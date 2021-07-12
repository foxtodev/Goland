package myjson

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"text/tabwriter"
)

type conf struct {
	Name         string
	Age          byte
	IsAdmin      bool
	IsEditor     bool
	PersonalPage string
}

func ReadConfigJson(fileName string) conf {
	configData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	config := conf{}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalln(err)
	}

	return config
}

func (c conf) Display() {
	v := reflect.ValueOf(c)
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', 0)
	for i := 0; i < v.NumField(); i++ {
		fmt.Fprintf(w, "%s\t: %v\n", v.Type().Field(i).Name, v.Field(i).Interface())
	}
	w.Flush()
}
