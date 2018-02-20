package main

import "fmt"

type AutoGenerated struct {
	Word struct {
		Adj struct {
			Means []string      `json:"means"`
			Tags  []interface{} `json:"tags"`
		} `json:"adj"`

		Verb struct {
			Means []string      `json:"means"`
			Tags  []interface{} `json:"tags"`
		} `json:"verb"`

		SciName string        `json:"sci_name"`
		Spell   string        `json:"spell"`
		Gtags   []string      `json:"gtags"`
		Spell   []interface{} `json:"!spell"`
	} `json:"angry"`
}

func main() {
	fmt.Printf("Hello, नमस्ते")
	w := Word{}
}
