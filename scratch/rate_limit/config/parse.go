package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Interface interface {
	Parse(configFilePath string) (*Config, error)
}

var _ Interface = &Parser{}

type Parser struct {
	doParse func(rawContent []byte) (*Config, error)
}

func NewParse() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(configFilePath string) (*Config, error) {
	rawContent, err := readFile(configFilePath)
	if nil != err {
		return nil, fmt.Errorf("Parse: %s", err)
	}

	rules, err := p.doParse(rawContent)
	if nil != err {
		return nil, fmt.Errorf("Parse: %s", err)
	}

	return rules, nil
}

type ParseJson struct {
	Parser
}

func NewParseJson() *ParseJson {
	pj := &ParseJson{}
	pj.doParse = pj.doParseImpl
	return pj
}

func (pj *ParseJson) doParseImpl(rawContent []byte) (*Config, error) {
	config := Config{}
	if err := json.Unmarshal(rawContent, &config); nil != err {
		return nil, fmt.Errorf("Parse: %s", err)
	}

	return &config, nil
}

type ParseYaml struct {
	Parser
}

func NewParseYaml() *ParseYaml {
	py := &ParseYaml{}
	py.doParse = py.doParseImpl
	return py
}

func (py *ParseYaml) doParseImpl(rawContent []byte) (*Config, error) {
	config := Config{}
	if err := yaml.Unmarshal(rawContent, &config); nil != err {
		return nil, fmt.Errorf("Parse: %s", err)
	}

	return &config, nil
}

func readFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, fmt.Errorf("readFile: %s", err)
	}

	return content, nil
}
