package router

import (
	"github.com/gin-gonic/gin"
	"task-manage/handler"
	"task-manage/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors()) //开启中间件 允许使用跨域请求
	r.Use(middleware.LogMiddleware())

	// 把api 前缀写一个组
	apiGroup := r.Group("/task_api")
	// 获取当前加密方式
	apiGroup.POST("/getEncryptType", handler.GetEncryptTypeHandler)
	// 登陆
	apiGroup.POST("/login", handler.AdminLoginHandler)

	// 通过oauth 获取用户信息 登录
	/*apiGroup.POST("/oauthLogin", handler.GetUserInfoHandler)
	// 下载文件
	apiGroup.GET("/directories/file/download", handler.DownloadFile)
	// statistics 导出
	apiGroup.GET("/statistics/group/export", handler.StatisticsExport)
	// 获取oauth 登录地址
	apiGroup.GET("/oauth/url", handler.GetOauthUrl)

	apiGroup.Use(middleware.CheckToken)

	// 用户重置密码
	apiGroup.POST("/user/reset-password", handler.ResetPassword)

	// 用户登出
	apiGroup.POST("/logout", handler.Logout)

	// 用户列表
	apiGroup.POST("/user/list", handler.ListUser)

	// 用户添加
	apiGroup.POST("/user/add", handler.AddUser)

	// 用户删除
	//apiGroup.POST("/user/del", handler.DeleteUser)
	// 用户更新
	apiGroup.POST("/user/edit", handler.EditUser)

	//部门列表
	apiGroup.POST("/org/list", handler.ListOrg)
	//角色列表
	apiGroup.POST("/role/list", handler.ListRole)
	// 角色编辑
	apiGroup.POST("/role/edit", handler.EditRole)

	// 团队管理
	apiGroup.POST("/team/add", handler.AddTeam)
	// 团队查询
	apiGroup.POST("/team/list", handler.ListTeam)
	// 团队编辑
	apiGroup.POST("/team/edit", handler.EditTeam)

	// 版本手工新增
	apiGroup.POST("/version/add", handler.AddVersion)
	// 获取所有版本
	apiGroup.POST("/version/list", handler.ListVersion)
	// 获取所有的版本
	apiGroup.POST("/version/list/all", handler.ListAllVersion)
	// 版本编辑
	apiGroup.POST("/version/edit", handler.EditVersion)
	// 版本删除
	apiGroup.POST("/version/delete", handler.DeleteVersion)

	// 需求导入
	apiGroup.POST("/file/import", handler.ImportFile)
	// 导入版本确认
	apiGroup.POST("/requirement/import/confirm", handler.ImportRequirementConfirm)

	// 需求新增
	apiGroup.POST("/requirement/add", handler.AddRequirement)
	// 通过关键字搜索需求
	apiGroup.POST("/requirement/search", handler.SearchRequirement)
	// 需求列表
	apiGroup.POST("/requirement/list", handler.ListRequirement)
	// 需求详情
	apiGroup.POST("/requirement/detail", handler.RequirementDetail)
	// 需求状态变更
	apiGroup.POST("/requirement/status/update", handler.StatusUpdate)
	// 需求评审文档设置
	apiGroup.POST("/requirement/review/doc", handler.AddReviewDoc)
	// 需求评审文列表
	apiGroup.POST("/requirement/review/doc/list", handler.ListReviewDoc)
	// 需求评审文档删除
	apiGroup.POST("/requirement/review/doc/delete", handler.DeleteReviewDoc)
	// 需求删除
	apiGroup.POST("/requirement/delete", handler.DeleteRequirement)

	// 小版本新增
	apiGroup.POST("/smallVersion/add", handler.AddSmallVersion)
	// 小版本重命名
	apiGroup.POST("/smallVersion/rename", handler.RenameSmallVersion)

	// 小版本列表
	apiGroup.POST("/smallVersion/list", handler.ListSmallVersion)
	// 小版本删除
	apiGroup.POST("/smallVersion/delete", handler.DeleteSmallVersion)

	// bug 导入确认
	apiGroup.POST("/bug/import/confirm", handler.ImportBugConfirm)
	// 修改bug 状态
	apiGroup.POST("/bug/status/edit", handler.BugStatusEdit)
	// 批量修改bug 状态
	apiGroup.POST("/bug/status/edit/batch", handler.BugStatusEditBatch)
	// bug 列表
	apiGroup.POST("/bug/list", handler.ListBug)
	// bug 新增
	apiGroup.POST("/bug/add", handler.AddBug)
	// bug 删除
	apiGroup.POST("/bug/delete", handler.DeleteBug)

	// 新建项目
	apiGroup.POST("/project/add", handler.AddProject)
	// 项目列表
	apiGroup.POST("/project/list", handler.ListProject)
	// 获取所有的项目
	apiGroup.POST("/project/list/all", handler.ListAllProject)
	// 项目删除
	apiGroup.POST("/project/delete", handler.DeleteProject)
	// 项目编辑
	apiGroup.POST("/project/edit", handler.EditProject)

	//项目问题新建

	apiGroup.POST("/project/bug/add", handler.AddProjectBug)
	// 项目问题使用文本新增
	apiGroup.POST("/project/bug/addByText", handler.AddProjectBugByText)
	//  项目问题删除
	apiGroup.POST("/project/bug/delete", handler.DeleteProjectBug)
	// 项目问题列表
	apiGroup.POST("/project/bug/list", handler.ListProjectBug)
	// 项目问题编辑
	apiGroup.POST("/project/bug/edit", handler.EditProjectBug)

	// 项目定制开发新增
	apiGroup.POST("/project/custom/add", handler.AddProjectCustom)
	// 项目定制开发列表
	apiGroup.POST("/project/custom/list", handler.ListProjectCustom)

	// 项目定制开发详情
	apiGroup.POST("/project/custom/detail", handler.ProjectCustomDetail)
	//  项目定制开发编辑
	apiGroup.POST("/project/custom/edit", handler.EditProjectCustom)
	// 项目定制开发删除
	apiGroup.POST("/project/custom/delete", handler.DeleteProjectCustom)

	// 其他任务添加
	apiGroup.POST("/other/task/add", handler.AddOtherTask)
	// 其他任务列表
	apiGroup.POST("/other/task/list", handler.ListOtherTask)
	// 其他任务详情
	apiGroup.POST("/other/task/detail", handler.OtherTaskDetail)
	// 其他任务删除
	apiGroup.POST("/other/task/delete", handler.DeleteOtherTask)
	// 其他任务编辑
	apiGroup.POST("/other/task/edit", handler.EditOtherTask)

	// 培训计划
	//  新增一个培训计划
	apiGroup.POST("/train/plan/add", handler.AddTrainPlan)
	// 培训计划列表
	apiGroup.POST("/train/plan/list", handler.ListTrainPlan)
	// 培训计划删除
	apiGroup.POST("/train/plan/delete", handler.DeleteTrainPlan)
	// 培训计划编辑
	apiGroup.POST("/train/plan/edit", handler.EditTrainPlan)

	// 培训记录

	// 培训文档列表
	apiGroup.POST("/train/list", handler.ListTrain)
	// 培训文档删除
	apiGroup.POST("/train/delete", handler.DeleteTrainDoc)

	// 文档
	//  获取当前路径下的文件和目录
	apiGroup.POST("/directories/file/list", handler.ListDirectoriesFile)
	// 获取分享链接中的文件和目录
	apiGroup.POST("/directories/share/list", handler.ListShareLink)

	// 在一个目录下新建一个目录
	apiGroup.POST("/directories/add", handler.AddDirectories)
	// 在一个目录下新建一个文件
	apiGroup.POST("/directories/file/add", handler.AddDirectoriesFile)
	// 删除一个目录
	apiGroup.POST("/directories/delete", handler.DeleteDirectories)
	//  删除一个文件
	apiGroup.POST("/file/delete", handler.DeleteDirectoriesFile)
	// 目录或者文件重命名
	apiGroup.POST("/directories/rename", handler.RenameDirectories)

	// 获取文件或者文件夹分享链接
	apiGroup.POST("/directories/share", handler.ShareDirectories)

	// 奖惩记录
	apiGroup.POST("/reward/list", handler.ListReward)
	// 新增奖惩记录
	apiGroup.POST("/reward/add", handler.AddReward)
	// 删除奖惩记录
	apiGroup.POST("/reward/delete", handler.DeleteReward)

	// 获取所有组的统计信息
	apiGroup.POST("/statistics/group/list", handler.ListGroupStatistics)
	// 获取人员变动信息
	apiGroup.POST("/statistics/person/list", handler.ListPersonStatistics)*/

	return r

}
