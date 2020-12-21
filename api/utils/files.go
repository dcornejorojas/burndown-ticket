package utils

import (
	"fmt"
	"os"
	"os/exec"
)

//GetStore returns the store number
func GetStore() string {
	storePath := os.Getenv("P_STORE")
	fmt.Println(storePath)
	out, err := exec.Command("sh", "-c", "grep \"STO\" "+storePath+" | awk -F\"=\" '{print $2}'").Output()
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Printf("in all caps: %q\n", out)
	return string(out)
}
