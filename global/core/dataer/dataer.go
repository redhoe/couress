package dataer

import (
	"fmt"
	"github.com/redhoe/couress/global/core/confer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// Living 数据库初始化(读写分离)

var dbLiving *gorm.DB = nil

func dbLivingInit() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(confer.AppConfServer.Mysql.MysqlDns()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	// 使用读写分离
	//if err := db.Use(dbresolver.Register(dbresolver.Config{
	//	Sources:  []gorm.Dialector{mysql.Open(confer.AppConfServer.Mysql.MysqlDns())},
	//	Replicas: []gorm.Dialector{mysql.Open(confer.AppConfServer.MysqlSlave.MysqlDns())},
	//	Policy:   dbresolver.RandomPolicy{},
	//})); err != nil {
	//	fmt.Println(err.Error())
	//	return err
	//}

	d, _ := db.DB()
	d.SetConnMaxIdleTime(20 * time.Second)
	d.SetConnMaxLifetime(20 * time.Second)
	d.SetMaxIdleConns(1000)
	d.SetMaxOpenConns(1500)
	err = d.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	dbLiving = db
	return nil
}

func Living() *gorm.DB {
	if dbLiving == nil {
		if err := dbLivingInit(); err != nil {
			panic(err)
		}
	}
	return dbLiving
}

// LivingMain 数据库初始化(单库)

var dbLivingMain *gorm.DB = nil

func dbLivingInitMain() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(confer.AppConfServer.Mysql.MysqlDns()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	d, _ := db.DB()
	d.SetConnMaxIdleTime(20 * time.Second)
	d.SetConnMaxLifetime(20 * time.Second)
	d.SetMaxIdleConns(1000)
	d.SetMaxOpenConns(1500)
	err = d.Ping()
	if err != nil {
		return err
	}
	dbLivingMain = db
	return nil
}

func LivingMain() *gorm.DB {
	if dbLivingMain == nil {
		if err := dbLivingInitMain(); err != nil {
			panic(err)
		}
	}
	return dbLivingMain
}
