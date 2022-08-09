package entity

// Entity used for name
type NameEntity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Dummy data for NameEntity
var NameDataDummy = []NameEntity{
	{
		Id:   1,
		Name: "A",
	},
	{
		Id:   2,
		Name: "B",
	},
	{
		Id:   3,
		Name: "C",
	},
	{
		Id:   4,
		Name: "D",
	},
}
