package tools

import (
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func StrToInt(err error, index string) int {
	result, err := strconv.Atoi(index)
	if err != nil {
		HasError(err, "string to int error"+err.Error(), -1)
	}
	return result
}

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		log.Print(err.Error())
		return false, err
	}
	return true, nil
}

// Assert conditional assertion
// trigger panic when the assertion condition is false
// The next code will not be executed for the current request, and the error message and error code in the specified format will be returned
func Assert(condition bool, msg string, code ...int) {
	if !condition {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// HasError error assertion
// trigger panic when error is not nil
// The next code will not be executed for the current request, and the error message and error code in the specified format will be returned
// If msg is empty, it defaults to the content in error
func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg == "" {
			msg = err.Error()
		}
		log.Println(err)
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}
