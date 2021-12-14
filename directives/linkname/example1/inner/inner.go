package inner

import (
	_ "unsafe"
	"fmt"
)

//go:linkname linka github.com/xiaowuzai/go-source-code/directives/linkname/example1/outer.LinkA
func linka(){
	fmt.Println("example1 linka")
}

