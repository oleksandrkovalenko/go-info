package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	PublicRepos int `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	// HTTP call
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//Decode JSON
	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func main() {
	user, err := userInfo("oleksandrkovalenko")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf("%+v\n", user)
}
