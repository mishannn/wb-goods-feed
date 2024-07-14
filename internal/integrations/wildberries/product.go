package wildberries

type Product struct {
	Time1           int64   `json:"time1"`
	Time2           int64   `json:"time2"`
	Wh              int64   `json:"wh"`
	Dtype           int64   `json:"dtype"`
	Dist            int64   `json:"dist"`
	ID              int64   `json:"id"`
	Root            int64   `json:"root"`
	KindID          int64   `json:"kindId"`
	Brand           string  `json:"brand"`
	BrandID         int64   `json:"brandId"`
	SiteBrandID     int64   `json:"siteBrandId"`
	Colors          []Color `json:"colors"`
	SubjectID       int64   `json:"subjectId"`
	SubjectParentID int64   `json:"subjectParentId"`
	Name            string  `json:"name"`
	Supplier        string  `json:"supplier"`
	SupplierID      int64   `json:"supplierId"`
	SupplierRating  float64 `json:"supplierRating"`
	SupplierFlags   int64   `json:"supplierFlags"`
	Pics            int64   `json:"pics"`
	Rating          int64   `json:"rating"`
	ReviewRating    float64 `json:"reviewRating"`
	Feedbacks       int64   `json:"feedbacks"`
	Volume          int64   `json:"volume"`
	ViewFlags       int64   `json:"viewFlags"`
	Sizes           []Size  `json:"sizes"`
	TotalQuantity   int64   `json:"totalQuantity"`
	Logs            *string `json:"logs,omitempty"`
	Meta            Meta    `json:"meta"`
	PresetType      string  `json:"preset_type"`
	PanelPromoID    *int64  `json:"panelPromoId,omitempty"`
	PromoTextCard   *string `json:"promoTextCard,omitempty"`
	PromoTextCat    *string `json:"promoTextCat,omitempty"`
	FeedbackPoints  *int64  `json:"feedbackPoints,omitempty"`
	IsNew           *bool   `json:"isNew,omitempty"`
}

type Color struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

type Meta struct {
	PresetID int64 `json:"presetId"`
}

type Size struct {
	Name           string `json:"name"`
	OrigName       string `json:"origName"`
	Rank           int64  `json:"rank"`
	OptionID       int64  `json:"optionId"`
	Wh             int64  `json:"wh"`
	Dtype          int64  `json:"dtype"`
	Price          Price  `json:"price"`
	SaleConditions int64  `json:"saleConditions"`
	Payload        string `json:"payload"`
}

type Price struct {
	Basic     int64 `json:"basic"`
	Product   int64 `json:"product"`
	Total     int64 `json:"total"`
	Logistics int64 `json:"logistics"`
	Return    int64 `json:"return"`
}

type Metadata struct {
	Name         string `json:"name"`
	CatalogType  string `json:"catalog_type"`
	CatalogValue string `json:"catalog_value"`
}
