package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "paulo"
  dbname   = "postgres"
)

type databaseUsersFetcher struct {
	db *sql.DB
}

func newDatabaseUsersFetcher () databaseUsersFetcher {
	//create db connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"dbname=%s sslmode=disable",
	host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return databaseUsersFetcher{db:db}
}


func (duf databaseUsersFetcher) ListBrokerClients(broker_id string) ([]User, error) {
		//query db
		rows, err := duf.db.Query("SELECT id, first_name, last_name, email FROM users U RIGHT JOIN clients C ON U.id = C.client_id WHERE C.broker_id ="+broker_id+";")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		fmt.Println("You got the rows")
		fmt.Println(rows)
		//translate db records to []User
	
		users := make([]User, 0)
		for rows.Next() {
			user := User{}
			err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
			if err != nil {
				return nil, err
			}
			fmt.Println(user)
			users = append(users, user)
		}
	
		return users, nil
}


func (duf databaseUsersFetcher) ListUsers() ([]User, error) {

	//query db
	rows, err := duf.db.Query("SELECT id,first_name,last_name,email FROM users")
  if err != nil {
    return nil, err
  }
	defer rows.Close()
	fmt.Println("You got the rows")
	
	//translate db records to []User

	users := make([]User, 0)
	for rows.Next() {
    user := User{}
    err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
    if err != nil {
      return nil, err
    }
		fmt.Println(user)
		users = append(users, user)
  }

	return users, nil
}
