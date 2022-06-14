package B

import (
	"fmt"
	"github.com/triumphpc/go-study/tricks/packages/1/A"
)

func init() {
	fmt.Println("b2")
}

func f() {
	_ = A.A2
}

var B2 = ""
