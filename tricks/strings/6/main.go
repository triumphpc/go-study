package main

import (
	"fmt"
	"strings"
)

func main() {

	trusteeIDs := make([]string, 0, 2)

	trusteeIDs = append(trusteeIDs, "YEST")
	trusteeIDs = append(trusteeIDs, "FDMmmm")

	for idx := range trusteeIDs {
		trusteeIDs[idx] = strings.ToLower(trusteeIDs[idx])
	}

	fmt.Println(trusteeIDs)

}
