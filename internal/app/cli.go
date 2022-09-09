package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
)

func (a *App) executor(in string) {
	cmd := strings.Split(in, " ")
	switch cmd[0] {
	case "start":
		if len(cmd) > 3 {
			fmt.Println("Invalid number of arguments")
			return
		}

		desc := ""
		if len(cmd) == 3 {
			desc = cmd[2]
		}

		err := a.d.Start(cmd[1], desc, time.Now().Unix())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task started successfully")
		}
	case "stop":
		if len(cmd) != 2 {
			fmt.Println("Invalid number of arguments")
			return
		}

		err := a.d.End(cmd[1], time.Now().Unix())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task stopped successfully")
		}
	case "list":
		if len(cmd) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}

		records, err := a.d.List("")
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, r := range records {
			endTime := time.Unix(r.End, 0).Format(time.RFC1123)
			if r.End == 0 {
				endTime = "in progress"
			}

			fmt.Printf("%s %s %s\n", r.Name, time.Unix(r.Start, 0).Format(time.RFC1123), endTime)
		}
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
	}
}

func (a *App) completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "start", Description: "Start task with specified name and optional description"},
		{Text: "stop", Description: "Stop task by name"},
		{Text: "add", Description: "Add a new custom record with the given name, start and end (UNIX time)"},
		{Text: "edit", Description: "Edit task by name, start and end (UNIX time)"},
		{Text: "delete", Description: "Delete task by name"},
		{Text: "list", Description: "List all tasks"},
		{Text: "export", Description: "Export tasks to CSV file"},
		{Text: "exit", Description: "Exit the program"},
	}

	cl := in.CurrentLine()
	var f []string
	f = strings.Fields(cl)
	if strings.HasSuffix(cl, " ") {
		f = append(f, " ")
	}
	if len(f) == 1 {
		return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
	}

	return prompt.FilterHasPrefix([]prompt.Suggest{}, in.GetWordBeforeCursor(), true)
}
