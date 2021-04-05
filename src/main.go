package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	paulo := User{
		ID: 0,
		FirstName: "Paulo",
		LastName: "Engelke",
		Email:"peengelke(at)gmail.com",
	}
	erik := User{
		ID: 1,
		FirstName: "Erik",
		LastName: "Haight",
		Email:"ehaight(at)sojern.com",
	}
	users := make([]User, 0)
	users = append(users, paulo) 
	users = append(users, erik)

	MyService:= Service{
		Host: ":8080",
		Users: users,
	}
	http.HandleFunc("/users", MyService.ListUsers)
	MyService.start()
}

type UsersFetcher interface {
	ListUsers() ([]User, error)
}
type Service struct {
	Host string
	Users []User
	Fetcher UsersFetcher 
}

func (S Service) start(){
	http.ListenAndServe(S.Host, nil)
}

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}

func (S Service) ListUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println(S.Host)
	fmt.Println(S.Users)

	bytes, err := json.Marshal(S.Users)
	if err != nil {
		fmt.Println(err)
		return 
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(bytes)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("hello")

}
