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
	temp = template.Must(template.ParseFiles("./template/" + path + ".txt"))
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
	{"comment", "Comment", "评论", 30},
	{"introduce", "Introduce", "内推记录", 30},
	{"message", "Message", "信息", 30},
	{"moments", "Moments", "动态", 30},
	{"organizationInformation", "OrganizationInformation", "组织信息", 30},
}

func maker(path string, list interface{}) {
	useTemplate(path)
	f, err := os.Create("./output/" + path + ".txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = temp.Execute(f, list)
	if err != nil {
		log.Fatalln(err)
	}
}

type webApi struct {
	MethodName string //方法名称
	Tags       string // 分类
	TagsName   string // 分类描述
	Summary    string // 描述
	Req        string //请求结构体
	Resp       string // 返回对象
	Router     string // 路径
	Method     string // 请求方法
}

func main() {
	// 生成api文件夹相关
	// var pathList = []string{"system_api", "menu", "casbin", "ensure_tables"}
	// for _, path := range pathList {
	// 	maker(path, apiList)
	// }

	// 生成web接口模板
	// var homeApiList = []webApi{
	// 	{"Info", "HomePage", "首页相关接口", "首页信息接口", "webReq.xxxx", "response.xxxxx", "/homepage", "GET"},
	// 	{"GetMomentList", "HomePage", "首页相关接口", "获取首页动态", "webReq.MomentsSearch", "response.PageResult", "/homepage/momentList", "GET"},
	// }

	var userApiList = []webApi{
		// {"GetUserInfo", "UserApi", "用户相关接口", "获取用户信息", "webReq.xxxx", "response.xxxxx", "/user/:id", "GET"},
		// {"GetMomentList", "UserApi", "用户相关接口", "获取用户相关动态", "webReq.xxxx", "response.xxxxx", "/user/:id/momentList", "GET"},
		// {"GetActivityList", "UserApi", "用户相关接口", "获取用户参加的活动", "webReq.xxxx", "response.xxxxx", "/user/:id/activityList", "GET"},
		// {"GetCommentList", "UserApi", "用户相关接口", "获取用户发表的评论", "webReq.xxxx", "response.xxxxx", "/user/:id/commentList", "GET"},

		// {"CreateArticle", "UserApi", "用户相关接口", "发布文章", "webReq.xxxx", "response.xxxxx", "/user/createArticle", "post"},
		// {"CreateActivity", "UserApi", "用户相关接口", "发布活动", "webReq.xxxx", "response.xxxxx", "/user/createActivity", "post"},
		// {"CreatComment", "UserApi", "用户相关接口", "发布评论", "webReq.xxxx", "response.xxxxx", "/user/creatComment", "post"},
		// {"CreateIntroduce", "UserApi", "用户相关接口", "发布内推", "webReq.xxxx", "response.xxxxx", "/user/createIntroduce", "post"},

		// 	- 参与活动
		//  - 填写个人简历、添加社团主页信息
		//  - 点赞/点踩评论
		//  - 收藏社团
		//  - 收藏文章
		// {"AttendActivity", "UserApi", "用户相关接口", "参与活动", "webReq.xxxx", "response.xxxxx", "/user/attendActivity", "post"},
		// {"CreateInformation", "UserApi", "用户相关接口", "填写简历/添加社团信息", "webReq.xxxx", "response.xxxxx", "/user/createInformation", "post"},
		// {"VoteComment", "UserApi", "用户相关接口", "点赞/点踩评论	", "webReq.xxxx", "response.xxxxx", "/user/voteComment", "post"},
		// {"Collection", "UserApi", "用户相关接口", "收藏社团/文章", "webReq.xxxx", "response.xxxxx", "/user/collection", "post"},
		{"GetHomePage", "WebOrganizationApi", "社团模块", "获取社团主页", "webReq.xxxx", "response.xxxxx", "/organization/homePage", "get"},
		{"GetDetail", "WebOrganizationApi", "社团模块", "社团详细页", "webReq.xxxx", "response.xxxxx", "/organization/detail", "get"},

		// 	- 获取社团主页(社团模块、社团评论)
		// 	- 社团详细页(社团详细、推荐相关社团)
		// 	- 获取社团评论
		// 	- 获取社团活动
	}

	maker("get_api", userApiList)
}

// TODO
// - 首页模块
// 	- 添加最新文章获取

// - userApiList
// 	- 发布文章 x
// 	- 发布活动 x
// 	- 发布评论 x
// 	- 发布内推 x

// 	- 参与活动 x
//  - 填写个人简历、添加社团主页信息 x
//  - 点赞/点踩评论 x
//  - 收藏社团 x
//  - 收藏文章 x

// 	- 拉取消息

// - 社团模块
// 	- 获取社团主页(社团模块、社团评论)
// 	- 社团详细页(社团详细、推荐相关社团)
// 	- 获取社团评论
// 	- 获取社团活动

// - 活动模块
// 	- 获取活动详细
// 	- 查看参与人员
// 	- 发送邮箱通知参与人员
// 	- 获取评论
