package types

type Category struct {
	Id       		string		`json:"id"`
	Name     		string 		`json:"name"`
	MaxAmount		int   		`json:"maxAmount"`
	SpentAmount 	int 		`json:"spentAmount"`
}

func (c Category) GetId() string{
	return c.Id
}