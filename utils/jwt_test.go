package utils

import (
	"fmt"
	"testing"
)

func TestAnalyseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjI0fQ.BHCcYIi2OOuGklleLE8O9m6hInGR8FHeudAGQVIdDR4"
	userClaims, _ := AnalyseToken(token)
	fmt.Println(userClaims.UserId)
}