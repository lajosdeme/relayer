package core

import (
	"fmt"
	"sync"

	"github.com/lajosdeme/transaction-relayer/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

var (
	once     sync.Once
	database *Database
)

func DB() *Database {
	once.Do(func() {
		database = initDB()
	})
	return database
}

func (db *Database) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.db.First(dest, conds...)
}

func (db *Database) Create(u types.User) error {
	res := db.db.Create(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func initDB() *Database {
	return &Database{
		db: initialMigrationDB(),
	}
}

func initialMigrationDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("relayer.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&types.User{})

	fmt.Println("DB loaded")

	return db
}
