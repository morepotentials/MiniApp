package main

import (
	"encoding/json"
	"io/ioutil"
)


type JSONRefresherUsersFetcher struct {
	filename string
}



func (jruf JSONRefresherUsersFetcher) ListUsers() ([]User, error) {
	buffer, err := ioutil.ReadFile(jruf.filename)
	if err != nil {return nil, err}
	jsonUsers := make([]User, 0)
	err = json.Unmarshal(buffer, &jsonUsers)
	if err != nil {return nil, err}
	return jsonUsers, nil
}

func NewAnyUsersFetcherFromJSONFile(filename string) AnyUsersFetcher {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	jsonUsers := make([]User, 0)
	err = json.Unmarshal(buffer, &jsonUsers)
	if err != nil {
		panic(err)
	}
	return AnyUsersFetcher{
		users: jsonUsers,
	}
}
