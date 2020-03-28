package golang

import (
	"github.com/andremedeiros/beaver/internal/config"
)

type Golang struct{}

func New() *Golang {
	return &Golang{}
}

func (*Golang) Name() string {
	return "Golang"
}

func (*Golang) DefaultTasks(_ string) map[string]config.Task {
	return map[string]config.Task{
		"test": config.Task{
			Description: "Run tests",
			Command:     "go test -v ./...",
		},
		"build": config.Task{
			Description: "Build project",
			Command:     "go build ./",
		},
		"deps": config.Task{
			Description: "Fetch dependencies",
			Command:     "go mod download",
		},
	}
}

func (*Golang) Generators() map[string]func(args []string) error {
	return map[string]func(args []string) error{}
}
