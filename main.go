package main

import (
	"log"
	"regexp"
	"strings"
)

var (
	flagMatchReg = regexp.MustCompile(`^(-\w)?(, )?(--[a-zA-Z0-9-]+)?( )?(.+)?$`)

	options = make([]*option, 0)
)

// Commander is the type used
type Commander struct{}

// OptionString is a command line argument
func (c *Commander) OptionString(flags string, desc string, parser func(interface{}) string, defaultVal string) {
	if !flagMatchReg.MatchString(flags) {
		log.Panicln("pattern '", flags, "' is not in the correct format")
	}
	opt := newOption(parseFlags(flags))

	options = append(options, opt)
}

func main() {
	c := &Commander{}
	c.OptionString("-d", "desc", nil, "")
	c.OptionString("--directory", "desc", nil, "")
	c.OptionString("-d, --directory", "desc", nil, "")
	c.OptionString("-d <dir>", "desc", nil, "")
	c.OptionString("--directory <dir>", "desc", nil, "")
	c.OptionString("-d, --directory <dir>", "desc", nil, "")
	c.OptionString("-d, --directory <dir>..", "desc", nil, "")
	c.OptionString("-d, --directory [dir]", "desc", nil, "")
	c.OptionString("-d, --directory-name [dir]", "desc", nil, "")
}

func parseFlags(flags string) (short string, long string, extra string) {
	m := flagMatchReg.FindAllStringSubmatch(flags, -1)[0]
	return m[1], m[3], m[5]
}

type option struct {
	HasValue   string
	IsOptional bool
	IsRange    bool
	IsRequired bool
	LongFlag   string
	ShortFlag  string
}

func newOption(short string, long string, extra string) *option {
	opt := &option{ShortFlag: short, LongFlag: long}

	if strings.IndexRune(extra, '[') == 0 {
		opt.IsOptional = true
	} else if strings.IndexRune(extra, '<') == 0 {
		opt.IsRequired = true
	}

	if strings.Index(extra, "...") != -1 {
		opt.IsRange = true
	}

	return opt
}
