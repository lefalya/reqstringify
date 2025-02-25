package reqstringify

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Username string   `json:"username"`
	Fullname string   `json:"fullname"`
	Packages []string `json:"packages"`
}

func TestStringify(t *testing.T) {

	testStruct := TestStruct{
		Username: "alpha1",
		Fullname: "Alpha One",
		Packages: []string{"alpha1", "alpha2"},
	}

	fmt.Println(Stringify(testStruct))
}
