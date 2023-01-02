package main

import (
	"log"
	"os"
	"text/template"
)

func init() {
	os.Mkdir("output", 0777)
}

// Declare type pointer to a template
var temp *template.Template

// Using the init function to make sure the template is only parsed once in the program
func useTemplate(path string) {
	// template.Must takes the reponse of template.ParseFiles and does error checking
	temp = template.Must(template.ParseFiles(path + ".txt"))
}

type api struct {
	Name     string // 小写名字
	Uname    string // 大写开头
	Describe string //描述
	ParentId int    // 父级目录
}

var apiList = []api{
	{"article", "Article", "文章", 30},
	{"activity", "Activity", "活动", 30},
	{"activityRecord", "ActivityRecord", "活动记录", 30},
	{"introduce", "Introduce", "内推记录", 30},
	{"message", "Message", "信息", 30},
	{"moments", "Moments", "动态", 30},
	{"organizationInformation", "OrganizationInformation", "组织信息", 30},
}

func maker(path string) {
	useTemplate(path)
	f, err := os.Create("./output/" + path + ".txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = temp.Execute(f, apiList)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var pathList = []string{"system_api", "menu", "casbin", "ensure_tables"}
	for _, path := range pathList {
		maker(path)
	}
}
