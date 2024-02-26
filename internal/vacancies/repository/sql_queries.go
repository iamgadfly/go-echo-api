package repository

const (
	Create = `INSERT INTO vacancies (name, vacancy_id, experience, remote, status, link, company_name, company_url, description, salary, created_at, updated_at)
VALUES (:name, :vacancy_id, :experience, :remote, :status, :link, :company_name, :company_url, :description, :salary, :created_at, :updated_at)`

	Update = `UPDATE vacancies SET name=:name,experience=:experience,salary=:salary, status=:status, link=:link, company_name=:company_name, remote=:remote, description=:description, updated_at=:updated_at WHERE vacancy_id=:vacancy_id`

	FindByVacancyId = `SELECT id, name, vacancy_id, experience, remote,  IFNULL(description, "") as description, IFNULL(salary, "") as salary, status, link, company_name, company_url, created_at, updated_at FROM vacancies WHERE vacancy_id=?`
	FindById        = `SELECT id, name, vacancy_id, experience, remote,  IFNULL(description, "") as description, IFNULL(salary, "") as salary, status, link, company_name, company_url, created_at, updated_at FROM vacancies WHERE id=?`
)
