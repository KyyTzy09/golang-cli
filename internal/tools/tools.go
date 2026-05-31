package tools

import "fmt"

type Tool struct {
	Name         string
	Requirements string
	Func         func(map[string]string) (any, error)
}

var Tools = map[string]Tool{
	"WriteFile": {
		Name:         "WriteFile",
		Requirements: "path, content",
		Func:         WriteFileContent,
	},
	"ScanFile": {
		Name:         "ScanFile",
		Requirements: "path",
		Func:         ScanFileContent,
	},
}

func RunTools(toolname string, args map[string]string) {
	exec, ok := Tools[toolname]
	if !ok {
		fmt.Println("Tool not found")
		return
	}

	exec.Func(args)
}