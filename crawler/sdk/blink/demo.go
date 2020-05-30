package main

import (
	"github.com/lxn/win"
	"github.com/raintean/blink"
	"log"
)

var mainWebView *blink.WebView

// set GOARCH=386
// go build -ldflags="-H windowsgui" -o ../blink.exe ./crawler/sdk/blink/demo.go
// go build -tags="bdebug" -ldflags="-H windowsgui" -o ../blink.exe ./crawler/sdk/blink/demo.go
func main() {
	//urlTarget := "https://www.baidu.com/"
	urlTarget := "http://app1.nmpa.gov.cn/datasearchcnda/face3/base.jsp?bcId=152904713761213296322795806604&tableId=25&tableName=TABLE25&title=%E5%9B%BD%E4%BA%A7%E8%8D%AF%E5%93%81"

	//调试模式
	blink.SetDebugMode(false)

	//初始化blink
	err := blink.InitBlink()
	if err != nil {
		log.Fatal(err)
	}

	mainWebView = blink.NewWebView(true,
		1266, 720, // 初始窗口大小
		int(win.GetSystemMetrics(win.SM_CXSCREEN)/5*4),
		int(win.GetSystemMetrics(win.SM_CYSCREEN)/5)) // 获取屏幕大小
	mainWebView.LoadURL(urlTarget)
	mainWebView.MoveToCenter()
	mainWebView.ShowDockIcon()
	mainWebView.ShowWindow()
	//view.ShowDevTools()
	mainWebView.ToTop()
	mainWebView.On("destroy", func(_ *blink.WebView) {
		mainWebView = nil
	})

	go func() {
		title := mainWebView.GetWebTitle()
		log.Println(title)

		mainWebView.Inject("title", "document.title")
		//mainWebView.Invoke()
	}()

	<-make(chan struct{})
}
