package repository

import (
	"fmt"

	// Un-comment below line to auto-generate table structure in MySQL
	// "go-web-template/modules/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IMySQLGorm interface {
	Get() *gorm.DB
	CreateMySQLConnection() *gorm.DB
}

var mysqlGorm *gorm.DB

type MySQLGorm struct{}

func CloseMySQL() {
	db, _ := mysqlGorm.DB()

	db.Close()
}

func (m *MySQLGorm) Get() *gorm.DB {
	if mysqlGorm == nil {
		return m.CreateMySQLConnection()
	}

	return mysqlGorm
}

func (*MySQLGorm) CreateMySQLConnection() *gorm.DB {
	if mysqlGorm != nil {
		return mysqlGorm
	}

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       "root:eee333rr@tcp(127.0.0.1:3307)/web?charset=utf8&parseTime=True&loc=Local",
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}),
		&gorm.Config{},
	)
	if err != nil {
		panic(fmt.Errorf("gorm initialization: MySQL connection failed:%s", err))
	}

	mysqlGorm = db

	// Un-comment below line to auto-generate table structure in MySQL
	// mysqlGorm.AutoMigrate(&model.User{}, &model.UserInfo{}, &model.UserRole{})

	return mysqlGorm
}
