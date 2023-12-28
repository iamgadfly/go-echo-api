package wb

import (
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ParseByLink(link string) (models.Product, error) {
	id := GetId(link)
	url := "https://card.wb.ru/cards/detail?appType=1&curr=rub&dest=-1257786&regions=80,38,83,4,64,33,68,70,30,40,86,75,69,1,31,66,110,48,22,71,114&spp=35&nm=" + id
	resp, err := sendData(url)
	if err != nil {
		return models.Product{}, err
	}

	ShopId, _ := strconv.Atoi(id)
	body, _ := ioutil.ReadAll(resp.Body)
	Name := gjson.Get(string(body), "data.products.0.name").String()
	SalePrice := gjson.Get(string(body), "data.products.0.salePriceU").Int() / 100
	Price := gjson.Get(string(body), "data.products.0.priceU").Int() / 100
	Color := gjson.Get(string(body), "data.products.0.colors.0.name").String()

	return models.Product{
		Name:      Name,
		SalePrice: SalePrice,
		Price:     Price,
		Color:     Color,
		ShopId:    ShopId,
		Link:      link,
	}, nil
}

func sendData(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func GetId(link string) string {
	raw := strings.Split(link, "/")
	return raw[len(raw)-2]
}
