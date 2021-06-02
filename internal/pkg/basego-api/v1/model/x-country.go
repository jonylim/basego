package model

// XCountry contains a country's information.
type XCountry struct {
	RedisNil        bool   `json:"-"            redis:"redisNil"`
	ID              int32  `json:"id"           redis:"id"`
	CommonName      string `json:"commonName"   redis:"commonName"`
	OfficialName    string `json:"officialName" redis:"officialName"`
	CountryCodeISO2 string `json:"iso2Code"     redis:"iso2Code"`
	CountryCodeISO3 string `json:"iso3Code"     redis:"iso3Code"`
	CallingCode     string `json:"callingCode"  redis:"callingCode"`
	CurrencyCode    string `json:"currencyCode" redis:"currencyCode"`
	IsEnabled       bool   `json:"isEnabled"    redis:"isEnabled"`
	IsHidden        bool   `json:"isHidden"     redis:"isHidden"`
}
