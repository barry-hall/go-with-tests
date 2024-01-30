package templating_test

import (
	"bytes"
	"go-with-tests/files"
	"go-with-tests/templating"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
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

	postRenderer, err := templating.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatalf("Render() err = %s", err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []files.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = files.Post{
			Title:       "Hello, World!",
			Body:        "This is a post",
			Description: "Lorem ipsum dolor sit amet",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := templating.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
