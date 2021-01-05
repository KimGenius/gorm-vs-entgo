package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	println("gorm test start")
	dialector := postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", "postgres", "password", "localhost", 5432, "gorm-vs-ent-gorm", "disable"))
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	println(db)
}
