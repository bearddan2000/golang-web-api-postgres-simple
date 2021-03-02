package main

import (

  "encoding/json"

  //"log"

  "net/http"

  //"database/sql"

  "fmt"

  "github.com/gorilla/mux"

  "github.com/jinzhu/gorm"

  "github.com/rs/cors"

  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Pop struct {

  gorm.Model

  Name   string

  Color  string
}

func main() {
  //.conn()

  router := mux.NewRouter()

  router.HandleFunc("/", GetPops).Methods("GET")

  handler := cors.Default().Handler(router)

  //log.Fatal(http.ListenAndServe(":8080", handler))

  http.ListenAndServe(":8080", handler)
}

func GetPops(w http.ResponseWriter, r *http.Request) {

  var db *gorm.DB

  const (
    host     = "db"
    port     = 5432
    user     = "maria"
    password = "pass"
    dbname   = "beverage"
  )
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db,err := gorm.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }

  if db != nil {
    db.Debug().DropTableIfExists(&Pop{})
    //Drops table if already exists
    var items = []Pop{

            {Name: "Cola", Color: "carmel"},

            {Name: "Gingerale",  Color: "brown"},

            {Name: "Lime",  Color: "green"},

            {Name: "Cherry",  Color: "red"},

            {Name: "Grape",  Color: "purple"},
        }
    db.Debug().AutoMigrate(&Pop{})
    //Auto create table based on Model
    for index := range items {

          db.Create(&items[index])

      }

    var pops []Pop

    db.Find(&pops)

    db.Close()

    json.NewEncoder(w).Encode(&pops)
  } else {
    fmt.Fprintf(w, "Hello astaxie!")
  }
}
