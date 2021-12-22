package database

import (
	"fmt"
	"log"
	"pedy/config"
	"time"

	"pedy/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func StartDatabase() {
	DbHost := config.GetConfig().DatabaseHost
	DbPort := config.GetConfig().DatabasePort
	DbUser := config.GetConfig().DatabaseUser
	DbName := config.GetConfig().DatabaseName
	DbSSlMode := config.GetConfig().DatabaseSslMode
	DbPass := config.GetConfig().DatabasePass
	DbMaxIddleConns := config.GetConfig().DatabaseMaxIdleConns
	DbMaxOpensConns := config.GetConfig().DatabaseMaxOpensConns

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", DbHost, DbPort, DbUser, DbName, DbSSlMode, DbPass)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := database.DB()

	config.SetMaxIdleConns(DbMaxIddleConns)
	config.SetMaxOpenConns(DbMaxOpensConns)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunAutoMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
