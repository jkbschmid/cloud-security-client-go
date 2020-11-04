// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Cloud Security Client Go contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sap-staging/cloud-security-client-go/auth"
	"github.com/sap-staging/cloud-security-client-go/env"
	"log"
	"net/http"
	"os"
)

// Main class for demonstration purposes.
func main() {
	r := mux.NewRouter()

	config, err := env.GetIASConfig()
	if err != nil {
		panic(err)
	}
	authMiddleware := auth.NewAuthMiddleware(config, auth.Options{
		UserContext:  "user",
		ErrorHandler: nil,
	})
	r.Use(authMiddleware.Handler)

	r.HandleFunc("/helloWorld", helloWorld).Methods("GET")

	address := ":8080"
	log.Println("Starting server on address", address)
	err = http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, r))
	if err != nil {
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*auth.OIDCClaims)
	_, _ = w.Write([]byte(fmt.Sprintf("Hello world!\nYou're logged in as %s", user.Email)))
}
