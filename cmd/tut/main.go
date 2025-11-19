package main

import (
	"fmt"

	"github.com/toon-format/toon-go"
)

type User struct {
	Id   int    `json:"id" toon:"id"`
	Name string `json:"name" toon:"name"`
	Role string `json:"role"  toon:"role"`
}

type Payload struct {
	Users []User
}

func main() {
	in := Payload{
		Users: []User{
			{Id: 1, Name: "Livi", Role: "admin"},
			{Id: 2, Name: "Sukinho", Role: "root mermão!!!"},
		},
	}

	encoded, err := toon.Marshal(in, toon.WithLengthMarkers(true))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))

	var out Payload
	if err := toon.Unmarshal(encoded, &out); err != nil {
		panic(err)
	}
	fmt.Printf("first user: %+v\n", out.Users[0])

}
