package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234ABC",
		Username:  "Alex",
	},
	"john": {
		AuthToken: "456GHI",
		Username:  "John",
	},
	"Doe": {
		AuthToken: "789JKL",
		Username:  "Doe",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "Alex",
	},

	"john": {
		Coins:    300,
		Username: "John",
	},
	"Doe": {
		Coins:    1000,
		Username: "Doe",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var ClientData = LoginDetails{}
	ClientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &ClientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var ClientData = CoinDetails{}
	ClientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &ClientData
}

func (d mockDB) SetupDatabase() error {
	return nil
}
