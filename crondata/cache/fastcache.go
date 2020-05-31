package main

import (
	"encoding/binary"
	"fmt"
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
	CacheSaveToFile  = time.Minute
	CacheWriterIndex = time.Now().Unix()
	CacheWriteMapper = map[int64]*CacheWriter{}
)

func CacheWriteBackgroundWorker() {
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

func (c *CacheWriter) saveWorker() {
	for {
		select {
		case <-c.Done:
			if c.Index > 0 {
				f := fmt.Sprintf("%d.%d.dat", c.Start, c.Index)
				err := c.Cache.SaveToFile(f)
				if err == nil {
					_ = ioutil.WriteFile(f+".err", []byte(err.Error()), 0644)
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
