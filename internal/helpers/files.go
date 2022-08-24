package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadCharArr() (error, []string) {
	jsonFile, err := os.Open("./resources/charArr.json")
	var arr []string

	if err != nil {
		return err, arr
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal([]byte(byteValue), &arr)

	defer jsonFile.Close()
	return err, arr
}
