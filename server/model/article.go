package model

import (
	"errors"
)

func GetList(page, size int) (res []Article, err error) {
	err = DB.Debug().Table("article").Select("id,title,introduce,author,addTime,watch,kind,kindName,cnName").
		Limit(size).Offset((page - 1) * size).Order("addTime desc").Scan(&res).Error

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("没有书籍")
	}

	return res, nil
}

func GetListByKind(page, size, kind int) (res []Article, err error) {

	err = DB.Table("article").Select("id,title,introduce,author,addTime,watch,kind,kindName,cnName").
		Where("kind = ? ", kind).Limit(size).Offset((page - 1) * size).Order("addTime desc").Scan(&res).Error

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("不存在此种书籍")
	}

	return res, nil
}

func (A *Article) GetArticleByID() error {
	err := DB.Table("article").Where("id = ? ", A.ID).Find(A).Error
	if err != nil {
		return err
	}

	if A.Kind == 0 {
		return errors.New("不存在文章")
	}

	return nil
}

func GetArticleKind() (res []ArticleKind, err error) {
	err = DB.Table("article").Select("kindName,cnName,kind,count(id) as count").Group("kindName,cnName,kind").Scan(&res).Error

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("文章列表为空")
	}

	return res, nil
}

func GetBloggerInfo(id int) (res BloggerInfo, err error) {
	err = DB.Table("setting").Where("id = ? ", id).Scan(&res).Error
	return res, err
}
