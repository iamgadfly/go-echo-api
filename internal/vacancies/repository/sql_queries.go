package repository

const (
	Create = `INSERT INTO vacancies (name, vacancy_id, experience, remote, status, link, company_name, company_url, description, salary, created_at, updated_at)
VALUES (:name, :vacancy_id, :experience, :remote, :status, :link, :company_name, :company_url, :description, :salary, :created_at, :updated_at)`

	FindByVacancyId = `SELECT id, name, vacancy_id, experience, remote,  IFNULL(description, "") as description, IFNULL(salary, "") as salary, status, link, company_name, company_url, created_at, updated_at FROM vacancies WHERE vacancy_id=?`
	FindById        = `SELECT id, name, vacancy_id, experience, remote,  IFNULL(description, "") as description, IFNULL(salary, "") as salary, status, link, company_name, company_url, created_at, updated_at FROM vacancies WHERE id=?`
)
