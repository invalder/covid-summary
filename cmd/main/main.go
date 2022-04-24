package main

import (
	"fmt"

	"github.com/invalder/covid-summary/pkg/controllers"
	"github.com/invalder/covid-summary/pkg/routes"
)

func main() {
	fmt.Println("Hello World")
	controllers.InitializePatients()

	r := routes.SetupRouter()
	r.Run(":8080")
}
