package textconverter

import (
	"bufio"
	"io"
	"strings"
)

const (
	lt  = "<"
	gt  = ">"
	amp = "&"
	nl  = "\n"
)

type runeConverter func(r rune) (string, bool)

// Converter handles the conversion of unicode text to other formats
type Converter struct {
	convertedLine []string
	result        []string
	runeConverter runeConverter
}

//ConvertToHTML converts unitoce text to HTML friendly text
func (c Converter) ConvertToHTML(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		c.basicHTMLEncode(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return c.result, err
	}
	return c.result, nil
}

func (c *Converter) basicHTMLEncode(s string) {
	for _, ch := range s {
		str, _ := c.runeConverter(ch)
		c.convertedLine = append(c.convertedLine, str)
	}
	c.addNewLine()
}

func (c *Converter) addNewLine() {

	c.result = append(c.result, strings.Join(c.convertedLine[:], ""))
	c.result = append(c.result, "<br />")

	c.convertedLine = c.convertedLine[:0]
}

type runeConverterChain struct {
	converters []runeConverter
}

func (c runeConverterChain) Convert(r rune) (string, bool) {
	for _, currentConverter := range c.converters {
		str, isApplied := currentConverter(r)
		if isApplied {
			return str, true
		}
	}
	return string(r), false
}

func NewDefaultConverter() runeConverter {
	composite := runeConverterChain{
		converters: []runeConverter{
			newOneToOneConvert(lt, "&lt;"),
			newOneToOneConvert(gt, "&gt;"),
			newOneToOneConvert(amp, "&amp;"),
			newOneToOneConvert(nl, "&lt;"),
			func(r rune) (string, bool) {
				if string(r) == nl {
					return string(r) + "</br>", true
				} else {
					return string(r), false
				}
			},
		},
	}

	return composite.Convert
}

func newOneToOneConvert(in string, out string) runeConverter {
	return func(r rune) (string, bool) {
		if string(r) == in {
			return out, true
		} else {
			return string(r), false
		}
	}
}
