package textconverter

import (
	"strings"
	"testing"
)

func TestConverter(t *testing.T) {

	t.Run("Do something", func(t *testing.T) {
		converter := NewDefaultConverter()
		c := Converter{runeConverter: converter}
		data := "coucou\nça va ?\nmoi oui\n <3"
		reader := strings.NewReader(data)
		content, err := c.ConvertToHTML(reader)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(content)
	})
}
