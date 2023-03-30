package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

func gitDir(u *url.URL) (string, error) {
	cleanPath := path.Clean(u.Path)
	pathSlice := strings.Split(cleanPath, "/")
	if len(pathSlice) < 3 {
		return "", errors.New("a forge URL should have at least a user and a project")
	}
	return path.Join(os.Getenv("HOME"), "src", u.Host, pathSlice[1]), nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("requires a git URL as an argument")
	}

	arg := os.Args[1]
	u, err := url.Parse(arg)
	if err != nil {
		log.Fatalf("error %T parsing url: %v", err, err)
	}
	dir, err := gitDir(u)
	if err != nil {
		log.Fatalf("error parsing URL: %v", err)
	}
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("error %T creating directory: %v", err, err)
	}
	err = os.Chdir(dir)
	if err != nil {
		log.Fatalf("error changing to directory %q: %v", dir, err)
	}
	clone := exec.Command("git", "clone", u.String())
	clone.Stdout = os.Stdout
	clone.Stderr = os.Stderr
	err = clone.Run()
	if err != nil {
		log.Fatalf("error %T cloning repo %q: %v", err, arg, err)
	}
	fmt.Println(dir)
}
