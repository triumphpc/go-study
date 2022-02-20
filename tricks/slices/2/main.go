package main

import "fmt"

type user struct {
	likes int
}

func main() {
	users := make([]user, 3)

	shareUser := &users[1]
	shareUser.likes++

	fmt.Println("****************")
	for i := range users {
		fmt.Printf("Users %d Likes: %d\n", i, users[i].likes)
	}

	users = append(users, user{})

	fmt.Println("****************")
	for i := range users {
		fmt.Printf("Users %d Likes: %d\n", i, users[i].likes)
	}

}
