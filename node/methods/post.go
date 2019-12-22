package methods

import (
	"fmt"
	"log"
	"context"
	"net/http"
	"encoding/json"
	"encoding/base64"

	"github.com/gorilla/mux"
	"github.com/google/uuid"

	"gopkg.in/macaroon-bakery.v2/bakery"
	"gopkg.in/macaroon-bakery.v2/bakery/checkers"

	"github.com/socketworld/user/node/utilities"
	"github.com/socketworld/user/node/ledger"

)

func Post (w http.ResponseWriter, r *http.Request) {
	// Load request variables.
	vars := mux.Vars(r)

	// Load the User ledger
	users := ledger.FetchUsers()

	// Check for the existence of the requested User
	if _, ok := users[vars["name"]]; ok {
		// Since the User already exists, it cannot be created to again.
		log.Printf("Error: POST %s when user already exists.", vars["name"])

		// Print Method Not Allowed since the User already exists.
		w.WriteHeader(http.StatusMethodNotAllowed)

		return

	}

	// Since the User doesn't exist, this request will claim it.
	// The following process will create a new User and a token for it.

	user := ledger.User{Name: vars["name"], Id: uuid.New()}
	users[user.Name] = user

	// Create an Authorization Token for the new User.
	caveats := [...]checkers.Caveat{}

	op := bakery.Op{
		Entity: fmt.Sprintf("/%s", vars["name"]),
		Action: "Write",
	}

	auth_token, err_auth_token := utilities.FetchOven().NewMacaroon(context.Background(), bakery.Version3, caveats[:], op)
	if err_auth_token != nil {
		log.Printf("error! %s", err_auth_token)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate JSON for the new User.
	user_json, _ := json.Marshal(user)

	// Log Success for the create User request.
	log.Printf("Success: POST %s %s", user.Name, user_json)

	// Generate JSON for the Authorization token.
	auth_json, err_auth_json := auth_token.MarshalJSON()
	if err_auth_json != nil {
		// Failed to generate Authorization JSON!
		log.Printf("Error: Cannot generate Authorization JSON (%s)", err_auth_json)

		// Report Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	// Return a Base64 encoded Authorization token in the Authorization header
	w.Header().Set("Authorization", base64.StdEncoding.EncodeToString([]byte(auth_json[:])))

	// Write 200 and print the JSON of the new User.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user_json))

}
