package ozon

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"strings"
)

func ParseByLink(link string) (models.Product, error) {
	url := "https://www.ozon.ru/api/composer-api.bx/page/json/v2?url=" + link

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var data string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &data, chromedp.ByQuery),
	); err != nil {
		return models.Product{}, err
	}
	fmt.Println(data)
	return models.Product{}, nil

	return models.Product{}, nil
}

func GetId(link string) string {
	firstSplit := strings.Split(link, "/")
	words := strings.Split(firstSplit[len(firstSplit)-2], "-")
	return words[len(words)-1]
}
