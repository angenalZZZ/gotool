package main

import (
	"github.com/lxn/win"
	"github.com/raintean/blink"
	"log"
	"time"
)

// go build -ldflags="-H windowsgui" -o ../table25.exe ./crawler/sdk/blink/demo.go
// go build -tags="bdebug" -ldflags="-H windowsgui" -o ../table25.exe ./crawler/sdk/blink/demo.go
func main() {
	//urlTarget := "https://www.baidu.com/"
	urlTarget := "http://app1.nmpa.gov.cn/datasearchcnda/face3/base.jsp?bcId=152904713761213296322795806604&tableId=25&tableName=TABLE25&title=%E5%9B%BD%E4%BA%A7%E8%8D%AF%E5%93%81"

	//启用调试模式
	//blink.SetDebugMode(true)

	//初始化blink
	err := blink.InitBlink()
	if err != nil {
		log.Fatal(err)
	}

	view := blink.NewWebView(true,
		1266, 720, // 初始窗口大小
		int(win.GetSystemMetrics(win.SM_CXSCREEN)/5*4),
		int(win.GetSystemMetrics(win.SM_CYSCREEN)/5)) // 获取屏幕大小
	view.LoadURL(urlTarget)
	view.ShowDockIcon()
	view.MoveToCenter()
	view.ShowWindow()
	view.ToTop()
	//view.ShowDevTools()

	go func() {
		title := view.GetWebTitle()
		log.Println(title)
		time.Sleep(3 * time.Second)

		js := `cb2020=function(id){
        request=new XMLHttpRequest();
        request.onreadystatechange=function(){
            if(request.readyState==4)
            {
                if(request.status==200)
                {
                    var res=request.responseText;
                    //console.log(res);
                    var t=res.substring(res.indexOf("<table "));
                    t = t.substring(0, t.indexOf("</table>")+8);
                    alert("服务器正常返回数据:国产药品:Id="+id+"  "+t);
                    request=null;
                }
                else
                {
                    alert("服务器未返回数据:国产药品:Id="+id)
                }
            }
        };
        request.open("GET","content.jsp?tableId=25&tableName=TABLE25&tableView=国产药品&Id="+id);
        request.setRequestHeader("Content-Type","text/html;encoding=gbk");
        request.send(null);
    };`

		log.Println(js)
		//C.runJSProxy(view.window, C.CString(``))
		//_, err = view.Invoke(js)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//_, err = view.Invoke(`cb2020`, `1`)
		//if err != nil {
		//	log.Fatalln(err)
		//}
	}()

	<-make(chan struct{})
}
