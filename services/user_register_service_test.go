package services

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	s := "我12"
	reg := regexp.MustCompile(`^.{1,2}$`)
	res := reg.FindString(s)
	fmt.Println(res=="")
}