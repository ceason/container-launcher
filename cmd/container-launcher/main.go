package main

import (
	"syscall"
	"os"
	"strings"
	"github.com/golang/glog"
	"fmt"
	"io/ioutil"
	"flag"
	"path/filepath"
)

const (
	itemDelimiter = "\n"
	kvDelimiter   = "="
)

func main() {
	flag.Parse()
	println("Launching this with special launcher!")

	// get env vars
	env := os.Environ()
	println("wtf1")
	if val, ok := os.LookupEnv("CONTAINERLAUNCHER_ENVIRONMENT"); ok {
		for _, item := range strings.Split(val, itemDelimiter) {
			if item == "" {
				continue
			}
			parts := strings.SplitN(item, kvDelimiter, 2)
			if len(parts) != 2 {
				glog.Fatalf("Must get environment vars in format <name>=<value> but got '%s'", item)
			}
			name := parts[0]
			value := getValue(parts[1])
			env = append(env, fmt.Sprintf("%s=%s", name, value))
		}
	}

	println("wtf2")
	// get files
	if val, ok := os.LookupEnv("CONTAINERLAUNCHER_FILES"); ok {
		for _, item := range strings.Split(val, itemDelimiter) {
			if item == "" {
				continue
			}
			parts := strings.SplitN(item, kvDelimiter, 2)
			filename := parts[0]
			content := getValue(parts[1])
			err := os.MkdirAll(filepath.Dir(filename), 0755)
			if err != nil {
				glog.Fatalf("Error creating dir '%s'", err.Error())
			}
			err = ioutil.WriteFile(filename, []byte(content), 0644)
			if err != nil {
				glog.Fatalf("Error writing file '%s'", err.Error())
			}
		}
	}

	println("wtf3")
	args := os.Args[1:]
	println(fmt.Sprintf("Execiting '%s' with args '%v' and environment '%v'", args[0], args, env))
	glog.Infof("Execiting '%s' with args '%v' and environment '%v'", args[0], args, env)
	syscall.Exec(args[0], args, env)
}

func getValue(url string) string {
	parts := strings.SplitN(url, ":", 2)
	if len(parts) != 2 {
		glog.Fatalf("Value must be in format <scheme>:<value>, got '%s'", url)
	}
	scheme := parts[0]
	value := parts[1]
	switch scheme {
	case "content":
		return value
	default:
		glog.Fatalf("Unrecognized scheme '%s' in url '%s'", scheme, url)
	}
	panic("Unreachable")
}
