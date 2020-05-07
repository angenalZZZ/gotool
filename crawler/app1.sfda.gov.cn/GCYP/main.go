package main

import (
	"github.com/zserge/webview"
)

// go build -ldflags="-H windowsgui" -o ../gcyp.exe ./crawler/app1.sfda.gov.cn/GCYP
func main() {
	w := webview.New(false)
	defer w.Destroy()

	url := "http://app1.nmpa.gov.cn/datasearchcnda/face3/base.jsp?tableId=25&tableName=TABLE25&title=国产药品&bcId=152904713761213296322795806604"

	w.SetSize(800, 600, webview.HintFixed)
	w.Navigate(url)
	w.Run()

	// Create UI with basic HTML passed via data URI
	//uiInit := "data:text/html," + url.PathEscape(`<html style="background-color:#fff"></html>`)
	//ui, err := lorca.New(uiInit, "", 1120, 680)
	////ui, err := lorca.New(uiInit, "", 1120, 680, "--headless") // Hide windows UI
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//defer func() { log.Fatal(ui.Close()) }()
	//
	//// Open url
	//if err = ui.Load(""); err != nil {
	//	log.Fatal(err)
	//	return
	//}

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
	//<-ui.Done()
}
