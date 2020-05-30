(function () {
    // create a mask layer
    var id0 = 'script-20200529';
    if (document.getElementById(id0)) return;
    var div0 = document.createElement('div');
    div0.id = id0;
    div0.style = 'background-color:rgba(0,0,0,0.5);position:fixed;z-index:999999;top:0;left:0;text-align:center;right:0;padding-top:calc(50vh-100px);bottom:0';
    document.body.append(div0);
    div0 = document.getElementById(id0);
    // create a show text layer
    var loading = document.createElement('div');
    loading.textContent = 'Loading...';
    loading.style = 'position:absolute;z-index:1;top:20px;left:20px;color:#fff;font-size:16px;font-weight:bold;cursor:default';
    div0.append(loading);
    var setContent = function (txt, add) { if (add) loading.innerHTML += txt; else loading.innerHTML = txt; };
    // create a close button
    var div1 = document.createElement('div');
    div1.style = 'position:absolute;z-index:2;top:50%;left:50%';
    var btn1 = document.createElement('button');
    btn1.style = 'color:#f39;font-size:16px;font-weight:bold;padding:6px;border-radius:.5em';
    btn1.innerText = 'Close';
    btn1.addEventListener('click', async () => {
        return await BlinkFunc.CloseWebPage();
    });
    div1.append(btn1); div0.append(div1);

    // create ajax request
    var getItemById = function (id, idMax, fn0, fn1) {
        var request = new XMLHttpRequest();
        request.onreadystatechange = function () {
            if (request.readyState === 4) {
                if (request.status === 200) {
                    var t = request.responseText;
                    t = t.substring(t.indexOf("<table "));
                    t = t.substring(0, t.indexOf("</table>") + 8);
                    fn0(id, t);
                } else {
                    fn1(id, request.responseText);
                }
                if (id++ <= idMax) {
                    getItemById(id, idMax, fn0, fn1);
                }
            }
        };
        request.open("GET", "content.jsp?tableId=25&tableName=TABLE25&tableView=国产药品&Id=" + id);
        request.setRequestHeader("Content-Type", "text/html;encoding=gbk");
        request.send(null);
    };

    // do ajax request
    getItemById(1, 3, function (id1, t1) {
        setContent('<p style="color:#3f9">成功: 国产药品: Id=' + id1 + '</p>', id1 > 1);
    }, function (t) {
        setContent('<p style="color:#f93">失败: 国产药品: Id=' + id1 + '</p>', id1 > 1);
    });
})();