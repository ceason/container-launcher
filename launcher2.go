package container_launcher

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

var (
	handlers           []SecretResolver
	launcherPrefix     = flag.String("containerlauncher-prefix", "containerlauncher", "Environment variable values starting with '<containerlauncher-prefix>:' will be resolved")
	NoMatchingResolver = errors.New("no matching SecretResolver")
)

// syntax is... "launcherPrefix:[optionalFilePath]:valueToResolve"

type SecretResolver interface {
	IsDefinedAt(str string) bool
	Resolve(str string, w io.Writer) error
	UsageText() string
}

func RegisterResolver(h SecretResolver) {
	handlers = append(handlers, h)
}

// Get help/usage text for all registered handlers
func GetRegisteredResolvers() (usageText []string) {
	for _, h := range handlers {
		usageText = append(usageText, h.UsageText())
	}
	return
}

// resolves using the first matching handler
func resolve(str string, w io.Writer) error {
	// find first matching handler
	for _, h := range handlers {
		if h.IsDefinedAt(str) {
			return h.Resolve(str, w)
		}
	}
	return NoMatchingResolver
}

// Return `os.Environ` with all containerlauncher references resolved.
// Will re-resolve each time it's invoked, so you probably want to reuse the
// result instead of calling this multiple times.
func Environ() ([]string, error) {
	var resolved []string
	for _, v := range os.Environ() {
		parts := strings.SplitN(v, "=", 1)
		name := parts[0]
		value := parts[1]
		vparts := strings.SplitN(value, ":", 2)
		prefix := vparts[0]
		// skip stuff that doesn't match the required prefix
		if prefix != *launcherPrefix {
			resolved = append(resolved, v)
			continue
		}
		if len(vparts) != 3 {
			return nil, errors.New(fmt.Sprintf("Expected format '<launcherPrefix>:[optionalFilePath]:...' but got '%s' for var '%s'", value, name))
		}
		filePath := vparts[1]
		resolvableStr := vparts[2]
		// resolve to an env var if no file path was given
		if filePath == "" {
			buf := strings.Builder{}
			err := resolve(resolvableStr, &buf)
			if err != nil {
				return nil, err
			}
			// set env var to the resolved value
			resolved = append(resolved, fmt.Sprintf("%s=%s", name, buf.String()))
		} else {
			err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
			if err != nil {
				return nil, err
			}
			f, err := os.Create(filePath)
			if err != nil {
				return nil, err
			}
			err = resolve(resolvableStr, f)
			if err != nil {
				return nil, err
			}
			// file has been written, so set env var to the path where we wrote the file
			resolved = append(resolved, fmt.Sprintf("%s=%s", name, filePath))
		}
	}
	return resolved, nil
}
