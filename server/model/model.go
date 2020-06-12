package model

import (
	"blog/conf"
	"errors"
	"time"
)

var (
	DB   = conf.DB
	Conf = conf.Config
)

type Article struct {
	ID        int       `gorm:"PRIMARY_KEY;UNIQUE;NOT NULL;AUTO_INCREMENT;column:id"`
	Title     string    `gorm:"NOT NULL;column:title;type:text"`
	Introduce string    `gorm:"NOT NULL;column:introduce;type:text"`
	Author    string    `gorm:"NOT NULL;column:author"`
	Context   string    `gorm:"NOT NULL;column:context;type:text"`
	AddTime   time.Time `gorm:"NOT NULL;column:addTime"`
	Watch     int       `gorm:"DEFAULT:0;column:watch"`
	Kind      int       `gorm:"NOT NULL;column:kind"`
	Remove    int       `gorm:"NOT NULL;default:0;column:remove"`
	KindName  string    `gorm:"NOT NULL;column:kindName"`
	IsTop     int       `gorm:"NOT NULL;column:isTop"`
	CNName    string    `gorm:"NOT NULL;column:cnName"`
}

func (A *Article) TableName() string {
	return "article"
}

type ArticleKind struct {
	Kind     int    `gorm:"column:kind"`
	KindName string `gorm:"column:kindName"`
	CNName   string `gorm:"column:cnName"`
	Count    int    `gorm:"column:count"`
}

type BloggerInfo struct {
	ID       int    `gorm:"column:id"`
	Describe string `gorm:"column:describe"`
}

func (A *Article) UpdateWatch() {
	A.Watch++
	DB.Model(A).Update(map[string]interface{}{"watch": A.Watch})
}

func (A *Article) AddArticle() error {
	isAdd := DB.NewRecord(A)
	if isAdd == false {
		return errors.New("文章已存在")
	}

	err := DB.Create(A).Error
	if err != nil {
		return err
	}

	isSuccess := DB.NewRecord(A)

	if isSuccess == true {
		return errors.New("文章添加失败")
	}

	return nil
}

//func init() {
//	//TableInit()
//}
//
//func TableInit() {
//	DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(
//		//&Article{},
//		//&ArticleKind{},
//		//&Comment{},
//	)
//}
