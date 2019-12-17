package main

import (
	"log"
	"regexp"
	"net/http"
	"encoding/json"
	"encoding/base64"

	"github.com/gorilla/mux"

	"gopkg.in/macaroon-bakery.v2/bakery"

	"github.com/socketworld/user/ledger"

)

func Get (w http.ResponseWriter, r *http.Request) {
	// Load request variables.
	vars := mux.Vars(r)

	// Load the User ledger.
	users := ledger.FetchUsers()

	// Load the requested User from the User ledger.
	user, ok := users[vars[`name`]];
	if !ok {
		// Since the User does not exists, it cannot be fetched.
		log.Printf("Error: GET %s when user does not exist.", vars["name"])

		// Print Not Found since the User does not exist.
		w.WriteHeader(http.StatusNotFound)

		return

	}

	// Generate JSON for the new User.
	user_json, _ := json.Marshal(user)

	// Log Success for the User retrieval.
	log.Printf("Success: GET %s %s", user.Name, user_json)

	// Define Authorization token format.
	auth_token_re := regexp.MustCompile(`^Bearer ([a-zA-Z0-9+/]+={0,2})$`)

	// Parse the Authorization token.
	auth_token_parsed := auth_token_re.FindStringSubmatch(r.Header.Get(`Authorization`))
	auth_json, err_auth_json := base64.StdEncoding.DecodeString(auth_token_parsed[1])
	if err_auth_json != nil {
		// Failed to generate Authorization JSON!
		log.Printf("Error: Cannot decode Authorization token (%s)", err_auth_json)

		// Report Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	// Parse Authorization JSON into Macaroon object.
	auth := bakery.Macaroon{}
	err_auth := auth.UnmarshalJSON(auth_json)
	if err_auth != nil {
		// Failed to generate Authorization JSON!
		log.Printf("Error: Cannot load Authorization JSON (%s)", err_auth)

		// Report Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	// Write 200 and print the JSON of the retrieved User.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user_json))

}
