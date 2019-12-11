package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

type Poc struct {
	Name   string  `yaml:"name"`
	Set    Set     `yaml:"set"`
	Rules  []Rules `yaml:"rules"`
	Detail Detail  `yaml:"detail"`
}

type Set struct {
	Filename    string `yaml:"filename"`
	FileContent string `yaml:"fileContent"`
}

type Headers struct {
	Cookie string `yaml:"Cookie"`
	ContentType string `yaml:"Content-Type"`
}

type Rules struct {
	Method          string  `yaml:"method"`
	Path            string  `yaml:"path"`
	Headers         Headers `yaml:"headers"`
	Body            string  `yaml:"body"`
	Search          string  `yaml:"search"`
	FollowRedirects bool    `yaml:"follow_redirects"`
	Expression      string  `yaml:"expression"`
}

type Detail struct {
	Author string   `yaml:"author"`
	Links  []string `yaml:"links"`
}

func (p *Poc) GetPoc(path string) *Poc {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("yaml_file Get err: %#v\n", err)
	}
	err = yaml.Unmarshal(yamlFile, p)
	if err != nil {
		log.Fatalf("Unmarshal err: %#v\n", err)
	}
	return p
}

// 利用反射，struct转map
func StructToMap(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]string)
	// 若传入指针类型，编译异常：panic: reflect: NumField of non-struct type
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).String()
	}
	return data
}
