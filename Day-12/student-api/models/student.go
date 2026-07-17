package models

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Students = []Student{
	{
		ID:   1,
		Name: "Satyam",
		Age:  21,
	},
	{
		ID:   2,
		Name: "Rahul",
		Age:  22,
	},
	{
		ID:   3,
		Name: "Krishna",
		Age:  23,
	},
}