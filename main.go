package main

import (
	jsonyaml "github.com/sp-yduck/format-converter/json_yaml"
)

func main() {
	jsonyaml.JsonPrint("./json_yaml/alerts.json")
}
