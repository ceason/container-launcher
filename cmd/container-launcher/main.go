package main

import (
	"flag"
	"github.com/ceason/container-launcher"
	"github.com/golang/glog"
	"syscall"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	glog.V(0).Info("Launching with container-launcher")
	argv := flag.Args()
	envv, err := container_launcher.Environ()
	if err != nil {
		glog.Fatalf(err.Error())
	}
	glog.V(1).Infof("Executing '%s' with args '%v' and environment '%v'", argv[0], argv[1:], envv)
	err = syscall.Exec(argv[0], argv, envv)
	if err != nil {
		glog.Fatal(err)
	}
}
