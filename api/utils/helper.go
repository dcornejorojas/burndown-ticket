package utils
import (
	"fmt"
	"strconv"
	"bytes"
	"io"
	"ticket/api/models"
)

func IsValidString (str string) bool {
	return (len(str) != 0 && str != "")
}

func IsValidTicket(arr []models.Transaction) bool {
	return (len(arr) > 0)
}

func IsValidInt (str string) bool {
	fmt.Println(len(str))
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

func StreamToBytes(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
		buf.ReadFrom(stream)
		return buf.Bytes()
}