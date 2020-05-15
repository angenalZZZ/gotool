package main

// go build -ldflags="-H windowsgui" -o ../table25.exe ./crawler/app1.sfda.gov.cn/table25/main.go
func main() {
	urlTarget := "http://app1.nmpa.gov.cn/datasearchcnda/face3/base.jsp?bcId=152904713761213296322795806604&tableId=25&tableName=TABLE25&title=%E5%9B%BD%E4%BA%A7%E8%8D%AF%E5%93%81"

	println(urlTarget)

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
}
