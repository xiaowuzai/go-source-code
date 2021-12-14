package outer

import (
	_ "unsafe"
	_ "github.com/xiaowuzai/go-source-code/directives/linkname/example2/inner"
)

//go:linkname linkb inner.linkb
func linkb()

func LinkB(){
	linkb()
}