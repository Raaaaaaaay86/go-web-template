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

type MySQLGorm struct {
	DB_MYSQL_USERNAME string
	DB_MYSQL_PASSWORD string
	DB_MYSQL_HOST     string
	DB_MYSQL_PORT     string
	DB_MYSQL_SCHEMA   string
}

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

func (mg *MySQLGorm) CreateMySQLConnection() *gorm.DB {
	if mysqlGorm != nil {
		return mysqlGorm
	}

	serverDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mg.DB_MYSQL_USERNAME,
		mg.DB_MYSQL_PASSWORD,
		mg.DB_MYSQL_HOST,
		mg.DB_MYSQL_PORT,
		mg.DB_MYSQL_SCHEMA,
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
