package database

import (
	"backEnd/models"
	"backEnd/pkg/mysql"
	"fmt"
)

// automatic migration
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
