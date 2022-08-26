package libs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

func (c *DbConfig) InitMysqlDB() *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	return db
}
