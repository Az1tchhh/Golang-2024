package main

import (
	"assignment_2/db"
	"assignment_2/gorm"
	"fmt"
)

func main() {

	conn, err := db.Connect()

	if err != nil {
		fmt.Print("Error: " + err.Error())
		return
	}

	db.CreateTable(conn)

	db.InsertUser(conn, "John", 19)
	db.InsertUser(conn, "John2", 54)
	db.InsertUser(conn, "John3", 26)
	db.InsertUser(conn, "John4", 43)

	if err := db.QueryUsers(conn); err != nil {
		fmt.Printf("Failed to query users: %s\n", err)
	}

	err = db.DeleteTable(conn)
	if err != nil {
		return
	}

	// GORM
	fmt.Println("--------------------------------")

	gormConn, gormErr := gorm.Connect()

	if gormErr != nil {
		fmt.Print("Error: " + gormErr.Error())
		return
	}

	gorm.InsertUser(gormConn, "Gorm_User1", 224)
	gorm.InsertUser(gormConn, "Gorm_User2", 322)
	gorm.InsertUser(gormConn, "Gorm_User3", 441)

	gorm.QueryUsers(gormConn)

	gorm.UpdateUser(gormConn, 1, "azaza", 21)
	gorm.QueryUsers(gormConn)
	gorm.DeleteTable(gormConn)

}
