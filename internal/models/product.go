package models

type Product struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       int64  `json:"price" db:"price"`
	SalePrice   int64  `json:"sale_price" db:"sale_price"`
	Description string `json:"description" db:"description"`
	Image       string `json:"image" db:"image"`
	ShopId      int    `json:"shop_id" db:"shop_id"`
	Address     string `json:"address" db:"address"`
	Color       string `json:"color" db:"color"`
	Link        string `json:"link" db:"link"`
}

//func NewNullString(str string) sql.NullString {
//	return sql.NullString{String: str, Valid: true}
//}

//func (ns *NullString) MarshalJSON() ([]byte, error) {
//	if !ns.Valid {
//		return []byte("null"), nil
//	}
//	return json.Marshal(ns.String)
//}

type ProductList struct {
	Products *[]Product `json:"products"`
}
