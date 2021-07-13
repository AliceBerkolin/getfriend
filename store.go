package main

import (
	// The sql go library is needed to interact with the database
	"database/sql"
)

// Store will have two methods, to add a new freind, and to get all existing people
type Store interface {
	CreateFriend(friend *Friend) error
	GetFriend() ([]*Friend, error)
}

// `dbStore` struct implements the `Store` interface. Variable `db` takes the pointer
// to the SQL database connection object.
type dbStore struct {
	db *sql.DB
}

// Create a global `store` variable of type `Store` interface. It will be initialized
// in `func main()`.
var store Store

func (store *dbStore) CreateFriend(friend *Friend) error {
	// 'Friend' is a struct which has "nama", "birthday", and "age" attributes.
	// Type SQL query to insert new person into our database.
	// Note: `peopleinfo` is the name of the table within our `peopleDatabase` postgresql database.
	_, err := store.db.Query(
		"INSERT INTO peopleinfo(name,birthday,age) VALUES ($1,$2,$3)",
		friend.Name, friend.Birthday, friend.Age)
	return err
}

func (store *dbStore) GetFriend() ([]*Friend, error) {
	// Query the database for all Friend, and return the result to the `rows` object.
	// Note: `peopleinfo` is the name of the table within our `peopleDatabase`
	rows, err := store.db.Query("SELECT name, birthday, age FROM peopleinfo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create an empty slice of pointers to `Friend` struct. This slice will be returned
	// by this function to its caller.
	friendList := []*Friend{}
	for rows.Next() {
		// For each row returned from the database, create a pointer to a `Friend` struct.
		friend := &Friend{}
		// Populate the `Name`, `Birthday`, and `Age` attributes of the person
		if err := rows.Scan(&friend.Name, &friend.Birthday, &friend.Age); err != nil {
			return nil, err
		}
		// Finally, append the new person to the returned slice, and repeat for the next row
		friendList = append(friendList, friend)
	}
	return friendList, nil
}