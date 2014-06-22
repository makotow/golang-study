package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

// database: http://dev.mysql.com/doc/sakila/en/
func main() {

  // connect database
  db, err := sql.Open("mysql","root@tcp(127.0.0.1:3306)/sakila")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  // create prepared statement
  stmt, err := db.Prepare("select actor_id, first_name from actor where actor_id=?")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  var id int
  var name string

  // where actor_id = 1
  err = stmt.QueryRow(1).Scan(&id, &name)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(id, name)


  // whre actor_id=10
  err = stmt.QueryRow(13).Scan(&id, &name)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(id, name)
}
