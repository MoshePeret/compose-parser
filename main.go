package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/MoshePeret/compose-parser/loader"
	"github.com/MoshePeret/compose-parser/types"
)

func main() {
	filename := "executor-compose.yml"
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("pass getwd")
	ymlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("pass readFile")
	cfg := types.ConfigDetails{
		Version:    "",
		WorkingDir: workingDir,
		ConfigFiles: []types.ConfigFile{
			{Filename: filename, Content: ymlFile},
		},
		Environment: nil,
	}

	project, err := loader.Load(cfg, func(options *loader.Options) {
		options.SkipConsistencyCheck = true
		options.SkipNormalization = true
		options.Name = filename
	})
	fmt.Println("pass load")
	if err != nil {
		panic(err)
	}
	fmt.Println("pass readFile")
	log.Println("Config file " + filename + " has been loaded!")
	str, _ := yaml.Marshal(project)
	fmt.Println(string(str))

	for _, service := range project.Services {
		fmt.Println(service.Name)
		fmt.Println(service.Labels)
		fmt.Println("init_container_policy: " + service.InitContainerPolicy)
		fmt.Printf("DependsOn: ")
		fmt.Println(service.DependsOn)
		for _, config := range service.InitContainer {
			empJSON, err := json.Marshal(config)
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Printf("\"name\":%s %s\n", config.Name, empJSON)
		}
	}
}

func isExist(depends []string, service string) bool {
	for _, v := range depends {
		if v == service {
			return true
		}
	}
	return false
}
