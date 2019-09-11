package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	setupConfig()
	resp, err := http.Get("https://api.github.com/gists/" + viper.GetString("GistId"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	value := gjson.Get(string(data), "files.kubeconfig.raw_url")
	fmt.Println("Downloading kubeconfig content from url.")
	r, err := http.Get(value.Str)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(viper.GetString("SyncFilePath"), result, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}

func setupConfig() {
	homeDir, _ := os.UserHomeDir()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filepath.Join(homeDir,".gist-sync/config"))
	err := viper.ReadInConfig();
	if err != nil {
		panic(err)
	}
}
