package jsonyaml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RuleGroups struct {
	Groups []RuleGroup `json:"groups"`
}

type RuleGroup struct {
	Name  string      `json:"name"`
	Rules []AlertRule `json:"rules"`
}

type AlertRule struct {
	Alert       string         `json:"alert"`
	Expr        string         `json:"expr"`
	For         string         `json:"for"`
	Labels      LabelType      `json:"labels"`
	Annotations AnnotationType `json:"annotations"`
}

type LabelType struct {
	Severity string `json:"severity"`
}

type AnnotationType struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func json_print(file_name string) {
	text, err := ioutil.ReadFile(file_name)
	if err != nil {
		panic(err.Error())
	}
	var rule_groups RuleGroups
	err = json.Unmarshal(text, &rule_groups)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rule_groups)
}

func main() {
	json_print("./json_yaml/alerts.json")
}
