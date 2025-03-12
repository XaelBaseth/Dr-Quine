package main

import (
	"fmt"
	"os"
	"os/exec"
)

const code = "package main%[1]cimport (%[1]c%[2]c%[3]cfmt%[3]c%[1]c%[2]c%[3]cos%[3]c%[1]c%[2]c%[3]cos/exec%[3]c%[1]c)%[1]c%[1]cconst code = %[3]c%[4]s%[3]c%[1]c%[1]cfunc main() {%[1]c%[2]ci := %[5]d%[1]c%[2]cif i <= 0 {%[1]c%[2]c%[2]creturn%[1]c%[2]c}%[1]c%[1]c%[2]cnewSrc := fmt.Sprintf(%[3]cSully_%%d.go%[3]c, i)%[1]c%[2]cnewExec := fmt.Sprintf(%[3]cSully_%%d%[3]c, i)%[1]c%[2]cfile, _ := os.Create(newSrc)%[1]c%[2]cdefer file.Close()%[1]c%[1]c%[2]cfmt.Fprintf(file, code, 10, 9, 34, code, i-1)%[1]c%[1]c%[2]cexec.Command(%[3]cgo%[3]c, %[3]cbuild%[3]c, %[3]c-o%[3]c, newExec, newSrc).Run()%[1]c%[2]ccmd := exec.Command(%[3]c./%[3]c + newExec)%[1]c%[2]ccmd.Start()%[1]c}%[1]c"

func main() {
	i := 5
	if i <= 0 {
		return
	}

	newSrc := fmt.Sprintf("Sully_%d.go", i)
	newExec := fmt.Sprintf("Sully_%d", i)

	file, _ := os.Create(newSrc)
	defer file.Close()

	fmt.Fprintf(file, code, 10, 9, 34, code, i-1)

	exec.Command("go", "build", "-o", newExec, newSrc).Run()
	cmd := exec.Command("./" + newExec)
	cmd.Start()
}
