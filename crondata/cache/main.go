package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"time"
)

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	handler := requestHandler
	if *compress {
		handler = fasthttp.CompressHandler(handler)
	}

	InitCacheBackgroundWorker(10 * time.Second)

	if err := fasthttp.ListenAndServe(*addr, handler); err == nil {
		log.Printf("Listen And Serve: %s", *addr)
	} else {
		log.Fatalf("Error in Listen And Serve: %s", err)
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
