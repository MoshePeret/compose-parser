package main

import (
	"fmt"
	"github.com/MoshPe/compose-parser/loader"
	"github.com/MoshPe/compose-parser/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := "docker-compose.yml"
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
		fmt.Println(service.DependsOn)
	}
}
