package wb

import (
	"context"
	"encoding/json"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ProductCat struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Cat  string `json:"cat"`
}

type Product struct {
	Name      string              `json:"name"`
	Price     int                 `json:"priceU"`
	SalePrice int                 `json:"salePriceU"`
	Color     []map[string]string `json:"colors"`
}

type Data struct {
	Prods []Product `json:"products"`
}

type RawData struct {
	Data Data `json:"data"`
}

type DataCat struct {
	Data struct {
		Prods []map[string]any `json:"products"`
	} `json:"data"`
}

const (
	wbUrlCatParse = `https://card.wb.ru/cards/detail?appType=1&curr=rub&dest=-1257786&regions=80,38,83,4,64,33,68,70,30,40,86,75,69,1,31,66,110,48,22,71,114&spp=35&nm=`
)

func ParseByLink(ctx context.Context, link string) (models.Product, error) {
	var data RawData
	id := GetId(link)
	resp, err := SendData(ctx, wbUrlCatParse+id)
	if err != nil {
		return models.Product{}, err
	}

	json.Unmarshal([]byte(resp), &data)
	ShopId, _ := strconv.Atoi(id)

	return models.Product{
		Name:      data.Data.Prods[0].Name,
		SalePrice: int64(data.Data.Prods[0].SalePrice / 100),
		Price:     int64(data.Data.Prods[0].Price / 100),
		Color:     data.Data.Prods[0].Color[0]["name"],
		ShopId:    ShopId,
		Link:      link,
	}, nil
}

func ParseProducts(ctx context.Context, link string, prods chan []models.Product, er chan error) {
	var products []models.Product
	var prodRaw DataCat

	body, err := SendData(ctx, link)
	if err != nil {
		er <- err
	}
	err = json.Unmarshal([]byte(body), &prodRaw)
	if err != nil {
		er <- err
	}

	for _, v := range prodRaw.Data.Prods {
		var color string
		if len(v["colors"].([]interface{})) != 0 {
			color = v["colors"].([]interface{})[0].(map[string]any)["name"].(string)
		}

		prod := models.Product{
			Name:      v["name"].(string),
			SalePrice: int64(v["salePriceU"].(float64) / 100),
			Price:     int64(v["priceU"].(float64) / 100),
			Color:     color,
			Link:      "https://www.wildberries.ru/catalog/" + strconv.Itoa(int(v["id"].(float64))) + "/detail.aspx",
		}
		products = append(products, prod)
	}
	er <- nil
	prods <- products
}

func SendData(ctx context.Context, url string) (string, error) {
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

func getCategories(body string) []ProductCat {
	var categories []ProductCat
	i := 0
	for {
		id := gjson.Get(body, strconv.Itoa(i)+".id").String()
		name := gjson.Get(body, strconv.Itoa(i)+".name").String()
		cat := gjson.Get(body, strconv.Itoa(i)+".query").String()
		categories = append(categories, ProductCat{
			Cat:  cat,
			Id:   id,
			Name: name,
		})

		if name == "" {
			break
		}
		i++
	}
	return categories
}
