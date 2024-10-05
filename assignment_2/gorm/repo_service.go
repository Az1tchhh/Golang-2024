package gorm

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

func InsertUser(db *gorm.DB, name string, age int) error {
	user := User{Name: name, Age: age}
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("User added successfully: %s", name)
	fmt.Println()
	return nil
}

func UpdateUser(db *gorm.DB, id uint64, name string, age int) error {
	user := User{}
	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return result.Error
	}

	user.Age = age
	user.Name = name
	result = db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	fmt.Printf("User updated successfully: %s", name)
	fmt.Println()
	return nil
}

func QueryUsers(db *gorm.DB) error {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		log.Fatalf("Failed to query users: %v", result.Error)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("id: %d \nName: %s \nAge: %d \n", user.ID, user.Name, user.Age)
	}

	return nil
}

func DeleteUser(db *gorm.DB, id uint64) error {
	result := db.Delete(&User{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func DeleteTable(db *gorm.DB) error {
	result := db.Migrator().DropTable(&User{})
	if result != nil {
		return fmt.Errorf("failed to drop table 'users': %w", result)
	}

	fmt.Println("Table 'users' dropped successfully.")
	return nil
}
