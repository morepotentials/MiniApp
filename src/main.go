package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newMockFetcher() JSONRefresherUsersFetcher {
	return JSONRefresherUsersFetcher{
		filename: "testdata/users.json",
	}
}
//localhost:8080/clients?broker_id="me"
//localhost:8080/brokers/me/clients
func main() { 
	r := mux.NewRouter()
	portPtr := flag.Int("port", 8080, "port to listen on")
	isMockPtr := flag.Bool("mock", false, "set to true to enable mockdata")
	flag.Parse()
	
	host := fmt.Sprintf(":%d", *portPtr)

	MyService := Service{
		Host: host,
	}

	if *isMockPtr{
		MyService.Fetcher = newMockFetcher()
		fmt.Println("Successfully initialized mock fetcher")
	} else {
		MyService.Fetcher = newDatabaseUsersFetcher()
		fmt.Println("Successfully initialized database fetcher")
	}

	r.HandleFunc("/broker/{me}/clients", MyService.HandleListClientsEndpoint)
	r.HandleFunc("/users", MyService.HandleListUsersEndpoint)
	http.Handle("/", r)
	MyService.start()
}
//handle func is going to extract the varibles

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
	Started		bool	 `json:"started"`
}

func (S Service) HandleListClientsEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println("me = ", vars["me"])
	broker_id := vars["me"]

	users, err := S.Fetcher.ListBrokerClients(broker_id)
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
