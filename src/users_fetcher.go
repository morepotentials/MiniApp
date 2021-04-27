package main

type UsersFetcher interface {
	ListUsers() ([]User, error)
	ListBrokerClients(broker_id string) ([]User, error)
}


// Generic Any Users, can be configured externally.
// Not only a single user, supports any # of users.
/*
type AnyUsersFetcher struct {
	users []User
}

func (auf AnyUsersFetcher) ListUsers() ([]User, error) {
	return auf.users, nil
}

//  Generic Any User. Can be configured externally.
// Instantiated through a global var. Or passed through main.
type AnyUserFetcher struct {
	user User
}

func (auf AnyUserFetcher) ListUsers() ([]User, error) {
	users := []User{auf.user}
	return users, nil
}

//  Instantiation of AnyUserFetcher, supplying Paulo user.
var PauloAnyUserFetcher = AnyUserFetcher{
	user: User{
		ID:        0,
		FirstName: "Paulo",
		LastName:  "Engelke",
		Email:     "peengelke(at)gmail.com",
	},
}

//  Instantiation of AnyUserFetcher, supplying Erik user.
var ErikAnyUserFetcher = AnyUserFetcher{
	user: User{
		ID:        1,
		FirstName: "Erik",
		LastName:  "Haight",
		Email:     "erik(at)sojern.com",
	},
}

*/