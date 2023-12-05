CREATE table users(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    surname varchar(255),
    password varchar(255),
    email varchar(255) UNIQUE,
    balance integer DEFAULT 0,
    is_register_social boolean not null default 0,
    is_paid boolean not null default 0,
    role enum('admin','copywriter', 'customer') DEFAULT 'customer'
)
