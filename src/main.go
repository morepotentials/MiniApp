package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func newMockFetcher () UsersFetcher{
	return JSONRefresherUsersFetcher{
		filename: "testdata/users.json", 
	}
}

func main() {  //flags mock, db
	host := ":8080"
	MyService := Service{
		Host:    host,
		Fetcher: newDatabaseUsersFetcher(),
	}

	http.HandleFunc("/users", MyService.HandleListUsersEndpoint)
	MyService.start()
}

type Service struct {
	Host    string
	Fetcher UsersFetcher
}

func (S Service) start() {
	fmt.Printf("starting service on port %s\n", S.Host)
	http.ListenAndServe(S.Host, nil)
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (S Service) HandleListUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	users, err := S.Fetcher.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hello")
}
