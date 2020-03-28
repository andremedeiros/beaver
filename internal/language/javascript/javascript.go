package javascript

import (
	"github.com/andremedeiros/beaver/internal/config"
)

type Javascript struct{}

func New() *Javascript {
	return &Javascript{}
}

func (*Javascript) Name() string {
	return "Javascript"
}

func (*Javascript) DefaultTasks(_ string) map[string]config.Task {
	return map[string]config.Task{
		"deps": config.Task{
			Description: "Fetch dependencies",
			Command:     "yarn install",
		},
	}
}

func (*Javascript) Generators() map[string]func(args []string) error {
	return map[string]func(args []string) error{}
}
