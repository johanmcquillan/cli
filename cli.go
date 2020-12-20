package cli

import (
	"io/ioutil"
)

var logger = DefaultCLI()

type Level uint8

const (
	LevelQuiet Level = iota
	LevelError
	LevelWarning
	LevelInfo
)

func (l Level) String() string {
	switch l {
	case LevelQuiet:
		return ""
	case LevelError:
		return "ERROR"
	case LevelWarning:
		return "WARNING"
	case LevelInfo:
		return "INFO"
	default:
		return "UNKNOWN"
	}
}

type CLI struct {
	Level Level

	levelToFormatter map[Level]Formatter
}

func DefaultCLI() *CLI {
	return &CLI{
		Level: LevelError,
		levelToFormatter: map[Level]Formatter{
			LevelError:   NewANSI("1+b"),
			LevelWarning: NewANSI("11+b"),
			LevelInfo:    NewANSI("12+b"),
		},
	}
}

func (c *CLI) WithLevel(level Level) *Sentence {
	if level == LevelQuiet || c.Level < level {
		s := NewSentence()
		s.quiet = true
		s.target = ioutil.Discard
		return s
	}
	formatter, ok := c.levelToFormatter[level]
	if !ok {
		formatter = PlainFormatter{}
	}
	return NewSentence().
		Plain(formatter.Format(level.String())).
		Plain(": ")
}

func (c *CLI) SetLevel(level Level) {
	c.Level = level
}

func (c *CLI) Info() *Sentence {
	return c.WithLevel(LevelInfo)
}

func (c *CLI) Warning() *Sentence {
	return c.WithLevel(LevelWarning)
}

func (c *CLI) Error() *Sentence {
	return c.WithLevel(LevelError)
}

func SetLevel(level Level) {
	logger.SetLevel(level)
}

func Info() *Sentence {
	return logger.Info()
}

func Warning() *Sentence {
	return logger.Warning()
}

func Error() *Sentence {
	return logger.Error()
}
