package dev

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mini/generator"
	"mini/helpers"
	"net/http"
	"os"
)

//Serve : works as a temporary server for dev purposes
func Serve() {
	http.Handle("/", http.FileServer(http.Dir("./webpage")))

	var conf generator.Configuration

	if _, err := os.Stat("./conf"); !os.IsNotExist(err) {
		config, err := os.Open("./conf/config.json")
		helpers.Check(err)

		byteConfig, err := ioutil.ReadAll(config)
		helpers.Check(err)

		config.Close()
		json.Unmarshal(byteConfig, &conf)

		fmt.Printf("Listening at port %s\n", conf.Port)

		http.ListenAndServe(conf.Port, nil)
	}
}
