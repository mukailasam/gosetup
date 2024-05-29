package main

import (
	"errors"
	"regexp"
)

const (
	expr     = "^[A-Za-z][A-Za-z0-9_-]*[A-Za-z0-9]$"
	badInput = "bad input"
)

func validate(input string) error {
	r := regexp.MustCompile(expr)
	checker := r.MatchString(input)
	if checker {
		return nil
	}
	return errors.New(badInput)
}
