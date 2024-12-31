package binance

import (
	binance_connector "github.com/binance/binance-connector-go"
	"pushbot/config"
)

func NewBNClient() *binance_connector.Client {
	apikey := config.HandleYaml().Exchange.Binance.Apikey
	secret := config.HandleYaml().Exchange.Binance.Secret
	client := binance_connector.NewClient(apikey, secret)
	return client
}
