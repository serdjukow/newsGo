package main

import (
	//"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/",  index)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}

// type User struct {
// 	Name string `json:"name"`
// 	Age  uint16 `json:"name"`
// }

// func main() {
// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	// insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('bob', 35)")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// defer insert.Close()

// 	res, err := db.Query("SELECT `name`, `age` FROM `users`")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for res.Next() {
// 		var user User
// 		err = res.Scan(&user.Name, &user.Age)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
// 	}

// }
