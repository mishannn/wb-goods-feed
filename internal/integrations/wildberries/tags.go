package wildberries

import (
	"fmt"

	"github.com/mishannn/wb-goods-feed/internal/entities/product"
	"github.com/mishannn/wb-goods-feed/internal/shared/httputils"
)

type breadcrumbsDataResponse struct {
	ResultState int64                        `json:"resultState"`
	Value       breadcrumbsDataResponseValue `json:"value"`
}

type breadcrumbsDataResponseValue struct {
	Data breadcrumbsData `json:"data"`
}

type breadcrumbsData struct {
	SitePath []sitePath `json:"sitePath"`
}

type sitePath struct {
	Name string `json:"name"`
}

func getProductTags(wbProduct Product) ([]product.Tag, error) {
	dataURL := fmt.Sprintf("https://www.wildberries.ru/webapi/product/%d/data?subject=%d&kind=%d&brand=%d", wbProduct.ID, wbProduct.SubjectID, wbProduct.KindID, wbProduct.BrandID)

	respBody, err := httputils.HttpGet[breadcrumbsDataResponse](dataURL)
	if err != nil {
		return nil, fmt.Errorf("can't get breadcrumb data: %w", err)
	}

	tags := make([]product.Tag, 0, len(respBody.Value.Data.SitePath))
	for _, sitePathItem := range respBody.Value.Data.SitePath {
		tags = append(tags, product.Tag{
			Name: sitePathItem.Name,
		})
	}

	return tags, nil
}
