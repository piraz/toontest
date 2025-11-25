package main

import (
	"fmt"

	"github.com/piraz/toonbench/gobench"
	"github.com/toon-format/toon-go"
)

func main() {
	in := gobench.Payload{
		Users: []gobench.User{
			{Id: 1, Name: "Livi", Role: "admin"},
			{Id: 2, Name: "Sukinho", Role: "root mermão!!!"},
		},
	}

	encoded, err := toon.Marshal(in, toon.WithLengthMarkers(true))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))

	var out gobench.Payload
	if err := toon.Unmarshal(encoded, &out); err != nil {
		panic(err)
	}
	fmt.Printf("first user: %+v\n", out.Users[0])

}
