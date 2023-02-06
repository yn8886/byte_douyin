package services

import (
	"fmt"
	"regexp"
	"testing"
	"unicode/utf8"
)

func TestRegexp(t *testing.T) {
	s := "我12"
	reg := regexp.MustCompile(`^.{1,2}$`)
	res := reg.FindString(s)
	fmt.Println(res=="")
}

func TestUtf8(t *testing.T) {
	s := "我12"
	n := utf8.RuneCountInString(s)
	fmt.Println(n)
}