package types


type Expense struct {
	Id        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Price     float32   `json:"price" gorm:"not null"`
	Date      string    `json:"date" gorm:"not null"`
	CategoryID uint      `json:"category_id" gorm:"foreignKey:Id"`
}

func (e Expense) GetId() uint {
	return e.Id
}

