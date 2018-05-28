package main

import "fmt"

func InstallWebFiles(service string) {
	if service == "" {

	}

	switch service {
	case "nginx":
		{
			InstallNginx()
		}
	case "apache":
		{
			InstallApache()
		}
	default:
		{
			fmt.Printf("Unable to configure web service")
		}
	}
}

func InstallNginx() {

}

func InstallApache() {

}
