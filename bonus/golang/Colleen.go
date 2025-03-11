package main

/*
A Golang quine
*/

import "fmt"

func function(s string) {
	fmt.Printf(s, 10, 34, 9, s)
}

func main() {
	/*
		A comment
	*/
	s := "package main%[1]c%[1]c/*%[1]cA Golang quine%[1]c*/%[1]c%[1]cimport %[2]cfmt%[2]c%[1]c%[1]cfunc function(s string) {%[1]c%[3]cfmt.Printf(s, 10, 34, 9, s)%[1]c}%[1]c%[1]cfunc main() {%[1]c%[3]c/*%[1]c%[3]c%[3]cA comment%[1]c%[3]c*/%[1]c%[3]cs := %[2]c%[4]s%[2]c%[1]c%[3]cfunction(s)%[1]c}%[1]c"
	function(s)
}
