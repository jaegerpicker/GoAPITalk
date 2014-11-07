package main

import (
	"fmt"
	"io/ioutil"
    "gopkg.in/yaml.v2"
	"log"
)


// Struct that defines possible settings values
// Sucks that you have to add it here when you add a new one but go isn't quite dynamic enough to just add it at runtime
// That's both good and bad IMO
type Settings struct {
	DbHost          	string
	DbUser          	string
	DbPass          	string
	DbPort          	string
	DbName				string
	Environment     	string
	LogLevel        	string
	LogType         	string
	LogLocation     	string
    Port                string
}


// Get the settings from the settings.yaml file which is a symlink to the current env's settings.yaml file
func (s Settings) GetSettings() (settings Settings, err error) {
	// Get settings file
	fileName := buildSettingsFileName()
	// read the file into a []byte
	yamlData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(fmt.Sprintf("Could not read settings file %s err: %v", fileName, err))
		return s, err
	}
	// This makes a dynamic interface(object in go) of types string
	m := make(map[interface{}]string)
	// Use the yaml lib to marshal the data into the go lang struct using a reference to the memory of the struct
	err = yaml.Unmarshal(yamlData, &m)
	if err != nil {
		log.Println(fmt.Sprintf("Invalid yaml in %s, exiting becuase of err: %v", fileName, err))
	}
	// Might be a better path than this but it's not horrible
	for key, value := range m {
		switch key {
		case "DbHost":
			s.DbHost = value
			break
		case "DbPass":
			s.DbPass = value
			break
		case "DbUser":
			s.DbUser = value
			break
		case "DbName":
			s.DbName = value
			break
		case "DbPort":
			s.DbPort = value
			break
		case "Environment":
			s.Environment = value
			break
		case "LogLevel":
			s.LogLevel = value
			break
		case "LogType":
			s.LogType = value
			break
		case "LogLocation":
			s.LogLocation = value
			break
        case "Port":
            s.Port = value
            break
		}
	}
	return s, nil

}

func buildSettingsFileName() (filename string) {
	filename = fmt.Sprintf("settings.yaml")
	return filename
}
