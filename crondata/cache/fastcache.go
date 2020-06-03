package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/angenalZZZ/gofunc/data/cache/fastcache"
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
	// maxBytes must be smaller than the available RAM size for the app,
	// since the cache holds data in memory.
	// If maxBytes is less than 32MB, then the minimum cache capacity is 32MB.
	CacheMaxBytes    = 1024 // default 1GB RAM size
	CacheWriterIndex = time.Now().Unix()
	CacheWriteMapper = map[int64]*CacheWriter{}
	CacheDataFileExt = "dat"
	// cache persist to disk directory
	CacheDataDirName = GetCurrentDir()
	// period of cache, beyond which it will be automatically persisted to disk
	CacheRenewFile = time.Minute
	// cycle of message processing
	ReleaseInterval = time.Minute
	// maximum number of messages that can be processed
	ReleaseMaxNumber = 60000
)

func InitCacheBackgroundWorker(renew time.Duration) {
	CacheRenewFile = renew
	readReady, writerReady := make(chan struct{}), make(chan struct{})
	go cacheReadBackgroundWorker(readReady)
	go cacheWriteBackgroundWorker(readReady, writerReady)
	<-writerReady
}

func cacheWriteBackgroundWorker(readReady <-chan struct{}, writerReady chan<- struct{}) {
	<-readReady
	start := time.Now().Unix()
	if _, ok := CacheWriteMapper[start]; ok == false {
		CacheWriteMapper[start] = &CacheWriter{
			Cache: fastcache.New(CacheMaxBytes),
			Done:  make(chan struct{}),
			Start: start,
			Index: 0,
		}
		go CacheWriteMapper[start].saveWorker()
	}
	CacheWriterIndex = start
	itemSeconds := int64(CacheRenewFile.Seconds())
	NextWriterIndex := start + itemSeconds
	writerReady <- struct{}{}
	//fmt.Println("waiting input new cache data ...")

	for {
		time.Sleep(time.Microsecond)
		start = time.Now().Unix()
		next := NextWriterIndex
		if start >= next {
			CacheWriteMapper[CacheWriterIndex].Done <- struct{}{}
			atomic.StoreInt64(&CacheWriterIndex, NextWriterIndex)
			NextWriterIndex += itemSeconds
		}
		if _, ok := CacheWriteMapper[next]; ok == false {
			CacheWriteMapper[next] = &CacheWriter{
				Cache: fastcache.New(CacheMaxBytes),
				Done:  make(chan struct{}),
				Start: next,
				Index: 0,
			}
			go CacheWriteMapper[next].saveWorker()
		}
	}
}

func cacheReadBackgroundWorker(readReady chan<- struct{}) {
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
			Start: start,
			Index: uint32(index),
		}
		CacheWriteMapper[start] = writer
	}
	if len(CacheWriteMapper) > 0 {
		//fmt.Println("loaded old cache data ...")
	}

	// read new data
	readReady <- struct{}{}
	for {
		time.Sleep(time.Microsecond)
		t, m := time.Now().Add(ReleaseInterval), ReleaseMaxNumber
		for start, c := range CacheWriteMapper {
			if m <= 0 || c == nil || start >= time.Now().Unix() {
				continue
			}
			m -= c.ReadAll(t, uint32(m))
		}
		if d := t.Sub(time.Now()); d.Milliseconds() > 0 {
			time.Sleep(d)
		}
	}
}

func GetCacheWriter() *CacheWriter {
	return CacheWriteMapper[CacheWriterIndex]
}

func (c *CacheWriter) ReadAll(endTime time.Time, maxNum uint32) (count int) {
	endTimeUnix := endTime.Unix()
	for i := uint32(1); i <= c.Index && i <= maxNum; i++ {
		if time.Now().Unix() >= endTimeUnix {
			return
		}

		dst := make([]byte, 0)
		buf := c.Get(dst, FromInt(i))
		if len(buf) <= 0 {
			count++
			continue
		}

		// Load the HTML document
		r := bytes.NewBuffer(buf)
		_, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, err)
			//f := c.filename()
			//_ = ioutil.WriteFile(f+".err", []byte(err.Error()), 0644)
		}

		//fmt.Println(doc)

		// Find the review items
		//doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		//	band := s.Find("a").Text()
		//	title := s.Find("i").Text()
		//	fmt.Printf("Review %d: %s - %s\n", i, band, title)
		//})

		count++
	}
	return
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
			time.Sleep(time.Second)
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
