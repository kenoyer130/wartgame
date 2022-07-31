package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kenoyer130/wartgame/models"
)



func UnmarshalJson(item models.JsonSerializable, path string) {
	file, _ := ioutil.ReadFile(path)
	err := json.Unmarshal([]byte(file), &item)

	if err != nil {
		fmt.Printf("json unmarshalling error when parsing %s: %s",path, err)
		panic("unable to parse json")
	}
} 