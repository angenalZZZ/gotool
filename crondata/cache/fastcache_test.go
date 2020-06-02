package main

import (
	"github.com/angenalZZZ/gofunc/data/random"
	"testing"
	"time"
)

// go test -v -bench=^BenchmarkCacheWriter$ -cpu=4 -benchmem -benchtime=20s github.com/angenalZZZ/gotool/crondata/cache
// go test -c -o %TEMP%\t01.exe github.com/angenalZZZ/gotool/crondata/cache && %TEMP%\t01.exe -test.v -test.bench ^BenchmarkCacheWriter$ -test.run ^$
func BenchmarkCacheWriter(b *testing.B) {
	b.StopTimer()
	InitCacheBackgroundWorker(10 * time.Second)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := GetCacheWriter()
		_, _ = w.Write([]byte(random.AlphaNumberLower(20)))
	}
}
