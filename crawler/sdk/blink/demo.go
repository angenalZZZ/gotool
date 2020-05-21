package main

import (
	"github.com/raintean/blink"
	"log"
)

// go build -tags="bdebug" -ldflags="-H windowsgui" -o ../table25.exe ./crawler/sdk/blink/demo.go
func main() {
	urlTarget := "https://www.baidu.com/"

	//启用调试模式
	blink.SetDebugMode(true)

	//初始化blink
	err := blink.InitBlink()
	if err != nil {
		log.Fatal(err)
	}

	view := blink.NewWebView(false, 1266, 720)
	view.LoadURL(urlTarget)
	view.MoveToCenter()
	view.ShowWindow()
	view.ShowDevTools()

	<-make(chan struct{})
}
