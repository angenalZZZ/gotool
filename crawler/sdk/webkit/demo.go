package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
	"github.com/sqs/gojs"
)

// set GOOS=linux
// go build -ldflags="-s -w -H windowsgui" -o ../ejs ./crawler/sdk/webkit/demo.go
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
		fmt.Fprintf(os.Stderr, "Notes:\n\n")
		fmt.Fprintf(os.Stderr, "\tBecause a headless WebKit instance is used, your $DISPLAY must be set. Use\n")
		fmt.Fprintf(os.Stderr, "\tXvfb if you are running on a machine without an existing X server. See\n")
		fmt.Fprintf(os.Stderr, "\thttps://sourcegraph.com/github.com/sourcegraph/go-webkit2/readme for more info.\n")
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

	runtime.LockOSThread()
	gtk.Init(nil)

	webView := webkit2.NewWebView()
	defer webView.Destroy()

	_, _ = webView.Connect("load-failed", func() {
		fmt.Println("Load failed.")
	})
	_, _ = webView.Connect("load-changed", func(_ *glib.Object, loadEvent webkit2.LoadEvent) {
		switch loadEvent {
		case webkit2.LoadFinished:
			webView.RunJavaScript(script, func(val *gojs.Value, err error) {
				if err != nil {
					logErr.Fatalf("javascript error: %s", err)
				} else {
					json, err := val.JSON()
					if err != nil {
						logErr.Fatalf("javascript serializable error: %v", err)
					}
					if len(json) > 0 {
						fmt.Println(string(json))
					}
				}
				if strings.Index(script, "(function") == 0 {
					<-make(chan struct{})
				} else {
					gtk.MainQuit()
				}
			})
		}
	})

	_, _ = glib.IdleAdd(func() bool {
		webView.LoadURI(pageURL)
		return false
	})

	gtk.Main()
}
