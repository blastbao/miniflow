package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Configs for miniflow
type Configs struct {
	Name     string  `json:"name"`
	Parallel int     `json:"parallel"`
	Tasks    []*Item `json:"tasks"`
}

// NewConfigs creates a new conf interface
func NewConfigs(fp string) *Configs {
	b := readJSONFile(fp)
	return load(b)
}

func readJSONFile(filePath string) []byte {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("read configs err: ", err)
	}
	return dat
}

func load(data []byte) *Configs {
	var c Configs
	json.Unmarshal(data, &c)
	return &c
}
