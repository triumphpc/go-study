// -------------------------------
// Declaring Fields, NOT embedding
// -------------------------------

package main

import (
	"fmt"

	"github.com/triumphpc/go-study/tricks/embedding/1/pack"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method notifies users
// of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

// admin represents an admin user with privileges.
// person user is not embedding. All we do here just create a person field based on that other
// concrete type named user.
type admin struct {
	person user // NOT Embedding
	level  string
}

type extPack struct {
	pack.Pack
	level string
}

func main() {
	// Create an admin user using struct literal.
	// Since person also has struct type, we use another literal to initialize it.
	ad := admin{
		person: user{
			name:  "Hoanh An",
			email: "hoanhan101@gmail.com",
		},
		level: "superuser",
	}

	// We call notify through the person field through the admin type value.
	ad.person.notify()

	adPack := extPack{
		pack.Pack{
			Name: "vvv",
		},
		"ааа",
	}
}
