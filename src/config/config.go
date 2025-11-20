package config

import (
	"fmt"
	"os"
	"prepare-code/src/types"
	"strings"
)

const yamlFileName = "sample_config.yaml"
const OldInstallDirPath = "$HOME/.prepare-code"

var config *Config

type Config struct {
	InstallDir   string `yaml:"install_dir"`
	BypassPrompt bool   `yaml:"bypass_prompt"`
}

func setConfig() error {
	yamlFile, err := os.ReadFile(yamlFileName)
	if err != nil {
		return err
	}

	// now read and assign
	yaml := string(yamlFile)
	configValue := make(map[string]interface{})
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

	config = &Config{
		InstallDir:   configValue["install_dir"].(string),
		BypassPrompt: types.GetBoolFromString(configValue["bypass_prompt"].(string)),
	}

	return nil
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
		return OldInstallDirPath
	}

	return config.InstallDir
}

func CanBypassPrompt() bool {
	setConfig()

	return config.BypassPrompt
}
