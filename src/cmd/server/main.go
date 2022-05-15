package main

import (
	"database/sql"
	"fmt"

	_ "go-server-docker/src/controllers/users"
	_ "go-server-docker/src/pkg/greeting"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	usr_id       int
	usr_name     string
	usr_username string
	usr_email    string
	usr_phone    string
	usr_address  string
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:root@tcp(localhost)/achats")
		// create array
		var users []User
		if err != nil {
			panic(err.Error())
		}

		err = db.Ping()
		if err != nil {
			panic(err.Error())
		}

		row, err := db.Query("SELECT * FROM achats")

		if err != nil {
			panic(err.Error())
		}

		for row.Next() {
			var usr_id int
			var usr_name string
			var usr_username string
			var usr_email string
			var usr_phone string
			var usr_address string

			err = row.Scan(&usr_id, &usr_name, &usr_username, &usr_email, &usr_phone, &usr_address)

			if err != nil {
				panic(err.Error())
			}

			// push to array
			users = append(users, User{usr_id, usr_name, usr_username, usr_email, usr_phone, usr_address})

			fmt.Println(usr_id, usr_name, usr_username, usr_email, usr_phone, usr_address)
		}

		defer db.Close()
		fmt.Printf("%v", users)

		c.JSON(200, gin.H{
			"message": "hi",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
