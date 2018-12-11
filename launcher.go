package container_launcher

import (
	"os"
	"strings"
	"errors"
	"fmt"
	"path/filepath"
	"io/ioutil"
)

const (
	itemDelimiter      = "\n"
	kvDelimiter        = "="
	EnvironmentEnvName = "CONTAINERLAUNCHER_ENVIRONMENT"
	FilesEnvName       = "CONTAINERLAUNCHER_FILES"
)

type launcher struct {
	files       map[string]string
	environment map[string]string
	argv        []string
	envv        []string
}

func NewFromEnvironment() (*launcher, error) {
	// pass through all non-containerlauncher env vars
	var envv []string
	for _, value := range os.Environ() {
		if strings.HasPrefix(value, EnvironmentEnvName+"=") ||
			strings.HasPrefix(value, FilesEnvName+"=") {
			continue
		}
		envv = append(envv, value)
	}
	launcher := &launcher{
		files:       make(map[string]string),
		environment: make(map[string]string),
		argv:        os.Args[1:],
		envv:        envv,
	}
	if val, ok := os.LookupEnv(EnvironmentEnvName); ok {
		for _, item := range strings.Split(val, itemDelimiter) {
			if item == "" {
				continue
			}
			parts := strings.SplitN(item, kvDelimiter, 2)
			if len(parts) != 2 {
				return nil, errors.New(fmt.Sprintf("Must get environment vars in format <name>=<value> but got '%s'", item))
			}
			launcher.environment[parts[0]] = parts[1]
		}
	}
	if val, ok := os.LookupEnv(FilesEnvName); ok {
		for _, item := range strings.Split(val, itemDelimiter) {
			if item == "" {
				continue
			}
			parts := strings.SplitN(item, kvDelimiter, 2)
			if len(parts) != 2 {
				return nil, errors.New(fmt.Sprintf("Must get file vars in format <name>=<value> but got '%s'", item))
			}
			launcher.files[parts[0]] = parts[1]
		}
	}
	return launcher, nil
}

func (l launcher) GetExecArgs() (argv0 string, argv []string, envv []string, error error) {
	argv0 = l.argv[0]
	argv = l.argv
	envv = l.envv

	// resolve values for env vars
	for key, value := range l.environment {
		resolved, err := getValue(value)
		if err != nil {
			return
		}
		envv = append(envv, fmt.Sprintf("%s=%s", key, resolved))
	}
	// place files as appropriate
	for filename, value := range l.files {
		content, err := getValue(value)
		if err != nil {
			error = err
			return
		}
		err = os.MkdirAll(filepath.Dir(filename), 0755)
		if err != nil {
			error = errors.New(fmt.Sprintf("Error creating dir for file '%s': %s", filename, err.Error()))
			return
		}
		err = ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			error = errors.New(fmt.Sprintf("Error writing file '%s': %s", filename, err.Error()))
			return
		}
	}

	return
}

func getValue(url string) (string, error) {
	if strings.HasPrefix(url, "content:") {
		return strings.TrimPrefix(url, "content:"), nil
	}
	if strings.HasPrefix(url, "arn:aws:secretsmanager:") {
		return getValue_awsSecretsmanager(url)
	}
	if strings.HasPrefix(url, "s3://") {
		trimmed := strings.TrimPrefix(url, "s3://")
		parts := strings.SplitN(trimmed, "/", 2)
		if len(parts) != 2 {
			return "", errors.New(fmt.Sprintf("Invalid S3 URI '%s' expected 's3://<bucket>/<key>'", url))
		}
		return getValue_awsS3(parts[0], parts[1])
	}
	if strings.HasPrefix(url, "arn:aws:ssm:") {
		return getValue_awsSsm(url)
	}
	return "", errors.New(fmt.Sprintf("Unrecognized url '%s'", url))
}
