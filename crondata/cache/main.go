package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

// go build -ldflags="-s -w" -o ../token.exe github.com/angenalZZZ/gotool/crondata/cache
func main() {
	flag.Parse()
	fmt.Println("Parse args ...")
	fmt.Printf(" addr = %q \n", *addr)
	fmt.Printf(" compress = %t \n", *compress)

	handler := requestHandler
	if *compress {
		handler = fasthttp.CompressHandler(handler)
	}

	fmt.Println("Init cache ...")
	InitCacheBackgroundWorker(time.Minute)

	fmt.Println("Start listen and serve ...")
	if err := fasthttp.ListenAndServe(*addr, handler); err != nil {
		fmt.Printf("Error in listen and serve: %s \n", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	w := GetCacheWriter()
	n, err := w.Write(ctx.Request.Body())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
	} else {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("text/plain; charset=utf8")
		_, _ = fmt.Fprintf(ctx, strconv.FormatInt(int64(n), 10))
	}
}
