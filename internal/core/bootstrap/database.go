package bootstrap

import (
	"database/sql"
	"fmt"
	"log"

	"goframe/internal/core/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
	rootDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort)

	sqlDB, err := sql.Open("mysql", rootDSN)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MySQL server: %v", err)
	}
	defer sqlDB.Close()

	var exists string
	err = sqlDB.QueryRow("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", cfg.DBName).Scan(&exists)
	if err != nil {

		fmt.Printf("⚠️ Database '%s' not found. Creating automatically...\n", cfg.DBName)
		_, err := sqlDB.Exec(fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", cfg.DBName))
		if err != nil {
			log.Fatalf("❌ Failed to create database: %v", err)
		}
		fmt.Printf("✅ Database '%s' successfully created!\n", cfg.DBName)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	DB = db
	fmt.Println("✅ Database connected")
}
