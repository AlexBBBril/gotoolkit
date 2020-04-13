package env

import (
	"fmt"
	"os"
	"strconv"
)

// Required will return the value of the environment variable 
// or cause a panic if there is no such variable
func Required(name string) string {
	env := os.Getenv(name)
	if "" == env {
		panic(fmt.Sprintf("Missing environment variable named: %s", name))
	}

	return env
}

// RequiredAsInt will return the value of the environment variable. Will try to translate string to int. 
// It will cause a panic if there is no such variable or when receiving an error translating string to int.
func RequiredAsInt(name string) int {
	i, err := strconv.Atoi(Required(name))
	if nil != err {
		panic(err)
	}

	return i
}
