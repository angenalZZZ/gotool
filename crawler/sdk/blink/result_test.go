package main

import (
	"github.com/angenalZZZ/gofunc/data/random"
	"testing"
)

func TestResultItem(t *testing.T) {
	s := `<table width=100% align=center>
<tr bgcolor="#659ace">
    <td height="25" colspan="2">
    	<div align="center" class=zs2 style="float:left;text-align:center;width:80%;padding-left:40px">国产药品</div>
    	<div style="float:right"><img src=images/data_fanhui.gif onclick=javascript:viewList() style=cursor:pointer></div>
    </td></tr>

<tr>
    <td bgcolor="#eaeaea" style="text-align:right" width=17% nowrap="true">批准文号</td>
    <td bgcolor="#eaeaea" width=83%>国药准字H21021039</td></tr>
           
<tr>
    <td bgcolor="#ffffff" style="text-align:right" width=17% nowrap="true">产品名称</td>
    <td bgcolor="#ffffff" width=83%>尼群地平片</td></tr>

<tr>
    <td bgcolor="#eaeaea" style="text-align:right" width=17% nowrap="true">英文名称</td>
    <td bgcolor="#eaeaea" width=83%>Nitrendipine  Tablets</td></tr>
           
<tr>
    <td bgcolor="#ffffff" style="text-align:right" width=17% nowrap="true">商品名</td>
    <td bgcolor="#ffffff" width=83%></td></tr>

<tr>
    <td bgcolor="#eaeaea" style="text-align:right" width=17% nowrap="true">剂型</td>
    <td bgcolor="#eaeaea" width=83%>片剂</td></tr>
           
<tr>
    <td bgcolor="#ffffff" style="text-align:right" width=17% nowrap="true">规格</td>
    <td bgcolor="#ffffff" width=83%>10mg</td></tr>

<tr>
    <td bgcolor="#eaeaea" style="text-align:right" width=17% nowrap="true">上市许可持有人</td>
    <td bgcolor="#eaeaea" width=83%></td></tr>
  
    
<tr>
    <td style="text-align:right">生产单位</td>
	
	
    <td><a href="javascript:commitForECMA(callbackC,'content.jsp?ytableId=25&tableId=34&tableName=TABLE34&linkId=COLUMN322&linkValue=富祥(大连)制药有限公司&Id=1',null);">富祥(大连)制药有限公司</a></td></tr>

<tr>
    <td bgcolor="#eaeaea" style="text-align:right" width=17% nowrap="true">生产地址</td>
    <td bgcolor="#eaeaea" width=83%>大连市旅顺经济开发区顺康街18号</td></tr>
           
<tr>
    <td bgcolor="#ffffff" style="text-align:right" width=17% nowrap="true">产品类别</td>
    <td bgcolor="#ffffff" width=83%>化学药品</td></tr>

<tr>
    <td bgcolor="#eaeaea" style="text-align:right" width=17% nowrap="true">批准日期</td>
    <td bgcolor="#eaeaea" width=83%>2015-07-30</td></tr>
           
<tr>
    <td bgcolor="#ffffff" style="text-align:right" width=17% nowrap="true">原批准文号</td>
    <td bgcolor="#ffffff" width=83%></td></tr>

<tr>
    <td bgcolor="#eaeaea" style="text-align:right" width=17% nowrap="true">药品本位码</td>
    <td bgcolor="#eaeaea" width=83%>86901110000020</td></tr>
           
<tr>
    <td bgcolor="#ffffff" style="text-align:right" width=17% nowrap="true">药品本位码备注</td>
    <td bgcolor="#ffffff" width=83%></td></tr>

<tr>
    <td style="text-align:right">相关数据库查询</td>
    <td>

        <a href="javascript:commitForECMA(callbackC,'content.jsp?ytableId=25&tableId=39&tableName=TABLE39&linkId=COLUMN422&linkValue=国药准字H21021039&Id=1',null)" >药品广告</a><br>

        <a href="javascript:commitForECMA(callbackC,'content.jsp?ytableId=25&tableId=22&tableName=TABLE22&linkId=COLUMN143&linkValue=国药准字H21021039&Id=1',null)" >中药保护品种库</a><br>
</td></tr>

<tr>
    <td></td>
    <td></td></tr>

<tr>
    <td bgcolor="#eaeaea" style="text-align:right">注</td>
    <td bgcolor="#eaeaea"><span style="FONT-SIZE: 14px; COLOR: #000066">说明 ：<br>　　企业用户如对药品数据信息有疑问，请及时与我局信息中心数据整理组联系，来电前请备好相应的批件证明材料以备工作人员查询。电话：88331520（工作日）；企业用户也可通过发邮件与我们联系：邮件地址yaopinshuju@nmpaic.org.cn，邮件主题请注明“药品批件问题”，邮件正文中请准确填写以下全部信息：1.药品批准文号/注册证号；2.药品批件号；3.药品批件类型（注册批件、补充批件、包材注册证、药品标准颁布件、再注册批件、其他）；4.问题描述（500字以内）；5.企业名称（全称）；6.统一社会信用代码；7.联系人姓名；8.联系电话（手机和座机）；9.电子邮件。以上内容请勿直接以电子邮件附件形式发送。</span></td></tr>
</table>`

	t.Log(len(s))

	s = random.AlphaNumberLower(5120)
	t.Log(len(s))
}
