package main

import (
	"fmt"
	"os"
	"prepare-code/src/platform"
	"prepare-code/src/session"
)

func main() {
	// start of the CLI application
	// accept URL
	var url string

	if len(os.Args) < 2 {
		fmt.Print("Enter the URL of problem: ")
		fmt.Scanln(&url)
	} else {
		url = os.Args[1]
	}

	// which platform it is
	platform, err := platform.GetPlatform(url)
	if err != nil {
		fmt.Println("Error while getting platform: ", err)
		return
	}

	fmt.Println("What a good time to do ", platform.GetPlatformName())
	fmt.Println("Problem code name is ", platform.GetCodeName())

	language := getLanguage()
	session, err := session.NewSession(platform, language)
	if err != nil {
		fmt.Println("Error while creating session: ", err)
		return
	}

	err = session.CreateFile()
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}

	fmt.Println("File created successfully!")
}

func getLanguage() string {
	var language string

	fmt.Print("Language to use (eg: java, cpp, py etc) [Default: Java]: ")
	fmt.Scanln(&language)
	if len(language) == 0 {
		language = "java"
	}

	return language
}
