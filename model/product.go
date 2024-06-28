package model

type (
	Product struct {
		Id    int     `json:"id" gorm:"column:id;autoIncrement"`
		Name  string  `json:"name" gorm:"column:name"`
		Price float64 `json:"price" gorm:"column:price"`
	}
	ProductUpdate struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	ProductCreate struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
)
