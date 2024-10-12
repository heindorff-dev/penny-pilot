package main

import (
	"api/pkg/database"
	"api/pkg/helpers"
	"api/pkg/router"
	"log"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func main() {
	database.Init()
	cassandraSession := database.Session
	defer cassandraSession.Close()

	issuerURL, _ := url.Parse(helpers.SafeGetEnv("AUTH0_DOMAIN"))
	audience := helpers.SafeGetEnv("AUTH0_AUDIENCE")
	provider := jwks.NewCachingProvider(issuerURL, time.Duration(5*time.Minute))
	jwtValidator, _ := validator.New(provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)
	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	r := router.New(jwtMiddleware)

	log.Print("Server listening on http://localhost:3001/")
	if err := http.ListenAndServe("0.0.0.0:3001", r); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
