package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/VictoriaMetrics/fastcache"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	//ReleaseCacheSec  = 100
	CacheSaveToFile  = time.Minute
	CacheWriterIndex = time.Now().Unix()
	CacheWriteMapper = map[int64]*CacheWriter{}
	CacheDataDirName = GetCurrentDir()
	CacheDataFileExt = "dat"
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
	oldFiles, _ := filepath.Glob(filepath.Join(CacheDataDirName, "*"+CacheDataFileExt))
	for _, oldFile := range oldFiles {
		_, f := filepath.Split(oldFile)
		s := strings.Split(f, ".")
		start, _ := strconv.ParseInt(s[0], 10, 0)
		index, _ := strconv.ParseInt(s[1], 10, 0)
		cache, err := fastcache.LoadFromFile(oldFile)
		if err != nil || cache == nil {
			continue
		}
		writer := &CacheWriter{
			Cache: cache,
			Done:  make(chan struct{}),
			Start: start,
			Index: uint32(index),
		}
		CacheWriteMapper[start] = writer
	}

	// read new data
	for {
		for _, c := range CacheWriteMapper {
			for i := uint32(1); i <= c.Index; i++ {
				dst := make([]byte, 0)
				v := c.Get(dst, FromInt(i))
				// Load the HTML document
				r := bytes.NewBuffer(v)
				doc, err := goquery.NewDocumentFromReader(r)
				if err != nil {
					//f := c.filename()
					//_ = ioutil.WriteFile(f+".err", []byte(err.Error()), 0644)
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
	return filepath.Join(CacheDataDirName, fmt.Sprintf("%d.%d.%s", c.Start, c.Index, CacheDataFileExt))
}

//func (c *CacheWriter) writeError(err error) {
//	_ = ioutil.WriteFile(c.filename()+".log", []byte(err.Error()), 0644)
//}

func (c *CacheWriter) saveWorker() {
	for {
		select {
		case <-c.Done:
			time.Sleep(time.Microsecond)
			if c.Index > 0 {
				if err := c.Cache.SaveToFileConcurrent(c.filename(), 0); err != nil {
					//c.writeError(err)
				}
			}
			return
		}
	}
}

// GetCurrentDir helper function
func GetCurrentDir() (basePath string) {
	basePath, _ = filepath.Abs(os.Args[0])
	basePath = filepath.Dir(basePath)
	return
}

// FromInt helper function which converts a int to a []byte in Big Endian
func FromInt(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

// ToInt helper function which converts a big endian []byte to an int
func ToInt(data []byte) uint32 {
	v := binary.BigEndian.Uint32(data)
	return v
}
