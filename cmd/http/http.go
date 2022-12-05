package http

import (
	"net/http"

	bootstrap "github.com/awi93/wallet-service/config"
	"github.com/awi93/wallet-service/src/endpoints"
	"github.com/spf13/cobra"
)

var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "Http Service",
	RunE: func(cmd *cobra.Command, args []string) error {
		bootstrap.ReadConfig("config.yml")
		endpoints.InitDependency(bootstrap.GetConfig())
		defer endpoints.DepoEmitter.Finish()

		router := endpoints.Router()
		return http.ListenAndServe(":"+bootstrap.GetConfig().GetString("http.port"), router)
	},
}
