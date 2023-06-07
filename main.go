package main

import (
	"database/sql"
	. "fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id        int
	name      string
	password  string
	hak_akses int
}

func connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/listrik_ku")
	// sql.Open(param1, param2)
	/*
		param1 = type of database
		param2 = address of database, template = user:password@tcp(host:port)/database_name
		ex: root@tcp(127.0.0.1:3306)/listrik_ku
			username: root
			password:
			host: 127.0.0.1 or localhost
			port: 3306
			database neme: listrik_ku
	*/
	if error != nil {
		return nil, error
	}
	return db, nil
}

func sqlQuery() {
	var db, err1 = connect()
	if err1 != nil {
		Println(err1.Error())
		return
	}
	defer db.Close()

	// var id = 2
	var query, err2 = db.Query("SELECT * FROM users")
	if err2 != nil {
		Println(err2.Error())
		return
	}
	defer query.Close()

	var users = []User{}

	for query.Next() {
		var user = User{}
		var err = query.Scan(&user.id, &user.name, &user.password, &user.hak_akses)
		if err != nil {
			Println(err.Error())
			return
		}
		users = append(users, user)
	}
	if err3 := query.Err(); err3 != nil {
		Println(err3.Error())
		return
	}

	for _, user := range users {
		Println(user.id, user.name, "|    Level:", user.hak_akses)
	}
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		Println(err.Error())
		return
	}
	defer db.Close()

	var user = User{}
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", 2).Scan(&user.id, &user.name, &user.password, &user.hak_akses)
	if err != nil {
		Println(err.Error())
		return
	}

	Println(user.id, user.name, "|   Level:", user.hak_akses)
}

func prepare() {
	var db, err = connect()
	if err != nil {
		Println(err.Error())
		return
	}
	defer db.Close()

	var statement, err2 = db.Prepare("SELECT * FROM users WHERE id = ?")
	if err2 != nil {
		Println(err2.Error())
		return
	}

	var user1 = User{}
	statement.QueryRow("1").Scan(&user1.id, &user1.name, &user1.password, &user1.hak_akses)
	var user2 = User{}
	statement.QueryRow("2").Scan(&user2.id, &user2.name, &user2.password, &user2.hak_akses)
	var user3 = User{}
	statement.QueryRow("5").Scan(&user3.id, &user3.name, &user3.password, &user3.hak_akses)

	Println(user1.id, user1.name)
	Println(user2.id, user2.name)
	Println(user3.id, user3.name)
}

func sqlInsert(name string, password string, hak_akses int) {
	var db, err = connect()
	if err != nil {
		Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO users VALUES (null, ?, ?, ?)", name, password, hak_akses)
	if err != nil {
		Println(err.Error())
		return
	}

	Println("Insert data berhasil")
}

func sqlUpdate(username string, password string) {
	var db, err = connect()
	if err != nil {
		Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("UPDATE users SET password = ? where username = ?", password, username)
	if err != nil {
		Println(err.Error())
		return
	}

	Println("Update password berhasil!")
}

func sqlDelete(username string) {
	var db, err = connect()
	if err != nil {
		Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE username = ?", username)
	if err != nil {
		Println(err.Error())
		return
	}

	Println("Delete user berhasil!")
}

func main() {
	sqlQuery()
	Println("------------------------------------")
	sqlQueryRow()
	Println("------------------------------------")
	prepare()
	Println("------------------------------------")
	sqlInsert("Budhi", "password", 2)
	Println("------------------------------------")
	sqlUpdate("Budhi", "passwordchanged")
	Println("------------------------------------")
	sqlDelete("Budhi")
}
