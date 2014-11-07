package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func setUpSettings() {
	err := os.Remove("settings.yaml")
	if err != nil {
		log.Fatal("Could not remove setting.yaml test file")
	}
	err = os.Symlink("settings-TEST.yaml", "settings.yaml")
	if err != nil {
		log.Fatal("Could not symlink to settings-TEST.yaml")
	}
}

func TestGetSettingsDev(t *testing.T) {
	setUpSettings()
	var settings Settings
	s, err := settings.GetSettings()
	if err != nil {
		t.Error(fmt.Sprintf("Error was encountered getting settings error: %v", err))
	}
	var mock_s Settings
	mock_s.DbUser = "root"
	mock_s.DbPass = "root"
	mock_s.DbHost = "localhost"
	mock_s.DbPort = "3306"
	mock_s.Environment = "TEST"
	mock_s.LogLevel = "INFO"
	mock_s.LogType = "FILE"
	mock_s.LogLocation = "log.txt"
	mock_s.DbName = "goapitalk_test"
	mock_s.Port = "3000"

	if !reflect.DeepEqual(s, mock_s) {
		//fmt.Errorf("Could not successfull get TEST settings in unit test")
		t.Error(fmt.Sprintf("Could not successfull get TEST settings in unit test settings: %v mock_settings: %v", s, mock_s))
	}
	tearDownSettings()
}

func tearDownSettings() {
	err := os.Remove("settings.yaml")
	if err != nil {
		log.Fatal("Could not remove setting.yaml test file")
	}
	err = os.Symlink("settings-LOCAL.yaml", "settings.yaml")
	if err != nil {
		log.Fatal("Could not symlink to settings-TEST.yaml")
	}
}
