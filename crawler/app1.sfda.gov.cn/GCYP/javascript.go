package main

const (
	makeVisibleScript = `setTimeout(function() {
	document.querySelector('#box1').style.display = '';
}, 3000);`

	injectCommitForECMA = `function injectCommitForECMA(url)
{
	request=createXMLHttp();
	request.onreadystatechange=function () {
		if(request.readyState==4)
		{
			if(request.status==200)
			{
				console.log(request.responseText);
				request=null;
			}
			else
			{
				console.log('服务器未返回数据')
			}
		}
	};
	request.open("GET",url);
	request.setRequestHeader("Content-Type","text/html;encoding=gbk");
	request.send(null);
}`
)

func DO() {
	// Get html content
	//ui.Eval(`window.onload=()=>{ document.body.innerHTML=''; }`)

	// Inject javascript
	// inject_commitForECMA("content.jsp?tableId=25&tableName=TABLE25&tableView=国产药品&Id=1")
}
