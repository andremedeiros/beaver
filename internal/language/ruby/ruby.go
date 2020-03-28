package ruby

import (
	"github.com/andremedeiros/beaver/internal/config"
)

type Ruby struct{}

func New() *Ruby {
	return &Ruby{}
}

func (*Ruby) Name() string {
	return "Ruby"
}

func (*Ruby) DefaultTasks(_ string) map[string]config.Task {
	return map[string]config.Task{
		"deps": config.Task{
			Description: "Fetch dependencies",
			Command:     "bundle",
		},
	}
}

func (*Ruby) Generators() map[string]func(args []string) error {
	return map[string]func(args []string) error{}
}
