package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Format struct {
	User  int    `json:"userid"`
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in Reading: ", err)
		return
	}

	var post Format
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error Decode: ", err)
		return
	}

	fmt.Println("userid: ", post.User)
	fmt.Println("id: ", post.Id)
	fmt.Println("title: ", post.Title)
	fmt.Println("body: ", post.Body)

}
