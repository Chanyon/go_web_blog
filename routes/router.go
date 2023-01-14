package routes

import (
	v1 "go_blog/api/v1"
	"go_blog/middleware"
	"go_blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// debug mode
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.MaxMultipartMemory = 8 << 20 // 限制8 MiB
	r.Static("/upload", "./upload")
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	// 后台管理路由接口
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// user 模块
		auth.GET("admin/users", v1.GetUsers)
		auth.POST("admin/users", v1.AddUser)
		auth.PUT("admin/user/:id", v1.EditUser)
		auth.GET("admin/user/:id", v1.GetUserInfo)
		auth.DELETE("admin/user/:id", v1.DeleteUser)
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)

		// 文章模块
		auth.POST("admin/article/add", v1.CreateArticle)
		auth.GET("admin/article/info/:id", v1.GetArticleInfo)
		auth.GET("admin/article", v1.GetAllArticles) //文章列表
		auth.PUT("admin/article/:id", v1.EditArticle)
		auth.DELETE("admin/article/:id", v1.DeleteArticle)
		auth.GET("admin/article/category/:id", v1.GetCategoryArticle)

		// 分类模块
		auth.POST("admin/category/add", v1.CreateCategory)
		auth.DELETE("admin/category/:id", v1.DeleteCategory)
		auth.PUT("admin/category/:id", v1.EditCategory)
		auth.GET("admin/category", v1.GetCategory)
		auth.GET("admin/category/:id", v1.GetCategoryInfo)
		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("admin/profile/:id", v1.UpdateProfile)
		auth.POST("admin/profile/", v1.CreatProfile)

		// 文件上传
		auth.POST("admin/upload", v1.UploadFile)

		// 评论模块
		auth.GET("admin/comment/list", v1.GetCommentList)
		auth.PUT("admin/check/comment/:id", v1.CheckComment)
		auth.PUT("admin/uncheck/comment/:id", v1.UnCheckComment)
		auth.DELETE("admin/delete/comment/:id", v1.DeleteComment)
	}

	login := r.Group("api/v1")
	{
		// 登录
		login.POST("admin/login", v1.Login)
		login.POST("user/login", v1.LoginFront)
	}

	//前台路由
	router := r.Group("api/v1")
	{
		// 用户模块
		router.POST("user/add", v1.AddUser) //用户注册
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章模块
		router.GET("user/article/info/:id", v1.GetArticleInfo)
		router.GET("user/article/", v1.GetAllArticles)
		router.GET("user/article/category/:id", v1.GetCategoryArticle)

		// 分类模块
		router.GET("user/category", v1.GetCategory)
		router.GET("user/category/:id", v1.GetCategoryInfo)

		// 获取个人设置
		router.GET("user/profile/:id", v1.GetProfile)
		// 评论模块
		router.POST("user/comment/create", v1.CreateComment)
		router.GET("user/comment/info/:id", v1.GetComment)
		router.GET("user/comment/front/:id", v1.GetCommentListFront)
		router.GET("user/comment/count/:id", v1.GetCommentCount)
	}
	_ = r.Run(utils.HttpPort)
}
