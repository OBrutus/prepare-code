package session

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"prepare-code/src/config"
	"prepare-code/src/platform/generic_platform"
	"prepare-code/src/system"
	"prepare-code/src/types"
	"prepare-code/src/uuid"
	"strings"
	"time"
)

var ExecutionMap = make(map[string]interface{})

type Session struct {
	Id       string
	Platform types.Instance
	Language string
	FileName string

	userName          string
	submissionVersion uint8
	startDate         time.Time
	targetDirectory   string

	// : apply weight group to aquire lock in terms of writes
}

func NewSession(platform types.Instance, language string) (Session, error) {
	fileName := platform.GetCodeName() + "." + language

	dir, err := os.Getwd()
	if err != nil || os.Getenv("redacted") != "" {
		fmt.Println("Error getting current working directory:", err)
		dir = "<unknown>"
	}

	return Session{
		Id:                uuid.New(),
		Platform:          platform,
		Language:          language,
		FileName:          fileName,
		userName:          system.GetUserName(),
		startDate:         time.Now(),
		targetDirectory:   dir,
		submissionVersion: 0,
	}, nil
}

func NewGenericSession(language string) (Session, error) {
	return NewSession(generic_platform.GetPlatform(), language)
}

func (s Session) GetFileName() string {
	return s.FileName
}

func (s Session) GetPlatform() types.Instance {
	return s.Platform
}

func (s Session) GetLanguage() string {
	return s.Language
}

func (s Session) CreateFile() error {
	// now I wanted to copy the contents of the template file
	// to newly created file
	templateLines, err := s.getTemplateLines()
	if err != nil {
		return errors.New("unable to get template lines: " + err.Error())
	}

	writeFile, err := os.Create(s.FileName)
	if err != nil {
		return errors.New("unable to create file: " + s.FileName + " err: " + err.Error())
	}
	defer writeFile.Close()

	dataItr := 0
	dataList := []string{
		s.userName,
		s.startDate.Format("2006-01-02"),
		"0.0",
		string(s.Platform.GetPlatformName()),
		s.targetDirectory,
	}

	// now based upon the details let's fill one by one
	writer := bufio.NewWriter(writeFile)
	for i, line := range templateLines {
		if strings.Contains(line, "%s") {
			templateLines[i] = fmt.Sprintf(line, dataList[dataItr])
			line = templateLines[i]
			dataItr++
		}

		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return errors.New("unable to write to file: " +
				s.FileName +
				" err: " +
				err.Error(),
			)
		}
	}
	writer.Flush()

	return nil
}

func (s Session) getTemplateLines() ([]string, error) {
	installDir := config.GetInstallDir()
	internalPath := fmt.Sprintf("/template/%s/template.%s", s.Language, s.Language)
	absolutePathFileName := filepath.Join(installDir, internalPath)
	templateContentBin, err := os.ReadFile(absolutePathFileName)
	if err != nil {
		return nil, errors.New(
			"unable to read template file: template." +
				s.Language + " err: " + err.Error(),
		)
	}

	if templateContentBin == nil || len(templateContentBin) == 0 {
		return nil, errors.New("template file is empty: template." + s.Language)
	}

	// now get the lines as string
	templateLines := strings.Split(string(templateContentBin), "\n")

	return templateLines, nil
}
