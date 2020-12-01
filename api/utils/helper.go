package utils
import (
	"fmt"
)

func IsValidString (str string) bool {
	fmt.Println(len(str))
	return (len(str) != 0 && str != "")
}