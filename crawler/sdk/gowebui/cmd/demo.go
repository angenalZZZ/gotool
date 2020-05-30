package main

import (
	"github.com/angenalZZZ/gotool/crawler/sdk/gowebui"
	"strconv"
)

var mb gowebui.WebView

// set GOARCH=386
// go build -ldflags="-H windowsgui" -o ../gowebui.exe ./crawler/sdk/gowebui/cmd/demo.go
func main() {
	//urlTarget := "https://www.baidu.com/"
	//urlTarget := "http://app1.nmpa.gov.cn/datasearchcnda/face3/base.jsp?bcId=152904713761213296322795806604&tableId=25&tableName=TABLE25&title=%E5%9B%BD%E4%BA%A7%E8%8D%AF%E5%93%81"

	gowebui.Initialize("node.dll", "gonode.dll")
	gowebui.BindJsFunction("showAlert", showAlert, 99, 3)

	mb.CreateWebWindow("窗口", 0, 0, 60, 60, 1266, 720)
	mb.ShowWindow(true)

	mb.LoadHTML(`<html><head><title>测试窗口</title></head><body>
	<a href="https://www.baidu.com">点击打开百度</a><br>
	<a href="javascript:showAlert('一1一',2,true)">点击显示alert</a>
	</body></html>`)
}

func showAlert(es gowebui.JsExecState, param uintptr) uintptr {
	gowebui.StartCallBack()
	defer gowebui.EndCallBack()

	mb.RunJS("alert('链接被点击了，第1个参数为：" + mb.GetJsString(es, mb.GetJsValueFromArg(es, 0)) + "')")
	mb.RunJS("alert('链接被点击了，第2个参数为：" + strconv.Itoa(int(mb.GetJsInt(es, mb.GetJsValueFromArg(es, 1)))) + "')")

	if mb.GetJsBool(mb.GetJsValueFromArg(es, 2)) == true {
		mb.RunJS("alert('链接被点击了，第3个参数为：true')")
	} else {
		mb.RunJS("alert('链接被点击了，第3个参数为：false')")
	}
	return 0
}
