package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/gcfg.v1"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	Name string
	Port string
	Key  string
}

type Mysql struct {
	User     string
	Password string
	Host     string
	DateBase string
	Port     int
	Dsn      string
}

type Conf struct {
	App
	Mysql
}

var (
	Config Conf
	DB     *gorm.DB
)

func (C *Conf) Init() {
	if err := gcfg.ReadFileInto(C, "./conf/server.ini"); err != nil {
		panic(err)
	}
}

func (M *Mysql) Init() *gorm.DB {
	MysqlDsn := fmt.Sprintf(M.Dsn, M.User, M.Password, M.Host, M.DateBase)
	DB, err := gorm.Open("mysql", MysqlDsn)

	if err != nil {
		panic(err)
	}

	// DB.LogMode(true)
	DB.DB().SetMaxOpenConns(5000)
	DB.DB().SetMaxIdleConns(2000)
	DB.SingularTable(true)
	return DB
}

func init() {
	Config.Init()
	DB = Config.Mysql.Init()
}
