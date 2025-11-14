package session

import (
	"errors"
	"log"
	"os"
	"prepare-code/src/types"
	"strings"
)

type Session struct {
	Platform types.Platform
	Language string
	FileName string

	submissionVersion int8
	// apply weight group to aquire lock in terms of writes
}

func NewSession(platform types.Platform, language string) (Session, error) {
	fileName := platform.GetCodeName() + "." + language
	return Session{
		Platform: platform,
		Language: language,
		FileName: fileName,
	}, nil
}

func (s Session) GetFileName() string {
	return s.FileName
}

func (s Session) GetPlatform() types.Platform {
	return s.Platform
}

func (s Session) GetLanguage() string {
	return s.Language
}

func (s Session) CreateFile(fileName string) error {
	// now I wanted to copy the contents of the template file
	// to newly created file

	templateContentBin, err := os.ReadFile("./templates/template." + s.Language)
	if err != nil {
		return errors.New("unable to read template file: template." + s.Language + " err: " + err.Error())
	}

	if templateContentBin == nil || len(templateContentBin) == 0 {
		return errors.New("template file is empty: template." + s.Language)
	}

	// now get the lines as string
	templateLines := strings.Split(string(templateContentBin), "\n")
	log.Println("template lines: ", templateLines)

	return nil
}
