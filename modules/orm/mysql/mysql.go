package mysql

import (
	"fmt"
	"os"

	// Un-comment below line to auto-generate table structure in MySQL
	// "go-web-template/modules/model"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//go:generate mockery --dir . --filename mock_mysql.go --name IMySQLGorm --output ../../mocks
type IMySQLGorm interface {
	Get() *gorm.DB
	CreateMySQLConnection() *gorm.DB
}

type MySQLGorm struct {
	DB_MYSQL_USERNAME string
	DB_MYSQL_PASSWORD string
	DB_MYSQL_HOST     string
	DB_MYSQL_PORT     string
	DB_MYSQL_SCHEMA   string
}

var mysqlGorm *gorm.DB

var MySQLOrmSet = wire.NewSet(
	wire.Bind(new(IMySQLGorm), new(*MySQLGorm)),
	MySQLGormProvider,
)

func MySQLGormProvider() *MySQLGorm {
	return &MySQLGorm{
		DB_MYSQL_USERNAME: os.Getenv("DB_MYSQL_USERNAME"),
		DB_MYSQL_PASSWORD: os.Getenv("DB_MYSQL_PASSWORD"),
		DB_MYSQL_HOST:     os.Getenv("DB_MYSQL_HOST"),
		DB_MYSQL_PORT:     os.Getenv("DB_MYSQL_PORT"),
		DB_MYSQL_SCHEMA:   os.Getenv("DB_MYSQL_SCHEMA"),
	}
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
