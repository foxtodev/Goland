package main

import "server/server"

func main() {
	server.ServerStart()
	//myyaml.ReadConfigYamlToJson("conf.yaml").Display()
	//myyaml.ReadConfig("conf.yaml").Display()
	//myjson.ReadConfigJson("conf.json").Display()
	//fmt.Printf("%+v\n", myjson.ReadConfigJson("./conf.json"))
}
