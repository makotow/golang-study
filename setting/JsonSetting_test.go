package setting

import (
	"testing"
	"fmt"
)

func TestJsonParse(t *testing.T) {
	conf, err :=  Parse("setting.json")
	if err != nil {
		t.Error("error: loading json file." , err)
	}

	actual := conf.Id
	expected := 1
	if actual != expected {
		t.Errorf("got %v\n wnat %v", actual, expected)
	}
	fmt.Println(actual)
}
