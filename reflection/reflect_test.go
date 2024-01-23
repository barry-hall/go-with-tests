package reflect

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"River"},
			[]string{"River"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"River", "London"},
			[]string{"River", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"River", 2},
			[]string{"River"},
		},
		{
			"Nested fields",
			Person{
				"River",
				Profile{1, "Manchester"},
			},
			[]string{"River", "Manchester"},
		},
		{
			"Pointers to things",
			&Person{
				"River",
				Profile{1, "Manchester"},
			},
			[]string{"River", "Manchester"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Manchester"},
			},
			[]string{"London", "Manchester"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Manchester"},
			},
			[]string{"London", "Manchester"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "London"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "London"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "London"}
		}

		var got []string
		want := []string{"Berlin", "London"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
