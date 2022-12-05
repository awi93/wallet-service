package http

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/awi93/wallet-service/config"
	"github.com/awi93/wallet-service/src/endpoints"
	"github.com/spf13/cobra"
)

var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "Http Service",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Reading Configuration File")
		config.ReadConfig("config.yml")
		log.Println("Configuration File Read")
		log.Println("Initializing Dependency")
		endpoints.InitDependency(config.GetConfig())
		log.Println("Dependency Initialized")
		defer endpoints.DepoEmitter.Finish()

		log.Println("Registering route to router")
		router := endpoints.Router()
		log.Println("Route registered")
		log.Printf("Starting Http Service at :%v\n", config.GetConfig().GetString("http.port"))
		return http.ListenAndServe(fmt.Sprintf(":%s", config.GetConfig().GetString("http.port")), router)
	},
}
