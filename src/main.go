package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
		http.HandleFunc("/users", ListUsers)
		http.ListenAndServe(":8080", nil)
}

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func ListUsers(w http.ResponseWriter, r *http.Request){
	paulo := User{
		ID: 0,
		FirstName: "Paulo",
		LastName: "Engelke",
	}
	erik := User{
		ID: 1,
		FirstName: "Erik",
		LastName: "Haight",
	}
	users := make([]User, 0)
	users = append(users, paulo) 
	users = append(users, erik)
	fmt.Println(users)

	bytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return 
	}
	_, err = w.Write(bytes)
	if err !=nil{
		fmt.Println(err)
		return
	}


}
