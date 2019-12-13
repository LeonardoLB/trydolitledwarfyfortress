package main

import (
	"encoding/json"
	"os"
)

type Item struct {
	Name string `json:name`
	Qtd  int32  `json:qtd`
}

func addInventory(item Item) {
	f, err := os.OpenFile("./inventory.json", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		logging("Cant get inventory")
	}
	defer f.Close()
	b, err := json.Marshal(item)
	f.Write(b)
	f.Close()
}

// func addInventory(item Item) {
// 	dataJson, err := ioutil.ReadFile("inventory.json")
// 	if err != nil {
// 		logging("Cant get inventory")
// 	}

// 	err = json.Unmarshal(dataJson, &item)
// 	if err != nil {
// 		logging("Cant to write in inventory")
// 	}
// }
