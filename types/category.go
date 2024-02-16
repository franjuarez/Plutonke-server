package types

type Category struct {
	Id       		uint		`json:"id" gorm:"primaryKey;autoIncrement"`
	Name     		string 		`json:"name"        gorm:"not null"`
	MaxAmount		float32		`json:"maxAmount"   gorm:"not null"`
	SpentAmount 	float32		`json:"spentAmount" gorm:"not null"`
}

func (c Category) GetId() uint{
	return c.Id
}


