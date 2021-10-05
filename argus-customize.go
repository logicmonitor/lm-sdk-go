package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open(os.Args[1])
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		fmt.Println("Failed to unmarshal json file")
		return
	}

	m := map[string]interface{}{
		"description": "Too Many Requests",
		"headers": map[string]interface{}{
			"x-rate-limit-limit": map[string]interface{}{
				"type":        "integer",
				"description": "Request limit per X-Rate-Limit-Window",
			},
			"x-rate-limit-remaining": map[string]interface{}{
				"type":        "integer",
				"description": "The number of requests left for the time window",
			},
			"x-rate-limit-window": map[string]interface{}{
				"type":        "integer",
				"description": "The rolling time window length with the unit of second",
			},
		},
	}
	userAgent := map[string]interface{}{
		"in":       "header",
		"name":     "User-Agent",
		"type":     "string",
		"required": false,
		"default":  fmt.Sprintf("Logicmonitor/SDK: Argus Dist-%s", os.Args[3]),
	}
	patchFields := map[string]interface{}{
		"in":       "query",
		"name":     "PatchFields",
		"type":     "string",
		"required": false,
	}
	paths := result["paths"].(map[string]interface{})
	for _, v := range paths {
		verbsMap := v.(map[string]interface{})
		for key, val := range verbsMap {
			verbDetails := val.(map[string]interface{})
			responses := verbDetails["responses"].(map[string]interface{})
			responses["429"] = m
			parameters := verbDetails["parameters"].([]interface{})
			if key == "patch" {
				parameters = append(parameters, patchFields)
			}
			parameters = append(parameters, userAgent)
			verbDetails["parameters"] = parameters
		}
	}
	file, _ := json.MarshalIndent(result, "", "  ")
	err = ioutil.WriteFile(os.Args[2], file, 0644)
	if err != nil {
		fmt.Println("Failed write to file", err)
		return
	}
}
