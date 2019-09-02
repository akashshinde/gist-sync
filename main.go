package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
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
	fmt.Printf("%+v",value)
	r, err := http.Get(value.Str)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(viper.GetString("SyncFilePath"), result, os.ModePerm)
}

func setupConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("/Users/akash/.gist-sync/config")
	err := viper.ReadInConfig();
	if err != nil {
		panic(err)
	}
}
