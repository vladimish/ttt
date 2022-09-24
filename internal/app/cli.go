package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
)

func (a *App) executor(in string) {
	cmd := strings.Split(in, " ")
	switch cmd[0] {
	case "start":
		if len(cmd) > 3 {
			fmt.Println("invalid number of arguments")
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
			fmt.Println("task started successfully")
		}
	case "stop":
		if len(cmd) != 2 {
			fmt.Println("invalid number of arguments")
			return
		}

		err := a.d.End(cmd[1], time.Now().Unix())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("task stopped successfully")
		}
	case "add":
		if len(cmd) == 3 {
			cmd = append(cmd, "")
		} else if len(cmd) != 4 {
			fmt.Println("invalid number of arguments")
			return
		}

		start, err := strconv.Atoi(cmd[1])
		if err != nil {
			fmt.Println("start must be an integer")
			return
		}

		end, err := strconv.Atoi(cmd[2])
		if err != nil {
			fmt.Println("end must be an integer")
			return
		}

		err = a.d.Start(cmd[0], cmd[3], int64(start))
		if err != nil {
			fmt.Printf("can't start task: %s\n", err.Error())
			return
		}

		err = a.d.End(cmd[0], int64(end))
		if err != nil {
			fmt.Printf("can't start task: %s\n", err.Error())
			return
		}
	case "delete":
		if len(cmd) != 1 {
			fmt.Println("invalid number of arguments")
			return
		}

		err := a.d.Delete(cmd[0])
		if err != nil {
			fmt.Printf("can't delete record: %s\n", err.Error())
			return
		}
	case "list":
		if len(cmd) != 1 {
			fmt.Println("invalid number of arguments")
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
		{Text: "add", Description: "Add a new custom record with the given name, start, end (UNIX time) and optional description."},
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
