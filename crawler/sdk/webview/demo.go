package main

import "github.com/zserge/webview"

func main() {
	urlTarget := "https://www.baidu.com/"

	w := webview.New(false)
	defer w.Destroy()
	w.SetSize(1266, 720, webview.HintFixed)
	w.Navigate(urlTarget)
	w.Run()
}
