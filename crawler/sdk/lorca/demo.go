package main

import (
	"github.com/zserge/lorca"
	"log"
)

// go build -ldflags="-s -w -H windowsgui" -o ../lorca.exe ./crawler/sdk/lorca/demo.go
func main() {
	urlTarget := "https://www.baidu.com/"

	// Create UI with basic HTML passed via data URI
	//ui, err := lorca.New(urlTarget, "", 1266, 720) // 默认
	//ui, err := lorca.New(urlTarget, "", 1266, 720, "--headless") // 不提供可视化页面
	ui, err := lorca.New(urlTarget, "", 1266, 720, "--disable-gpu", // 规避bug
		"--disable-http2", // 禁用http2
		"--disable-java",  // 禁用java
		//"--disable-javascript", // 禁用javascript
		"--disable-sync",    // 禁用同步
		"--disable-plugins", // 禁用插件
		"--hide-scrollbars", // 隐藏滚动条, 应对一些特殊页面
		"--disable-databases --disable-local-storage", // 禁用h5本地缓存数据库
		//"blink-settings=imagesEnabled=false",          // 不加载图片以提升速度
		//"--no-sandbox", // 以最高权限运行
		"--start-maximized", // 启动时最大化
		"--disable-crash-reporter")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() { log.Fatal(ui.Close()) }()

	// Open url
	if err = ui.Load(urlTarget); err != nil {
		log.Fatal(err)
		return
	}

	// Inject javascript

	// Get html content
	//ui.Eval(`window.onload=()=>{ document.body.innerHTML=''; }`)

	// Wait until UI window is closed
	<-ui.Done()
}
