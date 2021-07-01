package server

import (
	"fmt"
	"net/http"
	"reflect"
	"server/config"
	"server/myjson"
	"server/myyaml"
)

func ServerStart() {
	cfg := config.ReadConfig()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		dataYAML := myyaml.ReadConfigYamlToJson("conf.yaml")
		dataStr := "<p  style='margin: 20px; font-size: 20px;'>YAML</p>" + DisplayData(dataYAML)
		dataJSON := myjson.ReadConfigJson("conf.json")
		dataStr += "<p  style='margin: 20px; font-size: 20px;'>JSON</p>" + DisplayData(dataJSON)
		//_, err := fmt.Fprint(writer, DisplayData(data))
		_, err := writer.Write([]byte(dataStr))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe("localhost:"+cfg.Port, nil)
	if err != nil {
		return
	}
}

func DisplayData(data interface{}) string {
	html := "<table border='1' style='margin: 20px; border-collapse: collapse;'>"
	html += fmt.Sprintln("<tr style='background-color:#bbb;'>" +
		"<th style='padding: 5px'>Field</th>" +
		"<th style='padding: 5px'>Value</th>" +
		"<th style='padding: 5px'>Validation URL</th>" +
		"</tr>")
	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		html += fmt.Sprintf("<tr><td style='padding: 5px'>%s:</td>"+
			"<td style='padding: 5px'>%v</td>"+
			"<td style='padding: 5px'>%v</td></tr>",
			v.Type().Field(i).Name, v.Field(i).Interface(), config.IsUrl(v.Field(i).String()))
	}
	html += "</table>"
	return html
}
