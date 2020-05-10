package main

import (
	"github.com/zserge/lorca"
	"log"
)

// go build -ldflags="-H windowsgui" -o ../table25.exe ./crawler/app1.sfda.gov.cn/table25/main.go
func main() {
	urlTarget := "http://app1.nmpa.gov.cn/datasearchcnda/face3/base.jsp?bcId=152904713761213296322795806604&tableId=25&tableName=TABLE25&title=%E5%9B%BD%E4%BA%A7%E8%8D%AF%E5%93%81"

	//w := webview.New(false)
	//defer w.Destroy()
	//w.SetSize(1120, 680, webview.HintFixed)
	//w.Navigate(urlTarget)
	//w.Run()

	// Create UI with basic HTML passed via data URI
	//ui, err := lorca.New(urlTarget, "", 1120, 680)
	//ui, err := lorca.New(urlTarget, "", 1120, 680, "--headless") // 不提供可视化页面
	ui, err := lorca.New(urlTarget, "", 1120, 680,
		"--disable-gpu",                      // 加上这个属性来规避bug
		"--disable-java",                     // 禁用java
		"--disable-sync",                     // 禁用同步
		"--disable-plugins",                  // 禁用插件
		"--hide-scrollbars",                  // 隐藏滚动条, 应对一些特殊页面
		"blink-settings=imagesEnabled=false", // 不加载图片, 提升速度
		//"--no-sandbox", // 以最高权限运行
		//"--start-maximized", // 启动时最大化
		//"--disable-javascript", // 禁用JavaScript
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
	// inject_commitForECMA("content.jsp?tableId=25&tableName=TABLE25&tableView=国产药品&Id=1")
	//ui.Eval(`function inject_commitForECMA(url)
	//{
	//	request=createXMLHttp();
	//	request.onreadystatechange=function () {
	//		if(request.readyState==4)
	//		{
	//			if(request.status==200)
	//			{
	//				console.log(request.responseText);
	//				request=null;
	//			}
	//			else
	//			{
	//				console.log('服务器未返回数据')
	//			}
	//		}
	//	};
	//	request.open("GET",url);
	//	request.setRequestHeader("Content-Type","text/html;encoding=gbk");
	//	request.send(null);
	//}`)

	// Get html content
	//ui.Eval(`window.onload=()=>{ document.body.innerHTML=''; }`)

	// Wait until UI window is closed
	<-ui.Done()
}
