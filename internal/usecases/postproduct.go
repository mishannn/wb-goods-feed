package usecases

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"fmt"
	"strings"
	"text/template"

	"github.com/mishannn/wb-goods-feed/internal/entities/feed"
	"github.com/mishannn/wb-goods-feed/internal/entities/product"
)

//go:embed pricehistorychart.tmpl
var chartTemplate string

type HistoryChartData struct {
	Days   []string
	Prices []string
}

func generatePriceHistoryChartLink(priceHistory []product.PriceHistoryItem) (string, error) {
	funcs := template.FuncMap{"join": strings.Join}

	tmpl, err := template.New("pricehistorychart").Funcs(funcs).Parse(chartTemplate)
	if err != nil {
		return "", fmt.Errorf("can't parse template: %w", err)
	}

	days := make([]string, 0, len(priceHistory))
	prices := make([]string, 0, len(priceHistory))
	for _, item := range priceHistory {
		days = append(days, item.Date.Format("02.01"))
		prices = append(prices, fmt.Sprintf("%d", item.Price.RUB/100))
	}

	var mermaidCode bytes.Buffer
	err = tmpl.Execute(&mermaidCode, HistoryChartData{
		Days:   days,
		Prices: prices,
	})
	if err != nil {
		return "", fmt.Errorf("can't execute template: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(mermaidCode.Bytes())
	return fmt.Sprintf("https://mermaid.ink/img/%s", encoded), nil
}

type PostProduct struct {
	fetcher product.Fetcher
	chooser product.Chooser
	poster  feed.Poster
}

func NewPostProduct(fetcher product.Fetcher, chooser product.Chooser, poster feed.Poster) *PostProduct {
	return &PostProduct{
		fetcher: fetcher,
		chooser: chooser,
		poster:  poster,
	}
}

func (pp *PostProduct) Do() error {
	// Выбираем продукт

	products, err := pp.fetcher.GetProducts()
	if err != nil {
		return fmt.Errorf("can't get products: %w", err)
	}

	product, err := pp.chooser.ChooseProduct(products)
	if err != nil {
		return fmt.Errorf("can't choose product: %w", err)
	}

	// Забираем только 5 изображений

	var imagesLimit int
	if len(product.Images) >= 5 {
		imagesLimit = 5
	} else {
		imagesLimit = len(product.Images)
	}

	postImages := make([]feed.Image, 0, len(product.Images))
	for _, productImage := range product.Images[0:imagesLimit] {
		postImages = append(postImages, feed.Image(productImage))
	}

	// Если есть история цены, генерируем ссылку на график цен

	if len(product.PriceHistory) >= 2 {
		chartImageURL, err := generatePriceHistoryChartLink(product.PriceHistory)
		if err != nil {
			return fmt.Errorf("can't generate price history chart: %w", err)
		}
		postImages = append(postImages, feed.Image{
			URL: chartImageURL,
		})
	}

	// Устанавливаем актуальную цену на товар

	productPrice := product.PriceHistory[len(product.PriceHistory)-1].Price.RUB

	// Подготавливаем пост

	content := fmt.Sprintf("Цена 💳 %d руб.\nРейтинг ⭐️ %.1f на 💬 %d отзывов", productPrice/100, product.Rating, product.ReviewCount)

	if len(product.Tags) > 0 {
		tagsRowItems := make([]string, 0, len(product.Tags))
		for _, tag := range product.Tags {
			tagNameWithNoSpaces := strings.Replace(tag.Name, " ", `\_`, -1)
			tagsRowItems = append(tagsRowItems, fmt.Sprintf("#%s", tagNameWithNoSpaces))
		}

		content = fmt.Sprintf("%s\n\n%s", strings.Join(tagsRowItems, " "), content)
	}

	post := feed.Post{
		Title:   fmt.Sprintf("%s от %s", product.Name, product.Brand),
		Content: content,
		Images:  postImages,
		Link:    product.Link,
	}

	// Публикуем пост

	err = pp.poster.PublishPost(post)
	if err != nil {
		return fmt.Errorf("can't publish post: %w", err)
	}

	return nil
}
