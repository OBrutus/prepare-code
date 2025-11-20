package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"prepare-code/src/constants"
	"prepare-code/src/types"
	"strings"
)

var yamlFileName = "config.yaml"

var config *Config

type Config struct {
	InstallDir   string `yaml:"install_dir"`
	BypassPrompt bool   `yaml:"bypass_prompt"`
}

func getOldInstallDirPath() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, "."+constants.AppName)
}

func setConfig() error {
	// now read and assign
	configValue := GetConfigMap()

	config = &Config{
		InstallDir:   getOldInstallDirPath(),
		BypassPrompt: types.GetBoolFromString(configValue["bypass_prompt"].(string)),
	}

	return nil
}

func GetConfigMap() map[string]interface{} {
	configValue := make(map[string]interface{})

	yamlFilePath := filepath.Join(getOldInstallDirPath(), yamlFileName)
	yamlFile, err := os.ReadFile(yamlFilePath)
	if err != nil {
		fmt.Println("Error reading config file. Err:", err)
		return configValue
	}

	yaml := string(yamlFile)
	for _, line := range strings.Split(yaml, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}

		keyVal := strings.Split(line, ":")
		key := strings.TrimSpace(keyVal[0])
		val := strings.TrimSpace(keyVal[1])
		configValue[key] = val
	}

	return configValue
}

func GetInstallDir() string {
	if config != nil {
		return config.InstallDir
	}

	err := setConfig()
	if err != nil {
		// for backward compatibility sending
		// previous build's install dir path
		fmt.Printf("Unable to read the config yaml file. Error: %v", err)
		return getOldInstallDirPath()
	}

	return config.InstallDir
}

func CanBypassPrompt() bool {
	if config == nil {
		setConfig()
	}

	return config.BypassPrompt
}
