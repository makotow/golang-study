package setting

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Config struct {
	Id int
	Name string
	Array []int
}

func Parse(filename string) (Config, error) {
	var c Config
	jsonString, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("error: %v ", err)
		return c, err
	}
	err = json.Unmarshal(jsonString, &c)
	if err != nil {
		fmt.Printf("error: %v ", err)
		return c, err
	}
	return c , nil
}
