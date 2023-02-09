package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var db *sql.DB

type employess struct {
	ID        string
	FirstName string
	LastName  string
}

func init() {
	var err error
	//sql.Open("drivername" , "database://username:password@uri/databaseName?option")
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost/myGolangDatabase?sslmode=disable")
	if err != nil {
		log.Panic(err.Error())
	}
	// defer db.Close()
	err = db.Ping() // is connection alive
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("connect to DB")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(1000 * time.Millisecond)
	fmt.Printf("%+v \n", db.Stats())

}
func main() {
	r := httprouter.New()
	r.POST("/showEmployees", httprouter.Handle(ShowEmployees))
	r.GET("/showEmployeesbyID/:id", httprouter.Handle(ShowEmployeesID))
	r.POST("/createEmployeess", httprouter.Handle(createEmployeess))
	http.ListenAndServe(":8080", r)
}

func ShowEmployees(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query("select * From employess")
	if err != nil {
		log.Panic(err.Error())
	}

	defer rows.Close()
	emps := make([]employess, 0)
	for rows.Next() { //scan all
		emp := employess{}
		err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName)
		if err != nil {
			log.Panic(err.Error())
		}
		emps = append(emps, emp)
	}
	for i, v := range emps {
		fmt.Println(i, v.ID, v.FirstName, v.LastName)
	}
	fmt.Fprintln(w, emps)
	// fmt.Println("%V", emps)
	fmt.Println("%S", db.Stats())

}

func ShowEmployeesID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	fmt.Println(id)
	if id == "" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	row := db.QueryRow("select * From employess where id =$1", id)
	e := employess{}
	err := row.Scan(&e.ID, &e.FirstName, &e.LastName)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %s\n", id)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	}

	fmt.Fprintln(w, e)

	fmt.Println("%S", db.Stats())

}

func createEmployeess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	e := employess{}
	// e.ID = r.FormValue("id")
	e.LastName = r.FormValue("LastName")
	e.FirstName = r.FormValue("FirstName")
	if e.LastName == "" || e.FirstName == "" {
		http.Error(w, "406", http.StatusNotAcceptable)
		return
	}

	res, err := db.Exec("INSERT  INTO employess VALUES (nextval('employess_id_seq'),$1,$2)", e.LastName, e.FirstName)
	if err != nil {
		log.Panic(err.Error())
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v \n", res)

}
