package wb

import (
	"context"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ParseByLink(ctx context.Context, link string) (models.Product, error) {
	id := GetId(link)
	url := "https://card.wb.ru/cards/detail?appType=1&curr=rub&dest=-1257786&regions=80,38,83,4,64,33,68,70,30,40,86,75,69,1,31,66,110,48,22,71,114&spp=35&nm=" + id
	resp, err := sendData(ctx, url)
	if err != nil {
		return models.Product{}, err
	}

	ShopId, _ := strconv.Atoi(id)
	Name := gjson.Get(resp, "data.products.0.name").String()
	SalePrice := gjson.Get(resp, "data.products.0.salePriceU").Int() / 100
	Price := gjson.Get(resp, "data.products.0.priceU").Int() / 100
	Color := gjson.Get(resp, "data.products.0.colors.0.name").String()

	return models.Product{
		Name:      Name,
		SalePrice: SalePrice,
		Price:     Price,
		Color:     Color,
		ShopId:    ShopId,
		Link:      link,
	}, nil
}

func sendData(ctx context.Context, url string) (string, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

func GetId(link string) string {
	raw := strings.Split(link, "/")
	return raw[len(raw)-2]
}
