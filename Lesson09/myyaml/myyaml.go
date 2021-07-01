package myyaml

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

type conf struct {
	Port         string
	DB_url       string
	Jaeger_url   string
	Sentry_url   string
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
}

func ReadConfig(fileName string) conf {
	buffer := []string{}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		configLine := strings.Split(string(scanner.Text()), ": ")
		buffer = append(buffer, configLine[1])
		/*if strings.Contains(configLine[0], "url") {
			err := config.IsUrl(configLine[1])
			if err != nil {
				log.Fatalln(err)
			}
		}*/
	}

	config := conf{
		Port:         buffer[0],
		DB_url:       buffer[1],
		Jaeger_url:   buffer[2],
		Sentry_url:   buffer[3],
		Kafka_broker: buffer[4],
		Some_app_id:  buffer[5],
		Some_app_key: buffer[6],
	}

	return config
}

func ReadConfigYamlToJson(fileName string) conf {
	configData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	data := "{\n"
	configLines := strings.Split(string(configData), "\n")
	for _, strLine := range configLines {
		configLine := strings.Split(string(strLine), ": ")
		data += "\t\"" + configLine[0] + "\": \"" + configLine[1] + "\",\n"
	}
	data = data[:len(data)-2] + "\n}"

	config := conf{}
	err = json.Unmarshal([]byte(data), &config)
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
