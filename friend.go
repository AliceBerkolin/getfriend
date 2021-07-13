package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// friend is a struct decsribing its properties
type Friend struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Age      string `json:"age"`
}

func getFriend(w http.ResponseWriter, r *http.Request) {
	// Retrieve people from postgresql database using our `store` interface variable's
	// `func (*dbstore) GetFriend` pointer receiver method defined in `store.go` file
	friendList, err := store.GetFriend()

	// Convert the `friendList` variable to JSON
	friendListBytes, err := json.Marshal(friendList)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write JSON list of friends to response
	w.Write(friendListBytes)
}

func createFriend(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML form data received in the request
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Extract the field information about the friend from the form info
	friend := Friend{}
	friend.Name = r.Form.Get("name")
	friend.Birthday = r.Form.Get("birthday")
	friend.Age = r.Form.Get("age")

	// Write new friend details into postgresql database using our `store` interface variable's
	// `func (*dbstore) Createfriend` pointer receiver method defined in `store.go` file
	err = store.CreateFriend(&friend)
	if err != nil {
		fmt.Println(err)
	}

	//Redirect to the originating HTML page
	http.Redirect(w, r, "/", http.StatusFound)
}