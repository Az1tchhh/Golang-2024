package main

import (
	"assignment_2/db"
	_ "assignment_2/docs"
	"assignment_2/gorm"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
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

	r := gin.Default()

	r.Use(ErrorHandlingMiddleware())

	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	// Start the server
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			var status int
			var errMessage string

			for _, err := range c.Errors {

				if err.Type == gin.ErrorTypePublic {
					status = http.StatusBadRequest
					errMessage = err.Error()
					continue
				}

				// PSQL
				var pqError *pq.Error
				if errors.As(err.Err, &pqError) {
					switch pqError.Code {
					case "23505": // unique_violation
						status = http.StatusConflict
						errMessage = "Conflict: " + pqError.Message
					case "23503": // foreign_key_violation
						status = http.StatusBadRequest
						errMessage = "Bad Request: Foreign key violation"
					default:
						status = http.StatusInternalServerError
						errMessage = "Database error: " + pqError.Message
					}
				} else {
					status = http.StatusInternalServerError
					errMessage = "Internal Server Error: " + err.Error()
				}
			}

			c.JSON(status, gin.H{"errors": errMessage})
		}
	}
}
