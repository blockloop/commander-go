package main

import (
	"log"
	"regexp"
)

var (
	shortFlagReg = regexp.MustCompile(`^-[A-Za-z0-9]`)
	longFlagReg  = regexp.MustCompile(`^--[A-Za-z0-9-]+`)
	// flagMatchReg = regexp.MustCompile(`^(-\w)?(, )?(--[a-zA-Z]+)?( )?([<\[]\w+[\]>])?$`)
	flagMatchReg = regexp.MustCompile(`^(-\w)?(, )?(--[a-zA-Z-]+)?( )?([<\[]\w+[\]>])?(.*)$`)
)

// Commander is the type used
type Commander struct{}

// OptionString is a command line argument
func (c *Commander) OptionString(flags string, desc string, parser func(interface{}) string, defaultVal string) {
	if !flagMatchReg.MatchString(flags) {
		log.Panicln("pattern '", flags, "' is not in the correct format")
	}
	short, long, extra := parseFlags(flags)
	// fmt.Printf("ORIGINAL: '%s'\n", flags)
	// fmt.Println("SHORT:", short)
	// fmt.Println("LONG:", long)
	// fmt.Println("EXTRA:", extra)
	// fmt.Println("---")
	// fmt.Println("---")
}

func main() {
	c := &Commander{}
	c.OptionString("-d", "desc", nil, "")
	c.OptionString("--directory", "desc", nil, "")
	c.OptionString("-d, --directory", "desc", nil, "")
	c.OptionString("-d <dir>", "desc", nil, "")
	c.OptionString("--directory <dir>", "desc", nil, "")
	c.OptionString("-d, --directory <dir>", "desc", nil, "")
	c.OptionString("-d, --directory [dir]", "desc", nil, "")
	c.OptionString("-d, --directory-name [dir]", "desc", nil, "")
}

func parseFlags(flags string) (short string, long string, extra []string) {
	m := flagMatchReg.FindAllStringSubmatch(flags, -1)[0]
	return m[1], m[3], m[5:]
}
