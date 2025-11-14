package session

import (
	"log"
	"os"
	"os/user"
)

func getUserName() string {
	curUser, err := user.Current()
	if err != nil || curUser == nil {
		log.Println("Error getting current user:", err)
		curUser = &user.User{
			Username: "<unknown>",
		}
	}

	hostName, err := os.Hostname()
	if err != nil || hostName == "" {
		log.Println("Error getting hostname:", err)
		hostName = "<unknown>"
	}

	return curUser.Username + "@" + hostName
}
