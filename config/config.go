package config

import (
	"database/sql"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Instance struct {
	Name           string `yaml:"name"`
	Informixserver string `yaml:"informixserver"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	db             *sql.DB
}

type InstanceList struct {
	Servers []Instance `yaml:"servers"`
}

type Probes struct {
	Parametro   string `yaml:"parametro"`
	Types       string `yaml:"type"`
	Description string `yaml:"description"`
	Query       string `yaml:"query"`
	Label       string `yaml:"label"`
}

type ConfigYaml struct {
	Metrics []Probes `yaml:"metrics"`
}

func LoadConfig(filename *string) (*ConfigYaml, error) {

	bytes, err := ioutil.ReadFile(*filename)
	if err != nil {
		return &ConfigYaml{}, err
	}

	var c ConfigYaml
	err = yaml.Unmarshal([]byte(bytes), &c)
	if err != nil {

		return &ConfigYaml{}, err
	}

	return &c, nil
}

func LoadConfig2(filename *string) (*InstanceList, error) {

	bytes, err := ioutil.ReadFile(*filename)
	if err != nil {
		return &InstanceList{}, err
	}

	var c InstanceList
	err = yaml.Unmarshal([]byte(bytes), &c)
	if err != nil {

		return &InstanceList{}, err
	}
	
	return &c, nil
}
