package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age: 30,
		Married: true,
		Hobbies: []string{"Game", "Reading", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":30,"Married":true,"Hobbies":["Game","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age: 30,
		Married: true,
		Hobbies: []string{"Game", "Reading", "Coding"},
		Addresses: []Address{
			{
				Street: "Jl Belum ada",
				City: "Jakarta",
				PostalCode: "999",
			},
			{
				Street: "Jl sudah ada",
				City: "Bandung",
				PostalCode: "888",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":30,"Married":true,"Hobbies":["Game","Reading","Coding"],"Addresses":[{"Street":"Jl Belum ada","City":"Jakarta","PostalCode":"999"},{"Street":"Jl sudah ada","City":"Bandung","PostalCode":"888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}