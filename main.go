package main

import (
	"fmt"
	"server/posting"
)

func main() {
	// port := "8080"
	// http.HandleFunc("/user/", create.AccountCreator)
	// http.HandleFunc("/post/", posting.Post)
	// http.HandleFunc("/frnd/", friend.AddFriend)
	// http.ListenAndServe(":"+port, nil)
	fmt.Println(posting.Check("shivang", "roz3x"))
}
