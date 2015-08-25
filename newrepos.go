package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: newrepos <path/to/repo> \n\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func repoFullPath(path string) string {
	return fmt.Sprintf("%s/src/%s", os.Getenv("GOPATH"), path)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}
	path := flag.Arg(0)

	if err := os.MkdirAll(repoFullPath(path), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "create directory failed: %s\n", err)
		os.Exit(1)
	}

	o, err := exec.Command("git", "init", repoFullPath(path)).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize repository: %s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s\n", string(o))
}
