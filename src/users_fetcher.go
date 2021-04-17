package main

type UsersFetcher interface {
	ListUsers() ([]User, error)
}

type PauloUserFetcher struct{}

func (puf PauloUserFetcher) ListUsers() ([]User, error) {
	paulo := User{
		ID:        0,
		FirstName: "Paulo",
		LastName:  "Engelke",
		Email:     "peengelke(at)gmail.com",
	}
	users := []User{paulo}
	return users, nil
}

// TODO(paulo): add new implementation of UsersFetcher interface.
// should return erik User, similar to the PauloUserFetcher implementation.
// Swap the Service's UserFetcher implementation to use this.

// func main() {
// 	paulo :=
// 	erik := User{
// 		ID:        1,
// 		FirstName: "Erik",
// 		LastName:  "Haight",
// 		Email:     "ehaight(at)sojern.com",
// 	}
// 	users := make([]User, 0)
// 	users = append(users, paulo)
// 	users = append(users, erik)

// 	MyService := Service{
// 		Host:  ":8080",
// 		Users: users,
// 	}
// 	http.HandleFunc("/users", MyService.ListUsers)
// 	MyService.start()
// }
