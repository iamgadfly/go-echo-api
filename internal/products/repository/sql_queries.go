package repository

const (
	SearchByShopId = `SELECT * FROM products WHERE shop_id=?`
	CreateProduct  = `INSERT INTO products (name,price,sale_price,color,shop_id, link, description) VALUES (:name,:price,:sale_price,:color,:shop_id, :link, :description)`
	Update         = `UPDATE products SET price=:price,sale_price=:sale_price WHERE shop_id=:shop_id`
	GetProducts    = `SELECT id, name, price, sale_price, IFNULL(color, "") as color, IFNULL(link, "") as link from products`
)
