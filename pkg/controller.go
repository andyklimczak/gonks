package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

type Controller struct{}

type Config struct {
	APIKey string `json:"api_key"`
	Stocks []string `json:"stocks"`
}

func (c Controller) Handle() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(usr.HomeDir, ".gonks.json")
	file, _ := ioutil.ReadFile(filePath)
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	stockModel := Stock{}
	stocks := stockModel.fetch(config.Stocks, config.APIKey)

	view := View{}
	view.Display(stocks)
}
