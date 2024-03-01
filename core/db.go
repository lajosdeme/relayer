package core

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/lajosdeme/transaction-relayer/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const TotalQuota = 20000000

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

func (db *Database) SubscribeUser(userId string, resetDate int) error {
	var localUser types.User

	tx := db.db.First(&localUser, "ID = ?", userId)
	if tx.Error != nil {
		return tx.Error
	}

	updates := types.Quota{
		Quota:      TotalQuota,
		Unit:       "gas",
		TotalQuota: TotalQuota,
		ResetDate:  resetDate,
	}

	tx = db.db.Model(&localUser).Updates(updates)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *Database) UpdateQuotaForUser(userId string, quotaUsed int) error {
	var localUser types.User

	tx := db.db.First(&localUser, "ID = ?", userId)
	if tx.Error != nil {
		return tx.Error
	}

	remainingQuota := localUser.Quota.Quota - quotaUsed

	updates := map[string]interface{}{
		"Quota": remainingQuota,
	}

	tx = db.db.Model(&localUser).Updates(updates)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *Database) UpdateVerifiedAddresses(id uuid.UUID, verifiedAddresses types.VerifiedAddresses) error {
	var localUser types.User

	tx := db.db.First(&localUser, "ID = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	updates := map[string]interface{}{
		"VerifiedAddresses": verifiedAddresses,
	}

	tx = db.db.Model(&localUser).Updates(updates)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *Database) GetUser(userId string) (types.User, error) {
	var localUser types.User

	tx := db.db.First(&localUser, "ID = ?", userId)
	if tx.Error != nil {
		return types.User{}, tx.Error
	}

	return localUser, nil
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
