package cli

import (
	"io"
	"os"
	"strings"
)

type Formatter interface {
	Format(string) string
}

type PlainFormatter struct{}

func (PlainFormatter) Format(s string) string {
	return s
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

func (s *Sentence) Plain(w string) *Sentence {
	_, _ = s.builder.WriteString(w)
	return s
}

func (s *Sentence) Highlight(w string) *Sentence {
	return s.Plain(s.Highlighter.Format(w))
}

func (s *Sentence) Prose(w string) *Sentence {
	return s.Plain(s.Proser.Format(w))
}

func (s *Sentence) Print() {
	s.builder.WriteString("\n")
	if _, err := s.target.Write([]byte(s.builder.String())); err != nil {
		panic(err)
	}
}
