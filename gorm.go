package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Name      string  `json:"name"`
	OwnerName string  `json:"ownerName"`
	Items     []*Item `json:"items"`
}

type Item struct {
	gorm.Model
	Name    string `json:"name"`
	Price   uint   `json:"price"`
	StoreID uint   `json:"storeId"`
}

func main() {
	println("gorm test start")
	dialector := postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", "postgres", "password", "localhost", 5432, "gorm-vs-ent-gorm", "disable"))
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Store{}, &Item{})
	if err != nil {
		panic(err)
	}
	println(db)
}
