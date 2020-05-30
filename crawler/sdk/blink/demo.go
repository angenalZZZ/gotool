package main

import (
	"github.com/lxn/win"
	"github.com/raintean/blink"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// set GOARCH=386 // option
// go build -ldflags="-s -w -H windowsgui" -o ../demo.exe ./crawler/sdk/blink/demo.go
// go build -tags="bdebug" -ldflags="-s -w -H windowsgui" -o ../demo.exe ./crawler/sdk/blink/demo.go
// demo.exe "http://app1.nmpa.gov.cn/datasearchcnda/face3/base.jsp?tableId=25&tableName=TABLE25&title=%B9%FA%B2%FA%D2%A9%C6%B7&bcId=152904713761213296322795806604" "A:\go\src\github.com\angenalZZZ\gotool\crawler\sdk\blink\demo.js"
func main() {
	var (
		hasUri      bool
		urlTarget   string
		jsTarget    string
		mainWebView *blink.WebView
	)

	if len(os.Args) > 1 {
		if strings.Index(os.Args[1], "http") == 0 {
			hasUri, urlTarget = true, os.Args[1]
		}
		if len(os.Args) > 2 {
			jsTarget = os.Args[2]
		} else {
			jsTarget = os.Args[1]
		}
		if fi, err := os.Stat(jsTarget); err == nil && !fi.IsDir() {
			jsTargetBytes, err := ioutil.ReadFile(jsTarget)
			if err != nil {
				log.Fatal(err)
				return
			}
			jsTarget = string(jsTargetBytes)
		}
	}

	// set default url
	if hasUri == false {
		urlTarget = "https://www.baidu.com"
	}

	//调试模式
	blink.SetDebugMode(false)

	//初始化blink
	err := blink.InitBlink()
	if err != nil {
		log.Fatal(err)
		return
	}

	done := make(chan struct{})
	mainWebView = blink.NewWebView(true,
		1266, 720, // 初始窗口大小
		int(win.GetSystemMetrics(win.SM_CXSCREEN)/5*4),
		int(win.GetSystemMetrics(win.SM_CYSCREEN)/5)) // 获取屏幕大小

	//mainWebView.ShowDockIcon()
	mainWebView.MoveToCenter()
	mainWebView.ShowWindow()
	mainWebView.ToTop()
	mainWebView.LoadURL(urlTarget)

	// view.ShowDevTools()
	mainWebView.On("destroy", func(_ *blink.WebView) {
		mainWebView = nil
	})

	// await BlinkFunc.CloseWebPage()
	mainWebView.Inject("CloseWebPage", func() (int, error) {
		close(done) // exit
		return 0, nil
	})

	go func() {
		_ = mainWebView.GetWebTitle()
		if jsTarget != "" {
			//log.Println(jsTarget)
			time.Sleep(2 * time.Second)
			_, _ = mainWebView.Invoke(jsTarget)
		}
	}()

	// set exit
	<-done
}
