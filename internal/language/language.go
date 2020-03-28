package language

import (
	"github.com/andremedeiros/beaver/internal/config"
	"github.com/andremedeiros/beaver/internal/language/golang"
	"github.com/andremedeiros/beaver/internal/language/javascript"
	"github.com/andremedeiros/beaver/internal/language/ruby"
)

type Language interface {
	Name() string
	DefaultTasks(root string) map[string]config.Task
	Generators() map[string]func(args []string) error
}

var Languages = map[string]Language{
	"js":         javascript.New(),
	"javascript": javascript.New(),
	"go":         golang.New(),
	"golang":     golang.New(),
	"ruby":       ruby.New(),
}
