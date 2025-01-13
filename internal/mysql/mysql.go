package mysql

import (
	"shadowCloud/internal/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func New() (db *gorm.DB) {
	// 获取配置
	config := global.Config.Database
	//拼装dsn 定义数据库连接常量 dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=" + config.Charset + "&parseTime=True&loc=Local"

	// 创建gorm.DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	if global.Config.App.Debug {
		//开启mysql日志记录文本中
		// db.Logger = logger.NewMysqlLogger()

	}
	// 设置连接池大小
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}
	//SetMaxIdleConns设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(10)
	//SetMaxOpenConns设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//SetConnMaxLifetime设置连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db

}
