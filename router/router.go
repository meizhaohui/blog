package router

import (
	"blog/conf"
	"blog/control"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run 入口
func Run() {
	engine := echo.New()
	// 初始渲染引擎
	engine.Renderer = initRender()
	// 恢复
	engine.Use(middleware.Recover())
	// 日志记录
	// engine.Use(middleware.LoggerWithConfig(logConfig))
	engine.Use(midLog)
	// 跨域设置
	engine.Use(middleware.CORSWithConfig(crosConfig))
	// 不显示横幅
	engine.HideBanner = true
	// 自定义错误处理
	engine.HTTPErrorHandler = HTTPErrorHandler
	engine.Debug = conf.Debug
	// 注册文档
	RegDocs(engine)
	// 静态目录
	engine.Static(`/static`, "static")
	// 前后端分离页面
	engine.File("/dashboard/*", "static/dist/index.html")
	// 首页
	engine.GET(`/`, control.IndexView)
	// 归档
	engine.GET(`/archives`, control.ArchivesView)
	// 标签
	engine.GET(`/tags`, control.TagsView)
	// 具体某个标签
	engine.GET(`/tag/:tag`, control.TagPostView)
	// 分类
	engine.GET(`/cate/:cate`, control.CatePostView)
	// 关于
	engine.GET(`/about`, control.AboutView)
	// 友链
	engine.GET(`/links`, control.LinksView)
	// 具体某个文章
	engine.GET(`/post/*`, control.PostView)
	// 具体某个页面
	engine.GET(`/page/*`, control.PageView)
	// ico
	engine.File(`/favicon.ico`, "favicon.ico")
	// 判断账号是否存在
	engine.GET(`/user/exist/:num`, control.UserExist)
	// 登陆
	engine.POST(`/login`, control.UserLogin)
	// 注销
	engine.POST(`/logout`, control.UserLogout)

	// 需要登陆才能访问
	api := engine.Group(`/api`, midJwt)
	// 服务器信息
	api.GET(`/sys`, control.Sys)
	// 统计信息
	api.GET(`/collect`, control.Collect)
	// 获取当前登陆信息
	api.GET(`/auth`, control.UserAuth)
	// 图片上传
	api.POST(`/upload`, control.Upload)
	// 修改自身信息
	api.POST(`/user/edit/self`, control.UserEditSelf)
	// 修改密码
	api.POST(`/user/pass`, control.UserPass)
	// 分类列表
	api.GET(`/cate/all`, control.CateAll)
	// 删除分类
	api.GET(`/cate/del/:id`, control.CateDel)
	// 通过分类查询文章
	api.GET(`/cate/post/:cid`, control.CatePost)
	// 添加分类
	api.POST(`/cate/add`, control.CateAdd)
	// 编辑分类
	api.POST(`/cate/edit`, control.CateEdit)
	// 文章
	api.GET(`/post/get/:id`, control.PostGet)
	// 通过文章id获取标签
	api.GET(`/post/tag/ids/:id`, control.PostTagIds)
	// 文章/页面-编辑/添加
	api.POST(`/post/opts`, control.PostOpts)
	// 删除文章/页面
	api.GET(`/post/del/:id`, control.PostDel)
	// 页面
	api.GET(`/page/all`, control.PostPageAll)
	// 标签列表
	api.GET(`/tag/all`, control.TagAll)
	// 删除标签
	api.GET(`/tag/del/:id`, control.TagDel)
	// 添加标签
	api.POST(`/tag/add`, control.TagAdd)
	// 编辑标签
	api.POST(`/tag/edit`, control.TagEdit)
	// 配置
	api.GET(`/opts/base`, control.OptsBase)
	// 获取配置项
	api.GET(`/opts/:key`, control.OptsGet)
	// 编辑配置项
	api.POST(`/opts/edit`, control.OptsEdit)
	engine.Start(conf.Addr)
}
