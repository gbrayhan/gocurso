package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
type Post struct {
	UserId int `json : "userId"`
	Id int `json : "id"`
	Title string `json : "title"`
	Body string `json : "body"`
}

func main() {
	var cliente = &http.Client{Timeout: 10 * time.Second}
	var url = "https://jsonplaceholder.typicode.com/posts"

	resp, err := cliente.Get(url)
	if err != nil {
		panic(err.Error())
	}
	var usuarios []Post
	err = json.NewDecoder(resp.Body).Decode(&usuarios)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(usuarios[0].Body)

}
