package main

import (
	"fmt"
	"os"
)

const filename = "./Grace_kid.go"
const format = "package main%[1]c%[1]cimport (%[3]cfmt%[3]c; %[3]cos%[3]c)%[1]c%[1]cconst filename = %[3]cGrace_kid.go%[3]c%[1]cconst format = %[3]c%[4]s%[3]c%[1]c%[1]cfunc main() {%[1]c%[2]cfile, _ := os.Create(filename)%[1]c%[2]cdefer file.Close()%[1]c%[2]cfmt.Fprintf(file, format, 10, 9, 34, format)%[1]c}%[1]c"

func main() {
	file, _ := os.Create(filename)
	defer file.Close()
	fmt.Fprintf(file, format, 10, 9, 34, format)
}
