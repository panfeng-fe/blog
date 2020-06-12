package ctr

import (
	"blog/model"
	"errors"
	"net/http"

	"github.com/kataras/iris/v12"
)

func GetArticleList(ctx iris.Context) {

	page, pageErr := ctx.URLParamInt("page")
	size, sizeErr := ctx.URLParamInt("size")

	if pageErr != nil || sizeErr != nil || page < 1 || size < 15 {
		page = 1
		size = 15
	}

	dbRes, err := model.GetList(page, size)

	if err != nil || len(dbRes) < 1 {
		ctx.JSON(map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"msg":   err.Error(),
			"state": false,
		})
		return
	}

	ctx.JSON(map[string]interface{}{
		"data":  dbRes,
		"msg":   "获取数据成功",
		"state": true,
		"code":  http.StatusOK,
	})
}

func GetArticleListByKind(ctx iris.Context) {

	kind, KindErr := ctx.URLParamInt("kind")
	page, pageErr := ctx.URLParamInt("page")
	size, sizeErr := ctx.URLParamInt("size")

	if KindErr != nil {
		ctx.JSON(map[string]interface{}{
			"msg":   errors.New("文章类型参数错误"),
			"state": false,
			"code":  http.StatusInternalServerError,
		})
		return
	}

	if pageErr != nil || sizeErr != nil || page < 1 || size < 15 {
		page = 1
		size = 15
	}

	dbRes, err := model.GetListByKind(page, size, kind)

	if err != nil {
		ctx.JSON(map[string]interface{}{
			"msg":   err.Error(),
			"state": false,
			"code":  http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(map[string]interface{}{
		"data":  dbRes,
		"msg":   "获取数据成功",
		"state": true,
		"code":  http.StatusOK,
	})
}

func GetArticleByID(ctx iris.Context) {

	id, err := ctx.URLParamInt("id")

	if err != nil {
		ctx.JSON(map[string]interface{}{
			"msg":   "文章ID错误",
			"state": false,
			"code":  http.StatusInternalServerError,
		})
		return
	}

	article := model.Article{ID: id}

	err = article.GetArticleByID()

	if err != nil {
		ctx.JSON(map[string]interface{}{
			"msg":   err.Error(),
			"state": false,
			"code":  http.StatusInternalServerError,
		})
		return
	}

	go article.UpdateWatch()

	ctx.JSON(map[string]interface{}{
		"data":  article,
		"msg":   "获取数据成功",
		"state": true,
		"code":  http.StatusOK,
	})
}

func GetArticleKind(ctx iris.Context) {

	dbRes, err := model.GetArticleKind()

	if err != nil {
		ctx.JSON(map[string]interface{}{
			"msg":   err.Error(),
			"state": false,
			"code":  http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(map[string]interface{}{
		"data":  dbRes,
		"msg":   "获取数据成功",
		"state": true,
		"code":  http.StatusOK,
	})
}

func GetBloggerInfo(ctx iris.Context) {

	id, err := ctx.URLParamInt("id")

	if err != nil {
		ctx.JSON(map[string]interface{}{
			"msg":   err.Error(),
			"state": false,
			"code":  http.StatusInternalServerError,
		})
		return
	}

	dbRes, err1 := model.GetBloggerInfo(id)

	if err1 != nil {
		ctx.JSON(map[string]interface{}{
			"msg":   err1.Error(),
			"state": false,
			"code":  http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(map[string]interface{}{
		"data":  dbRes,
		"msg":   "获取数据成功",
		"state": true,
		"code":  http.StatusOK,
	})
}
