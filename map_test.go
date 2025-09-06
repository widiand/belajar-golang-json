package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Macbook pro","imageUrl":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["imageUrl"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{} {
		"id": "P0001",
		"name": "Apple Macbook pro",
		"imageUrl": "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}