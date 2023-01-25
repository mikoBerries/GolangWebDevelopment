/*
Basic CRUD on RDBMS mysql localhost
*/package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//connect DB and ping it
	//usermame : password@tcp(localhost:5555)/dbname?charset=utf8
	db, err := sql.Open("mysql",
		// "awsuser:mypassword@tcp(mydbinstance.cakwl95bxza0.us-west-1.rds.amazonaws.com:3306)/test02?charset=utf8")
		"root:mysqlpassword@tcp(localhost:3306)/mydb?charset=utf8")

	// or using sql.OpenDB()
	checkErr(err)
	defer db.Close()

	err = db.Ping() // pinging conection to db
	checkErr(err)
	fmt.Println("connected to localhost")
	//Create table on db
	// {
	// 	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	// 	checkErr(err)
	// 	defer stmt.Close()

	// 	r, err := stmt.Exec()
	// 	checkErr(err)

	// 	n, err := r.RowsAffected()
	// 	checkErr(err)

	// 	fmt.Println("CREATED TABLE customer", n)
	// }
	//simple select DB
	{
		rows, err := db.Query(`SELECT aName FROM mydb.user;`)
		checkErr(err)
		defer rows.Close()

		// data to be used in query
		var s, name string
		s = "RETRIEVED RECORDS:\n"

		// query
		for rows.Next() {
			err = rows.Scan(&name)
			checkErr(err)
			s += name + "\n"
		}
		fmt.Println(s)
	}
	//simple insert query
	{
		stmt, err := db.Prepare(`INSERT INTO mydb.user (aName) VALUES ("James");`)
		checkErr(err)
		defer stmt.Close()

		r, err := stmt.Exec()
		checkErr(err)

		n, err := r.RowsAffected()
		checkErr(err)
		fmt.Println("INSERTED RECORD", n)
	}
	//simple query select * from db
	{
		rows, err := db.Query(`SELECT * FROM user;`)
		checkErr(err)
		defer rows.Close()

		var name string
		var id int
		for rows.Next() {
			err = rows.Scan(&id, &name)
			checkErr(err)
			fmt.Println("RETRIEVED RECORD:", id, "-", name)
		}
	}
	//simple update query
	{
		stmt, err := db.Prepare(`UPDATE user SET aName="Jimmy" WHERE aName="James";`)
		checkErr(err)
		defer stmt.Close()

		r, err := stmt.Exec()
		checkErr(err)

		n, err := r.RowsAffected()
		checkErr(err)

		fmt.Println("UPDATED RECORD", n)
	}
	// simple delete query db
	{
		stmt, err := db.Prepare(`DELETE FROM user WHERE aName="Jimmy";`)
		checkErr(err)
		defer stmt.Close()

		r, err := stmt.Exec()
		checkErr(err)

		n, err := r.RowsAffected()
		checkErr(err)

		fmt.Println("DELETED RECORD", n)
	}
	//simple drop db
	// {
	// 	stmt, err := db.Prepare(`DROP TABLE customer;`)
	// 	checkErr(err)
	// 	defer stmt.Close()

	// 	_, err = stmt.Exec()
	// 	checkErr(err)

	// 	fmt.Println("DROPPED TABLE customer")

	// }
	//similar to java db conection
}

func checkErr(e error) {
	if e != nil {
		fmt.Print(e.Error())
	}
}
