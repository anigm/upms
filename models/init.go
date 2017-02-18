package models

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID          int64      `gorm:"primary_key" json:"id"`
	CreatedAt   *time.Time `gorm:"type:datetime"  json:",omitempty"`
	UpdatedAt   *time.Time `gorm:"type:datetime"  json:",omitempty"`
	DeletedAt   *time.Time `gorm:"type:datetime" sql:"index" json:",omitempty"`
	Description string     `json:",omitempty"`
}

func ParseJson(str string, obj interface{}) {
	if err := json.Unmarshal([]byte(str), obj); err != nil {
		beego.Error(err)
		obj = nil
	}
}

func ToJson(obj interface{}) string {
	if bytes, err := json.Marshal(obj); err == nil {
		return string(bytes)
	} else {
		beego.Error(err)
		return ""
	}
}

func ParseJsonToMap(jsonStr string) map[string]string {
	var jsonMap map[string]string
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err == nil {
		return jsonMap
	} else {
		beego.Error(err)
		return nil
	}
}

func DBModelToJson(v interface{}) string {
	if bytes, err := json.Marshal(v); err == nil {
		r := regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+08:00)`)
		jsonstr := r.ReplaceAllStringFunc(string(bytes), func(s string) string {
			if t, err := time.Parse(time.RFC3339Nano, s); err != nil {
				panic(err)
			} else {
				return t.Format("2006-01-02 15:04:05")
			}
		})
		return jsonstr
	} else {
		panic(err)
	}
}

var connstr string

func DB() (*gorm.DB, error) {
	if db, err := gorm.Open("mysql", connstr); err != nil {
		return nil, err
	} else {
		if err = db.DB().Ping(); err != nil {
			return nil, err
		}
		if beego.AppConfig.String("runmode") == "dev" {
			//db.LogMode(true)
		}
		return db, nil
	}
}

func init() {
	beego.Info("Initializing Database...")
	connstr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqladdr"),
		beego.AppConfig.String("mysqlport"),
		beego.AppConfig.String("mysqldb"),
	)
	if db, err := DB(); err != nil {
		beego.Error("Initialize Database Error!")
		beego.Error(err)
		return
	} else {
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		// 	&User{}, &UserGroup{},
		// 	&Permission{}, &PermissionGroup{},
		// 	&Role{}, &RoleGroup{},
		// )
		beego.Info("Database Initializing Success！")
	}
	//InitAdmin()
	//InitTestData()
}

// func InitAdmin() {
// 	admin := User{Account: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", Name: "管理员"}
// 	SaveUser(&admin)
// }
