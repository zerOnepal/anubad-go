package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("normal Marshal op only")
	os.Stdout.Write(b)
	fmt.Println("\n\nMarshall Indent is good")
	b, _ = json.MarshalIndent(group, "", " ")
	os.Stdout.Write(b)

}
