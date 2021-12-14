package main

import (
	outer1 "github.com/xiaowuzai/go-source-code/directives/linkname/example1/outer"
	outer2 "github.com/xiaowuzai/go-source-code/directives/linkname/example2/outer"
)

func main(){
	outer1.LinkA()
	outer2.LinkB()
}
