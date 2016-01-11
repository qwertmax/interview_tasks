package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Database struct {
	DB *gorm.DB
}

type DB struct {
	UserName string `json:"user_name"`
	Password string `json:"passsword"`
	Name     string `json:"name"`
	SSLMode  string `json:"ssh_mode"`
	Address  string `json:"address"`
	Port     string `json:"port"`
}

type Config struct {
	DB      DB     `json:"db"`
	AppPort string `json:"app_port"`
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
	dbconn :=
		"user=" + conf.DB.UserName +
			" password=" + conf.DB.Password +
			" dbname=" + conf.DB.Name +
			" sslmode=" + conf.DB.SSLMode +
			" host=" + conf.DB.Address +
			" port=" + conf.DB.Port
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
		AppPort: "3000",
	}

	db := Database{}
	dbHandler, err := db.getDb(conf)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	dbHandler.DB().Ping()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("qqq"))
	})

	http.HandleFunc("/conf", func(w http.ResponseWriter, r *http.Request) {
		js, err := json.Marshal(conf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	err = http.ListenAndServe(":"+conf.AppPort, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
