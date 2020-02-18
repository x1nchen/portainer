package model

type Extension struct {
	// Extension identifier
	Id int32 `json:"Id,omitempty"`
	// Extension name
	Name string `json:"Name,omitempty"`
	// Is the extension enabled
	Enabled bool `json:"Enabled,omitempty"`
	// Short description about the extension
	ShortDescription string `json:"ShortDescription,omitempty"`
	// URL to the file containing the extension description
	DescriptionURL string `json:"DescriptionURL,omitempty"`
	// Is the extension available for download and activation
	Available bool `json:"Available,omitempty"`
	// List of screenshot URLs
	Images []string `json:"Images,omitempty"`
	// Icon associated to the extension
	Logo string `json:"Logo,omitempty"`
	// Extension price
	Price string `json:"Price,omitempty"`
	// Details about extension pricing
	PriceDescription string `json:"PriceDescription,omitempty"`
	// URL used to buy the extension
	ShopURL string `json:"ShopURL,omitempty"`
	// Is an update available for this extension
	UpdateAvailable bool `json:"UpdateAvailable,omitempty"`
	// Extension version
	Version string              `json:"Version,omitempty"`
	License *LicenseInformation `json:"License,omitempty"`
}
