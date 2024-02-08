package types

type Expense struct {
	Id       		string 		`json:"id"`
	Name     		string 		`json:"name"`
	Price    		uint   		`json:"price"`
	Date     		string 		`json:"date"`
	Category 		string 		`json:"category"`
}