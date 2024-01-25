package files_test

import (
	"go-with-tests/files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":   {Data: []byte("Title: Post 1")},
		"hello_world_2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := files.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := files.Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
