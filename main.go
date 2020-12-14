package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"github.com/nsf/jsondiff"
)

func parseJSONToByteArray(pathFile string) []byte {
	var data interface{}

	jsonFile, _ := os.Open(pathFile)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal([]byte(byteValue), &data)

	output, _ := json.Marshal(data)

	return output
}

func main() {
	// var data interface{}
	// var data2 interface{}

	// jsonFile, _ := os.Open("./test_files/test_case_from_assignment/1.json")
	// jsonFile2, _ := os.Open("./test_files/test_case_from_assignment/2.json")

	// defer jsonFile.Close()
	// defer jsonFile2.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// byteValue2, _ := ioutil.ReadAll(jsonFile2)

	// _ = json.Unmarshal([]byte(byteValue), &data)
	// _ = json.Unmarshal([]byte(byteValue2), &data2)

	byteValue := parseJSONToByteArray("./test_files/test_case_from_assignment/1.json")
	byteValue2 := parseJSONToByteArray("./test_files/test_case_from_assignment/2.json")

	opts := jsondiff.DefaultConsoleOptions()
	diff, text := jsondiff.Compare([]byte(byteValue), []byte(byteValue2), &opts)

	// fmt.Println(data)
	// fmt.Println(data2)
	fmt.Println(text)
	fmt.Println(diff.String())
}