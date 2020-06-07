package main

import "C"
import (
	"flag"
	"fmt"
	"github.com/lxn/win"
	"github.com/raintean/blink"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

// set GOARCH=386 // option
// go build -ldflags="-s -w -H windowsgui" -o ../ejs.exe ./crawler/sdk/blink/demo.go
// go build -tags="bdebug" -ldflags="-s -w -H windowsgui" -o ../ejs.exe ./crawler/sdk/blink/demo.go
// ejs.exe "http://app1.nmpa.gov.cn/datasearchcnda/face3/base.jsp?tableId=25&tableName=TABLE25&title=%B9%FA%B2%FA%D2%A9%C6%B7&bcId=152904713761213296322795806604" ./demo.js
func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "ejs eval js expression in the context of a web page\n")
		fmt.Fprintf(os.Stderr, "running in a headless instance of the Web browser.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "\tejs url script-file\n\n")
		fmt.Fprintf(os.Stderr, "url is the web page to execute the script in, and script-file is a local file\n")
		fmt.Fprintf(os.Stderr, "with the JavaScript you want to evaluate.\n\n")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "Example usage:\n\n")
		fmt.Fprintf(os.Stderr, "\tTo return the value of `document.title` on https://www.baidu.com:\n")
		fmt.Fprintf(os.Stderr, "\t    $ echo document.title | ejs https://www.baidu.com /dev/stdin\n")
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}
	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	logErr := log.New(os.Stderr, "", 0)

	pageURL := flag.Arg(0)
	script, scriptFile := "", flag.Arg(1)

	if _, err := url.Parse(pageURL); err != nil {
		logErr.Fatalf("Failed to parse URL %q: %s", pageURL, err)
		return
	}

	if fi, err := os.Stat(scriptFile); err == nil && !fi.IsDir() {
		jsTargetBytes, err := ioutil.ReadFile(scriptFile)
		if err != nil {
			logErr.Fatalf("Failed to open script file %q: %s", scriptFile, err)
			return
		}
		script = string(jsTargetBytes)
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
	mainWebView := blink.NewWebView(true,
		1266, 720, // 初始窗口大小
		int(win.GetSystemMetrics(win.SM_CXSCREEN)/5*4),
		int(win.GetSystemMetrics(win.SM_CYSCREEN)/5)) // 获取屏幕大小

	//mainWebView.ShowDockIcon()
	mainWebView.MoveToCenter()
	mainWebView.ShowWindow()
	mainWebView.ToTop()
	mainWebView.LoadURL(pageURL)

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
		if script != "" {
			time.Sleep(2 * time.Second)
			if strings.Index(script, ";") == 0 || strings.Index(script, "(function") == 0 {
				if _, err := mainWebView.Invoke(script); err != nil {
					logErr.Print(err)
				}
				return
			}
			if result, err := mainWebView.Invoke(script); err != nil {
				logErr.Print(err)
			} else {
				fmt.Println(result)
			}
		}
	}()

	// set exit
	<-done
}
