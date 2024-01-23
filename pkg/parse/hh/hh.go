package hh

import (
	"context"
	"encoding/json"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ResultApi struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Salary     map[string]any    `json:"salary"`
	Experience map[string]string `json:"experience"`
	Remote     map[string]string `json:"schedule"`
	Employment map[string]string `json:"employment"`
	Desc       string            `json:"description"`
	Employer   map[string]any    `json:"employer"`
	CreatedAt  string            `json:"created_at"`
	Status     map[string]string `json:"type"`
}

func Parse(ctx context.Context, link string) (models.Vacancy, error) {
	var res ResultApi
	var salary string
	splitLink := strings.Split(link, "/")
	id := splitLink[4][:strings.IndexAny(splitLink[4], "?")]
	idInt, _ := strconv.Atoi(id)
	resp, err := getDataById(ctx, id)
	if err != nil {
		return models.Vacancy{}, err
	}
	if err := json.Unmarshal([]byte(resp), &res); err != nil {
		return models.Vacancy{}, err
	}

	if res.Salary != nil {
		from := strconv.FormatFloat(res.Salary["from"].(float64), 'f', 0, 64)
		to := strconv.FormatFloat(res.Salary["to"].(float64), 'f', 0, 64)
		salary = from + "-" + to + " " + res.Salary["currency"].(string)
	}

	return models.Vacancy{
		Name:        res.Name,
		VacancyId:   idInt,
		Experience:  res.Experience["name"],
		Remote:      res.Remote["name"],
		Salary:      salary,
		Status:      res.Status["name"],
		Link:        link,
		Desc:        res.Desc,
		CompanyName: res.Employer["name"].(string),
		CompanyUrl:  res.Employer["alternate_url"].(string),
		CreatedAt:   res.CreatedAt,
	}, nil
}

func getDataById(ctx context.Context, id string) (string, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.hh.ru/vacancies/"+id, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	return string(body), nil
}
