package helper

import (
	"fmt"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/internal/config"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConMySQL(mySQLConf config.MySQLConf) (db *gorm.DB, err error) {

	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=true&writeTimeout=%s",
		mySQLConf.Username,
		mySQLConf.Password,
		mySQLConf.Host,
		mySQLConf.Port,
		mySQLConf.Database,
		mySQLConf.Charset,
		mySQLConf.Collation,
		mySQLConf.Timeout,
	)

	db, err = gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err = AutoMigrate(db); err != nil {
		return db, err
	}

	return db, err
}

func AutoMigrate(db *gorm.DB) (err error) {

	var tables = []interface{}{
		&models.Upload{},
	}

	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			err = db.AutoMigrate(table)
		}
	}

	return err
}
