package main

import (
	"fmt"
	"github.com/MoshePeret/compose-parser/loader"
	"github.com/MoshePeret/compose-parser/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type depends struct {
}

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
	//s := "s"
	//for _, serviceA := range project.Services {
	//	preDepends := serviceA.GetPreRunDependencies()
	//	startOrderDepends := serviceA.GetStartOrderDependencies()
	//	for _, service := range project.Services {
	//		service.Name = "a"
	//		if isExit(preDepends, service.Name) {
	//			if serviceA.DependsPre == nil {
	//				*serviceA.DependsPre = &s
	//			}
	//			//serviceA.DependsPre[serviceA.Name] = service
	//		} else if isExit(startOrderDepends, service.Name) {
	//			if serviceA.DependsStart == nil {
	//				serviceA.DependsStart = &s
	//			}
	//			//serviceA.DependsStart[serviceA.Name] = service
	//		}
	//	}
	//	fmt.Println(*serviceA.DependsPre)
	//}

	for _, service := range project.Services {
		fmt.Println(service.Name)
		fmt.Println("pre_run_policy: " + service.PreRunPolicy)
		fmt.Println(service.DependsOnPreRun)
		fmt.Println(service.DependsOnStartOrder)
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
