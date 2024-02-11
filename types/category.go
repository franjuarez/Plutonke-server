package types

type Category struct {
	Id       		string		`json:"id"`
	Name     		string 		`json:"name"`
	MaxAmount		float32   		`json:"maxAmount"`
	SpentAmount 	float32 		`json:"spentAmount"`
}

func (c Category) GetId() string{
	return c.Id
}