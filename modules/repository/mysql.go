package repository

import (
	"fmt"
	"os"

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

	serverDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_MYSQL_USERNAME"),
		os.Getenv("DB_MYSQL_PASSWORD"),
		os.Getenv("DB_MYSQL_HOST"),
		os.Getenv("DB_MYSQL_PORT"),
		os.Getenv("DB_MYSQL_SCHEMA"),
	)

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       serverDSN,
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
