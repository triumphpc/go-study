package main

import "fmt"

type MyType struct {
}

func main() {

	m := new(MyType)
	m.testCache()
	m.testCache()
	m.testCache()
	m.testCache()
	m.testCache()

}

var localCache map[string]map[string]uint

func (r *MyType) testCache() {
	if localCache == nil {
		fmt.Println("her")
		localCache = make(map[string]map[string]uint, 2)
	}

	fmt.Println(localCache)

	res := re()
	localCache["test"] = res

	res2 := make(map[string]uint, 2)
	localCache["test2"] = res2

}

func re() map[string]uint {

	res := make(map[string]uint, 2)
	res["fddfdf"] = 3

	return res

}
