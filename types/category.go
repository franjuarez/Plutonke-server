package types

type Category struct {
	Id       		string		`json:"id"`
	Name     		string 		`json:"name"`
	MaxAmount		uint   		`json:"price"`
	SpentAmount 	uint 		`json:"date"`
}