package main

import "fmt"

func main() {
	type AA struct {
		ID int64
	}

	arraySt := []AA{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}

	tmpMap := map[int64]*AA{}

	for _, val := range arraySt {
		tmpMap[val.ID] = &val
	}

	for i := range tmpMap {
		fmt.Printf("%+v \n", *tmpMap[i]) // {ID:3} {ID:3} {ID:3}

	}
}
