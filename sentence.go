package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Formatter interface {
	Format(string, ...interface{}) string
}

type PlainFormatter struct{}

func (PlainFormatter) Format(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

type Sentence struct {
	Proser      Formatter
	Highlighter Formatter

	target  io.Writer
	builder strings.Builder
	quiet   bool
}

func NewSentence() *Sentence {
	return &Sentence{
		Proser:      NewANSI("2"),
		Highlighter: NewANSI("1+b"),
		target:      os.Stderr,
	}
}

func (s *Sentence) Plain(format string, args ...interface{}) *Sentence {
	_, _ = s.builder.WriteString(fmt.Sprintf(format, args...))
	return s
}

func (s *Sentence) Highlight(format string, args ...interface{}) *Sentence {
	return s.Plain(s.Highlighter.Format(format, args...))
}

func (s *Sentence) Quote(format string, args ...interface{}) *Sentence {
	return s.
		Prose(`"`).
		Highlight(strings.ReplaceAll(fmt.Sprintf(format, args...), `"`, `\"`)).
		Prose(`"`)
}

func (s *Sentence) Prose(format string, args ...interface{}) *Sentence {
	return s.Plain(s.Proser.Format(format, args...))
}

func (s *Sentence) Print() {
	s.builder.WriteString("\n")
	if _, err := s.target.Write([]byte(s.builder.String())); err != nil {
		panic(err)
	}
}
