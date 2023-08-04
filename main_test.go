package main

import (
	"net/url"
	"testing"
)

func TestPaths(t *testing.T) {
	gotExp := "got:%v expected:%v"

	type test struct {
		argURL       string
		truncatedURL *url.URL
		gitPath      string
	}

	goodURL, err := url.Parse("https://example.com/one/two")
	if err != nil {
		t.Fatal(err)
	}

	cases := []test{
		test{
			"https://example.com/one/two/three",
			goodURL,
			"example.com/one/two",
		},
		test{
			"https://example.com/one/two",
			goodURL,
			"example.com/one/two",
		},
	}

	for _, c := range cases {
		t.Log(c.argURL)
		u, err := url.Parse(c.argURL)
		if err != nil {
			t.Fatal(err)
		}
		gitU, gitP, err := paths(u)
		if err != nil {
			t.Fatal(err)
		}
		if gitU.String() != c.truncatedURL.String() {
			t.Fatalf(gotExp, gitU.String(), c.truncatedURL.String())
		}
		if gitP != c.gitPath {
			t.Fatalf(gotExp, gitP, gitP)
		}
	}

	// fail on truncated url
	u, err := url.Parse("https://example.com/insufficient")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
	_, _, err = paths(u)
	if err == nil {
		t.Fatalf("%v should have caused error, got nil", u.String())
	}
}
