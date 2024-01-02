package ozon

import (
	"fmt"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func ParseByLink(link string) (models.Product, error) {
	//url := "https://www.ozon.ru/api/composer-api.bx/page/json/v2?url=" + link
	//
	//ctx, cancel := chromedp.NewContext(context.Background())
	//defer cancel()
	//
	//var data string
	//if err := chromedp.Run(ctx,
	//	chromedp.Navigate(url),
	//	chromedp.OuterHTML("html", &data, chromedp.ByQuery),
	//); err != nil {
	//	return models.Product{}, err
	//}
	//fmt.Println(data)
	//return models.Product{}, nil
	//
	return models.Product{}, nil
}

func Parse(link string) (models.Product, error) {
	id := GetId(link)

	//url := "https://api.ozon.ru/composer-api.bx/page/json/v2?url=/products/" + id
	url := "https://api.ozon.ru/composer-api.bx/page/json/v2?url=/products/" + id + "/?advert%3Dh9FGxsyoWK_Ydrh90mU6gYwDqbPcXox_y9zanSsmAsDc1wuOepPqesQSuf0v_5uEtMzeVKX2uLW7OFpfxsqDJrZ3xmAAD7I_7WwUe2RgiSiroMN65AKUy9pyeZ1FR99z5B7cd0Ieenq-UX9wRJFU7WACH1KajEd9cOz-nrdZhzZxHNl5pwQ3xDHias7CnNghp9n98_j8bsOjS_0sv86yCMUuX57-hYFwOSI7bbpusErvbIMorx5uUtbA6hqgjZzxEzZzb2v1k4CJiBNp-Oppooo9LmULqJDWogCGcWaKzPAJ6btdLZ0YuUjn9vrHkZJmSIJ5mj_vV27E4sJoHcxuCX7UeFz7HK0%26avtc%3D1%26avte%3D2%26avts%3D1694673786%26keywords%3DIPhone%2B14%26layout_container%3DappPdpPage3%26layout_page_index%3D3%26sh%3D-y9PwlqrqA"
	response, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	}

	// read response body
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	response.Body.Close()

	fmt.Println(string(body))

	return models.Product{}, nil

}

func GetId(link string) string {
	firstSplit := strings.Split(link, "/")
	words := strings.Split(firstSplit[len(firstSplit)-2], "-")
	return words[len(words)-1]
}
