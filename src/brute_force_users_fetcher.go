package main

// Brute Force, single User Paulo
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

// Brute Force, single User Erik
type ErikUserFetcher struct{}

func (euf ErikUserFetcher) ListUsers() ([]User, error) {
	erik := User{
		ID:        1,
		FirstName: "Erik",
		LastName:  "Haight",
		Email:     "erik(at)sojern.com",
	}
	users := []User{erik}
	return users, nil
}

// Reuse Brute Force, single Users Erik/Paulo
type ErikPauloUserFetcher struct{}

func (epuf ErikPauloUserFetcher) ListUsers() ([]User, error) {
	erikFetcher := ErikUserFetcher{}
	erikUsers, err := erikFetcher.ListUsers()
	if err != nil {
		return nil, err
	}
	pauloFetcher := PauloUserFetcher{}
	pauloUsers, err := pauloFetcher.ListUsers()
	if err != nil {
		return nil, err
	}
	users := make([]User, 0)
	users = append(users, erikUsers...)
	users = append(users, pauloUsers...)
	return users, err
}
