package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

func LowerCaseString(str string) string {
	a := []rune(str)
	a[0] = unicode.ToLower(a[0])
	lcStr := string(a)
	return lcStr
}

func ParseStringToBool(strVal interface{}) (*bool, error) {

	// since strVal is technically an interface, we check to see if it's actually a string
	v := reflect.ValueOf(strVal)
	typeOfStrval := v.Type()

	if typeOfStrval.Name() != "string" {
		err := errors.New("parsestringtobool util must be a string")
		return nil, err
	}

	// since strconv.ParseBool must take a string, we convert the interface to a string
	strValConverted := fmt.Sprintf("%v", strVal)

	// do conversion, return vals
	boolVal, err := strconv.ParseBool(strValConverted)
	if err != nil {
		fmt.Println("error in ParseStringToBool util: %w", err)
		return nil, err
	}
	return &boolVal, nil

}
