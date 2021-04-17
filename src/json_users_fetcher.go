package main

import (
	"encoding/json"
	"io/ioutil"
)

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
