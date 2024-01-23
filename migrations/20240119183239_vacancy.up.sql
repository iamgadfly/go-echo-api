CREATE table vacancies(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL,
    experience varchar(255),
    vacancy_id int DEFAULT 0,
    salary varchar(255),
    remote varchar(255),
    status varchar(255),
    link varchar(255),
    `description` TEXT,
    `skills` TEXT,
    company_name varchar(255),
    company_url varchar(255),
    created_at varchar(255),
    updated_at varchar(255)
)
