package node

import (
	"log"
	"net/http"
	"github.com/spf13/cobra"
	"github.com/gorilla/mux"
	"github.com/socketworld/user/node/methods"
)

var Command = &cobra.Command{
	Use:   "node",
	Short: "`node`",
	Long:  "`node`",

	Run: func(cmd *cobra.Command, args []string) {
		// Create a router and setup route handlers.
		router := mux.NewRouter().StrictSlash(true)
		router.Methods(`GET`).Path(`/{name}`).HandlerFunc(methods.Get)
		router.Methods(`POST`).Path(`/{name}`).HandlerFunc(methods.Post)

		// Start HTTP server, log error on exit.
		err := http.ListenAndServe(":8080", router)
		log.Fatal(err)

	},

}
