package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/VictoriaMetrics/fastcache"
	"io/ioutil"
	"sync/atomic"
	"time"
)

type CacheWriter struct {
	*fastcache.Cache
	Done  chan struct{}
	Start int64
	Index uint32
}

var (
	ReleaseCacheSec  = 100
	CacheSaveToFile  = time.Minute
	CacheWriterIndex = time.Now().Unix()
	CacheWriteMapper = map[int64]*CacheWriter{}
)

func CacheWriteBackgroundWorker() {
	time.Sleep(time.Second)
	for {
		start := time.Now().Unix()
		end := start + int64(CacheSaveToFile.Seconds())
		tOk := start <= CacheWriterIndex && CacheWriterIndex <= end

		if tOk == false {
			CacheWriteMapper[CacheWriterIndex].Done <- struct{}{}
			atomic.StoreInt64(&CacheWriterIndex, start)
		}

		if _, ok := CacheWriteMapper[CacheWriterIndex]; ok == false {
			CacheWriteMapper[CacheWriterIndex] = &CacheWriter{
				Cache: fastcache.New(4),
				Done:  make(chan struct{}),
				Start: start,
				Index: 0,
			}
			go CacheWriteMapper[CacheWriterIndex].saveWorker()
		}

		time.Sleep(time.Microsecond)
	}
}

func CacheReadBackgroundWorker() {
	time.Sleep(time.Microsecond)
	// load old data

	for {
		for _, c := range CacheWriteMapper {
			var i uint32 = 1
			for ; i <= c.Index; i++ {
				dst := make([]byte, 0)
				v := c.Get(dst, FromInt(i))
				// Load the HTML document
				r := bytes.NewBuffer(v)
				doc, err := goquery.NewDocumentFromReader(r)
				if err != nil {
					f := c.filename()
					_ = ioutil.WriteFile(f+".err", []byte(err.Error()), 0644)
				}
				// Find the review items
				doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
					band := s.Find("a").Text()
					title := s.Find("i").Text()
					fmt.Printf("Review %d: %s - %s\n", i, band, title)
				})
			}
		}

		time.Sleep(time.Microsecond)
	}
}

func GetCacheWriter() *CacheWriter {
	if _, ok := CacheWriteMapper[CacheWriterIndex]; ok == false {
		time.Sleep(time.Microsecond)
	}
	return CacheWriteMapper[CacheWriterIndex]
}

func (c *CacheWriter) Write(p []byte) (n int, err error) {
	i := atomic.AddUint32(&c.Index, 1)
	c.Cache.Set(FromInt(i), p)
	return int(i), nil
}

func (c *CacheWriter) filename() string {
	return fmt.Sprintf("%d.%d.dat", c.Start, c.Index)
}

func (c *CacheWriter) writeError(err error) {
	_ = ioutil.WriteFile(c.filename()+".err", []byte(err.Error()), 0644)
}

func (c *CacheWriter) saveWorker() {
	for {
		select {
		case <-c.Done:
			time.Sleep(time.Microsecond)
			if c.Index > 0 {
				if err := c.Cache.SaveToFile(c.filename()); err != nil {
					c.writeError(err)
				}
			}
			return
		}
	}
}

// helper function which converts a int to a []byte in Big Endian
func FromInt(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

// helper function which converts a big endian []byte to an int
func ToInt(data []byte) uint32 {
	v := binary.BigEndian.Uint32(data)
	return v
}
