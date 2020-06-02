package main

import (
	"github.com/angenalZZZ/gofunc/data/random"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCacheDirName(t *testing.T) {
	basePath, _ := filepath.Abs("./")
	t.Log(basePath)

	basePath, _ = filepath.Abs(os.Args[0])
	basePath = filepath.Dir(basePath)
	t.Log(basePath)
}

func BenchmarkCacheWriter(b *testing.B) {
	b.StopTimer()
	CacheSaveToFile = 10 * time.Second
	CacheBackgroundWorker()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := GetCacheWriter()
		_, _ = w.Write([]byte(random.AlphaNumberLower(20)))
	}
}
