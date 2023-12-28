CREATE table products(
                      id BIGINT AUTO_INCREMENT PRIMARY KEY,
                      name varchar(255),
                      description varchar(255),
                      price integer DEFAULT 0,
                      sale_price integer DEFAULT 0,
                      shop_id integer DEFAULT 0,
                      image varchar(255),
                      address varchar(255),
                      color varchar(255),
                      link varchar(255)
)

