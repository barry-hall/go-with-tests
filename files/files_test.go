package files_test

import (
	"go-with-tests/files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Go
is
awesome`
	)

	fs := fstest.MapFS{
		"hello_world.md":   {Data: []byte(firstBody)},
		"hello_world_2.md": {Data: []byte(secondBody)},
	}

	posts, err := files.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], files.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})

}

func assertPost(t *testing.T, got files.Post, want files.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
