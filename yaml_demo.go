package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
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
    Cookie      string `yaml:"Cookie"`
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

func (p *Poc) PocGetter(path string) *Poc {
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatalf("yaml_file Get err: %#v\n", err)
    }
    err = yaml.Unmarshal(yamlFile, p)
    if err != nil {
        log.Fatalf("yaml.Unmarshal err: %#v\n", err)
    }
    return p
}
