package impex

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	path "path/filepath"
)

// Read YAML file and unmarshal its content into the `dst` data structure,
// using the `yaml` annotation of the structs.
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

// Marshal the content from the `dst` data structure, and savel into a YAML file
// using the `yaml` annotation of the structs.
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
