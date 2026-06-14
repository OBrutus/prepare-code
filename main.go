package main

import (
	"fmt"
	"os"
	"prepare-code/src/config"
	"prepare-code/src/editor"
	"prepare-code/src/platform"
	"prepare-code/src/session"
	"prepare-code/src/system"
)

func main() {
	// start of the CLI application
	// accept URL
	var url string

	if len(os.Args) < 2 {
		fmt.Print("Enter the URL of problem: ")
		fmt.Scanln(&url)
	} else if os.Args[1] == "config" {
		configMap := config.GetConfigMap()
		fmt.Printf("Your config map is : ")
		fmt.Println(configMap)
		return
	} else if os.Args[1] == "raw" {
		genericHandover()
		return
	} else if os.Args[1] == "update" {
		system.Update()
	} else {
		url = os.Args[1]
	}

	// which platform it is
	platform, err := platform.ResolvePlatform(url)
	if err != nil {
		fmt.Println("[-] Error while getting platform: ", err)
		return
	}

	fmt.Println("What a good time to do ", platform.GetPlatformName())
	fmt.Println("Problem code name is ", platform.GetCodeName())

	language := getLanguage()
	session, err := session.NewSession(platform, language)
	if err != nil {
		fmt.Println("[-] Error while creating session: ", err)
		return
	}

	err = session.CreateFile()
	if err != nil {
		fmt.Println("[-] Error while creating file: ", err)
		return
	}

	fmt.Println("[+] File created successfully!")

	err = editor.TryOpenInEditor(session.FileName)
	if err != nil {
		fmt.Println("[-] Err: ", err)
	}
}

func getLanguage() string {
	var language string

	fmt.Print("Language to use (eg: java, go, cpp, py etc) [Default: Java]: ")
	fmt.Scanln(&language)
	if len(language) == 0 {
		language = "java"
	}

	return language
}

func genericHandover() {
	fmt.Printf("Giving for raw template")
	if len(os.Args) < 3 {
		fmt.Printf("You need to choose a language explicitly for raw template")
		return
	}

	language := os.Args[2]
	session, err := session.NewGenericSession(language)
	if err != nil {
		fmt.Println("Error while creating raw session: ", err)
		return
	}

	// check file name
	if len(os.Args) >= 4 {
		session.FileName = os.Args[3]
	}

	err = session.CreateFile()
	if err != nil {
		fmt.Println("Error while creating raw file: ", session.FileName, "\nError: ", err)
		return
	}

	fmt.Println("Raw file created successfully with name: ", session.FileName)
}
