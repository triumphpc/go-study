package main

import "log"

import "github.com/ssrathi/go-scrub"

// https://github.com/ssrathi/go-scrub/
// Удаляет конфидициальные поля

func main() {

	// Have a struct with some sensitive fields.
	type testScrub struct {
		Username string
		Password string
		Codes    []string
	}

	// Create a struct with some sensitive data.
	T := testScrub{
		Username: "administrator",
		Password: "my_secret_passphrase",
		Codes:    []string{"pass1", "pass2", "pass3"},
	}

	// Create a set of field names to scrub (default is 'password').
	fieldsToScrub := map[string]bool{
		"password": true,
		"codes":    true,
	}

	// Call the util API to get a JSON formatted string with scrubbed field values.
	out := scrub.Scrub(&T, fieldsToScrub)

	// Log the scrubbed string without worrying about prying eyes!
	log.Println(out) // {"Username":"administrator","Password":"********","Codes":["********","********","********"]}

}
