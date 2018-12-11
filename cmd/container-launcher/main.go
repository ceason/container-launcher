package main

import (
	"syscall"
	"github.com/golang/glog"
	"flag"
	"github.com/ceason/container-launcher"
)

func main() {
	flag.Parse()
	glog.Infof("Launching this with special launcher!")

	launcher, err := container_launcher.NewFromEnvironment()
	if err != nil {
		glog.Fatalf(err.Error())
	}
	argv0, argv, envv, err := launcher.GetExecArgs()
	if err != nil {
		glog.Fatalf(err.Error())
	}
	glog.Infof("Execiting '%s' with args '%v' and environment '%v'", argv0, argv, envv)
	syscall.Exec(argv0, argv, envv)

}
