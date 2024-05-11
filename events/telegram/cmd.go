package telegram

import (
	"EfimBot/e"
	"EfimBot/storage"
	"context"
	"errors"
	"log"
	"strings"
)

const (
	startCmd  = "/start"
	helpCmd   = "/помощь"
	initCmd   = "/init"
	createCmd = "/создать"
)

var ErrBlankBody = errors.New("command body is blank")
var ErrInsufficientArguments = errors.New("can't create table, insufficient arguments")
var ErrInputIsNotCommand = errors.New("input string doesn't start with </>")
var ErrInvalidInput = errors.New("number of arguments doesn't match the pattern")

func (p *Processor) doCmd(text string, chatID int, username string) error {
	if text[0] != '/' {
		return e.Wrap("incorrect input", ErrInputIsNotCommand)
	}

	text = strings.ToLower(text)

	cmd, cmdBody := p.parseInput(text)

	if (cmd != "/start" && cmd != "/help") && cmdBody == "" {
		return e.Wrap("command doesn't contain body", ErrBlankBody)
	}

	log.Printf("Got new command %s from %s", text, username)

	switch cmd {
	case initCmd:
		return p.initCmd(cmdBody)
	case createCmd:
		return p.createCmd(cmdBody)
	}

	return nil
}

func (p *Processor) parseInput(text string) (leftPart string, rightPart string) {
	for i := 0; i < len(text)-1; i++ {
		if text[i] == ' ' {
			return text[:i], text[i+1:]
		}
	}
	return text, ""
}

func (p *Processor) initCmd(cmdBody string) error {
	table, _ := createTableFromCmdBody(cmdBody)
	return p.storage.Init(context.TODO(), table)
}

func (p *Processor) createCmd(cmdBody string) error {
	obj, args := p.parseInput(cmdBody)

	switch obj {
	case "проект":
		return p.createProject(args)
	case "задачу":
		return p.createTask(args)
	case "отдел":
		return p.createDepartment(args)
	case "суботдел":
		return p.createSubDepartment(args)
	}
}

func (p *Processor) createProject(args string) error {
	if strings.Count(args, ` `) != 3 {
		return e.Wrap("can't create project", ErrInvalidInput)
	}

	table := storage.Table{
		Name:      "projects",
		Columns:   "codename,project_name,project_manager",
		Dimension: 3,
	}

	parameters := make([]string, 3)

	count := 0
	firstIndex := 0
	for i := 0; i < len(args); i++ {
		if string(args[i]) == ` ` {
			count++
		}
		if count == 1 && string(args[i]) == ` ` {
			parameters[0] = args[:i]
			firstIndex = i
		}
		if count == 2 {
			parameters[1] = args[firstIndex+1 : i]
			parameters[2] = args[i+1:]
			break
		}
	}
	return p.storage.Insert(context.TODO(), table, parameters)
}

func (p *Processor) createTask(args string) error {

}

func (p *Processor) createDepartment(args string) error {

}

func (p *Processor) createSubDepartment(args string) error {

}

func createTableFromCmdBody(cmdBody string) (storage.Table, error) {
	args := strings.Split(cmdBody, ` `)

	if len(args) < 2 {
		return storage.Table{}, e.Wrap("not enough args to create table", ErrInsufficientArguments)
	}

	name := args[0]

	dirtyColumns := strings.TrimSpace(strings.Join(args[1:], ""))

	columns := correctColumnsFormat(dirtyColumns)

	return storage.Table{Name: name, Columns: columns}, nil
}

func correctColumnsFormat(cmdBody string) string {
	corrected := make([]string, 0, len(cmdBody))
	sep := ""
	startingCommasSkipped := false
	for i := 0; i < len(cmdBody); i++ {
		char := string(cmdBody[i])
		if char == "," && startingCommasSkipped {
			sep = ","
		}
		if char != "," {
			corrected = append(corrected, sep+char)
			sep = ""
			startingCommasSkipped = true
		}
	}
	return strings.TrimSpace(strings.Join(corrected, ""))
}

func parseArgs() []string {

}

func parseArgsCareName(namePos int) []string {

}
