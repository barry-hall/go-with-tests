package templating_test

import (
	"bytes"
	"go-with-tests/files"
	"go-with-tests/templating"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = files.Post{
			Title:       "Hello, World!",
			Body:        "This is a post",
			Description: "Lorem ipsum dolor sit amet",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := templating.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello, World!</h1>`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
