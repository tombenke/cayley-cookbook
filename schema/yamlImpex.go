package main

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	path "path/filepath"
)

func ReadFromYaml(yamlPath string, dst interface{}) error {

	absPath, err := path.Abs(yamlPath)
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	buf, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	err = yaml.Unmarshal([]byte(buf), dst)
	if err != nil {
		log.Printf("YAML unmarshal error: %v", err)
		return err
	}
	return nil
}

func ReadFromYamlStream(yamlPath string, dst interface{}) error {

	absPath, err := path.Abs(yamlPath)
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	f, err := os.Open(absPath)
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	defer f.Close()

	yamlReader := bufio.NewReader(f)

	dec := yaml.NewDecoder(yamlReader)
	return dec.Decode(dst)
}

func SaveToYaml(yamlPath string, src interface{}) error {
	buf, err := yaml.Marshal(src)
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	absPath, err := path.Abs(yamlPath)
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return ioutil.WriteFile(absPath, buf, 0644)
}
