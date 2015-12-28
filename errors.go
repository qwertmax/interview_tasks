package main

import (
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Database struct {
	DB *gorm.DB
}

type DB struct {
	UserName string
	Password string
	Name     string
	SSLMode  string
	Address  string
	Port     string
}

type Config struct {
	DB      DB
	AppPort string
}

type Error struct {
	error
	err  string
	code int32
}

func (e Error) Error() string {
	return e.err
}

func (db *Database) getDb(conf Config) (gorm.DB, error) {
	dbconn := "user=" + conf.DB.UserName + " password=" + conf.DB.Password + " dbname=" + conf.DB.Name + " sslmode=" + conf.DB.SSLMode + " host=" + conf.DB.Address + " port=" + conf.DB.Port
	// fmt.Printf("dbconn = %s\n", dbconn)
	return gorm.Open("postgres", dbconn)
}

func main() {
	conf := Config{
		DB: DB{
			UserName: "postgres",
			Password: "1",
			Name:     "upwork",
			SSLMode:  "disable",
			Address:  "192.168.99.100",
			Port:     "5432",
		},
	}

	db := Database{}
	dbHandler, err := db.getDb(conf)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	dbHandler.DB().Ping()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hey")
	})

	err = http.ListenAndServe(":"+conf.AppPort, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
