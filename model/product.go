package model

type Product struct {
	ID    int     `json:"Id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
