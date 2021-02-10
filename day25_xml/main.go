package main

import (
	"log"

	"github.com/beevik/etree"
)

//ExistNode  判断节点是否存在
func ExistNode(doc etree.Element, nodeName string) (exist bool) {
	path := etree.MustCompilePath(nodeName)
	bb := doc.FindElementPath(path)
	return bb != nil
}

// GetXpath 取指定的结点
func GetXpath(xmlstr string, nodename string) []*etree.Element {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlstr); err != nil {
		log.Panic(err)
	}
	return doc.FindElements(nodename)
	//  过去数据方法和xpath规则一致
	// for _, t := range doc.FindElements("//book[@category='WEB']/title") {
	// 	fmt.Println("Title:", t.Text())
	// }
}

// GetTree 取指定的结点
func GetTree(xmlstr string) *etree.Document {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlstr); err != nil {
		log.Panic(err)
	}
	return doc
}

// ScanDocument 取指定的结点
func ScanDocument(doc *etree.Document, nodename string) []*etree.Element {
	return doc.SelectElements(nodename)
}

// Scanpath 取指定的结点
func FindElements(doc *etree.Element, nodename string) []*etree.Element {
	return doc.FindElements(nodename)
}

// FindElement 取指定的结点
func FindElement(doc *etree.Element, nodename string) *etree.Element {

	return doc.FindElement(nodename)
}

func main() {
	xmlstr := ` 
	<input name="empcode" desc="员工编号"></input>
	<action result="Emp_Code,btime,etime" type="sql_select" >
	  <connstr>link:oaalwayson_read</connstr>
	  <sqlstr> <![CDATA[
	   select t.[Emp_Code],b.btime,b.etime from (
	   select  ? as [Emp_Code]   
	   ) t left join (
		  SELECT
		  [Emp_Code],substring(convert(varchar(20),[Att_Time],120),12,5) as btime
			  ,substring(convert(varchar(20),Att_EndTime,120),12,5) as etime
			  ,convert(decimal(18,2),datediff(minute,[Att_Time],getdate())/60.0) as worktime
	  FROM [TJOL_Attendance].[dbo].[T_Attendance_Init] where kqDay=convert(date,getdate()) 
	  and  [Emp_Code]  =? ) b
	  on t.[Emp_Code]=b.[Emp_Code]
	
	]]>
	  </sqlstr>
	  <sqlpara>{empcode}</sqlpara>
	  <sqlpara>{empcode}</sqlpara>
	</action>
	
	
	<resullt>
	  <return  name="Emp_Code" value="{empcode}" />
		<return  name="btime" value="{btime}" />
		<return  name="etime" value="{etime}" />
	 </resullt>

	 `
	for _, t := range GetXpath(xmlstr, "input") {
		log.Println("input:", t.Tag)
		log.Println("inputAttr:", t.Attr)
		log.Println("inputChild:", t.Child)
	}
	for _, t := range GetXpath(xmlstr, "action") {
		log.Println("action:", t.Text())
	}
	for _, t := range GetXpath(xmlstr, "resullt") {
		log.Println("resullt:", t.Text())
	}
}
