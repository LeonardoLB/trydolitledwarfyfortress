package main

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Name string
	qtd  int32
}

func addInventory(item Item) {
	dataJson, err := ioutil.ReadFile("inventory.json")
	if err != nil {
		logging("Cant get inventory")
	}
	err = json.Unmarshal(dataJson, &item)
	if err != nil {
		logging("Cant to write in inventory")
	}
}
