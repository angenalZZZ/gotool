package main

import (
	"github.com/angenalZZZ/gofunc/data/random"
	"testing"
	"time"
)

// go test -v -cpu=4 -benchtime=15s -benchmem -bench=^BenchmarkCacheWriter$ github.com/angenalZZZ/gotool/crondata/cache
// go test -c -o %TEMP%\t01.exe github.com/angenalZZZ/gotool/crondata/cache && %TEMP%\t01.exe -test.v -test.bench ^BenchmarkCacheWriter$ -test.run ^$
func BenchmarkCacheWriter(b *testing.B) {
	b.StopTimer()
	InitCacheBackgroundWorker(10 * time.Second)
	//l := 5120 // every time 5kB data request: cpu=4 1200k/qps 0.8ms/op
	//l := 2018 // every time 2kB data request: cpu=4 1800k/qps 0.5ms/op
	//l := 1024 // every time 1kB data request: cpu=4 2400k/qps 0.4ms/op
	l := 128 // every time 128B data request: cpu=4 3200k/qps 0.3ms/op
	p := []byte(random.AlphaNumberLower(l))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := GetCacheWriter()
		_, _ = w.Write(p)
	}
}
