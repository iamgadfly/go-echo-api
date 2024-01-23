package models

type Vacancy struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	VacancyId   int    `json:"vacancy_id" db:"vacancy_id"`
	Experience  string `json:"experience" db:"experience"`
	Remote      string `json:"remote" db:"remote"`
	Salary      string `json:"salary" db:"salary"`
	Status      string `json:"status" db:"status"`
	Link        string `bson:"link" db:"link"`
	Desc        string `json:"desc" db:"description"`
	Skills      string `json:"skills" db:"skills"`
	CompanyName string `json:"company_name" db:"company_name"`
	CompanyUrl  string `json:"company_url" db:"company_url"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}
