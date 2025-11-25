package gobench

type User struct {
	Id   int    `json:"id" toon:"id"`
	Name string `json:"name" toon:"name"`
	Role string `json:"role"  toon:"role"`
}

type Payload struct {
	Users []User
}
