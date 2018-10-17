package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

type jsonConfig struct {
	Lang string `json:"lang"`
}

var (
	configDir  string
	configPath string
)

func setConfig(c *cli.Context) {
	config := loadConfig()

	if c.String("lang") != "" {
		config.Lang = c.String("lang")
	}
	bdata, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Cannot encode json:", err.Error())
		os.Exit(ExitCodeError)
	}

	if _, err := os.Stat(configDir); err != nil {
		os.Mkdir(configDir, 0777)
	}

	err = ioutil.WriteFile(configPath, []byte(bdata), os.ModePerm)
	if err != nil {
		fmt.Println("Cannot save json:", err.Error())
		os.Exit(ExitCodeError)
	}
}

func loadConfig() jsonConfig {
	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Println("Cannot find homedir:" + err.Error())
		os.Exit(ExitCodeError)
	}
	configDir = homeDir + "/.trivia"
	configPath = configDir + "/settings.json"

	var config jsonConfig
	if fileExists(configPath) {
		bytes, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Println("Cannot read config:", err.Error())
			os.Exit(ExitCodeError)
		}
		if err := json.Unmarshal(bytes, &config); err != nil {
			fmt.Println("Cannot load config:", err.Error())
			os.Exit(ExitCodeError)
		}
	}
	return config
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
