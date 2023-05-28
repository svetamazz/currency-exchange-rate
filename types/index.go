package helpers

type ExchangeRate struct {
	BTCtoUAH float64 `json:"btc_to_uah"`
}

type EmailList struct {
	Emails []string `json:"emails"`
}