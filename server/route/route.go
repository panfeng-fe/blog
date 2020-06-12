package route

import (
	"blog/ctr"

	"github.com/kataras/iris/v12"
)

func RouterInit(router iris.Party) {
	blogInit(router.Party("blog"))
	//AdminInit(router.Party("admin/"))
}

func blogInit(route iris.Party) {
	route.Get("/getArticleList", ctr.GetArticleList)
	route.Get("/getArticleListByKind", ctr.GetArticleListByKind)

	route.Get("/getArticleByID", ctr.GetArticleByID)
	route.Get("/getArticleKind", ctr.GetArticleKind)
	route.Get("/getBloggerInfo", ctr.GetBloggerInfo)

	//BlogRoute.Post("/postComment", ctr.PostComment)
	//BlogRoute.Get("/commentByArticleID", ctr.CommentByArticleID)
}
