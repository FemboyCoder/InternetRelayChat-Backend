package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	DATABASE_LOCATION = "data/database.sqlite"
)

var (
	database *sql.DB
)

func Init() {
	if database == nil {
		createDatabseFile()
		var err error
		database, err = sql.Open("sqlite3", DATABASE_LOCATION)
		if err != nil {
			log.Println("error opening database: " + err.Error())
		} else {
			log.Println("Opened Connection to Database")
		}
	}

	createUserDataTable()
	createAuthenticationTable()

	AddUser("Royalty")

}

func Close() {
	log.Println("Closing Connection to Database")
	err := database.Close()
	if err != nil {
		log.Println("Error closing database: " + err.Error())
	}
}

func AddUser(username string) {
	{
		statement, err := database.Prepare("insert into USER_DATA(USERNAME) VALUES (?)")
		if err != nil {
			log.Fatalln("Error creating AddUser statement: " + err.Error())
		}
		_, err = statement.Exec(username)
		if err != nil {
			log.Fatalln("Error executing AddUser statement: " + err.Error())
		}
	}
	{
		statement, err := database.Prepare("insert into USER_AUTHENTICATION(USERNAME) VALUES (?)")
		if err != nil {
			log.Fatalln("Error creating AddUser statement: " + err.Error())
		}
		_, err = statement.Exec(username)
		if err != nil {
			log.Fatalln("Error executing AddUser statement: " + err.Error())
		}
	}
}

func createUserDataTable() {
	statement, err := database.Prepare("" +
		"create table if not exists USER_DATA" +
		"( " +
		"id integer primary key autoincrement not null, " +
		"username text not null, " +
		"nickname text not null default '' " +
		")",
	)
	if err != nil {
		log.Fatalln("Error creating createUserDataTable statement: " + err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln("Error executing createUserDataTable statement: " + err.Error())
	}
}

func createAuthenticationTable() {
	statement, err := database.Prepare("" +
		"create table if not exists USER_AUTHENTICATION" +
		"( " +
		"username text primary key not null, " +
		"password text, " +
		"authentication_key text" +
		")",
	)
	if err != nil {
		log.Fatalln("Error creating createAuthenticationTable statement: " + err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln("Error executing createAuthenticationTable statement: " + err.Error())
	}
}

func createDatabseFile() {
	_, err := os.Create(DATABASE_LOCATION)
	if err != nil && err != os.ErrExist {
		log.Println("Error creating databse file: " + err.Error())
	}
}

type UserData struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type AuthenticationData struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	AuthenticationKey string `json:"authentication_key"`
}
