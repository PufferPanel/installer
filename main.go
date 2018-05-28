package main

import "flag"

func main() {
	var noDaemon bool
	var configureWeb string
	var version bool
	var install bool

	flag.BoolVar(&noDaemon, "noDaemon", false, "Do not install daemon")
	flag.StringVar(&configureWeb, "configureWeb", "", "Configures the specified web service")
	flag.BoolVar(&version, "version", false, "Gets version of installer and panel")
	flag.BoolVar(&install, "install", false, "Installs the panel")

	flag.Parse()

	if version {
		GetVersion()
		return
	}

	if install {
		InstallPanel(configureWeb)
		return
	}
}
