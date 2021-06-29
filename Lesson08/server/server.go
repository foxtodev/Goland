package server

import (
	"fmt"
	"net/http"
	"reflect"
	"server/config"
)

func ServerStart() {

	cfg := config.ReadConfig()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := "<table border='1' style='margin: 20px; border-collapse: collapse;'>"
		data += fmt.Sprintln("<tr style='background-color:#bbb;'><th style='padding: 5px'>Field</th><th style='padding: 5px'>Value</th><th style='padding: 5px'>Validation URL</th></tr>")
		v := reflect.ValueOf(cfg)
		for i := 0; i < v.NumField(); i++ {
			data += fmt.Sprintf("<tr><td style='padding: 5px'>%s:</td><td style='padding: 5px'>%v</td><td style='padding: 5px'>%v</td></tr>", v.Type().Field(i).Name, v.Field(i).Interface(), config.IsUrl(v.Field(i).String()))
		}
		data += "</table>"
		_, err := fmt.Fprint(writer, data)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe("localhost:"+cfg.Port, nil)
	if err != nil {
		return
	}

}
