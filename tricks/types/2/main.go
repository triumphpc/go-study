package main

import "fmt"

type tokenForCertIDRequestArgs struct {
	CertID      string
	CallBackURL string
}

func main() {

	var user map[string]string

	user2 := make(map[string]string)

	fmt.Println(user)
	fmt.Println(user2)

	token := new(tokenForCertIDRequestArgs)

	//usersMap :=  map[string]

	//var args interface{}

	//args = res()
	//res :=  &args

	//fmt.Printf("%T", args)
	//var opts *tokenForCertIDRequestArgs
	//var ok bool
	//opts, ok = args.(tokenForCertIDRequestArgs)

	//fmt.Println(opts, ok)

	//res := res()
	//
	//result := res.(*tokenForCertIDRequestArgs)
	//
	//fmt.Println(result)
	//fmt.Println(result.CallBackURL == "")
}

func res() interface{} {
	return &tokenForCertIDRequestArgs{
		CertID: "fddf",
		//CallBackURL: "vvvv",
	}

}
