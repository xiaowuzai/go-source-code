package inner

import (
"fmt"
_ "unsafe"
)

//go:linkname linkb inner.linkb
func linkb(){
	fmt.Println("example2 linkb")
}
