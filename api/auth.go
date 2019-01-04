package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
	"net/http"
	"os"
	"strings"
)

func (c *Controller) JwtAuthentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/login"}                       //List of endpoints that doesn't require auth
		specialPrivileges := []string{"/api/user/new", "/api/users"} //List of endpoint that need special admin privileges
		requestPath := r.URL.Path                                    //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header
		if tokenHeader == "" {                       //Token is missing, returns with error code 403 Unauthorized
			forbidden(w, errors.New("Missing auth token"))
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			forbidden(w, errors.New("Invalid/Malformed token"))
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &model.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv(database.JWT_PASSWORD)), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			forbidden(w, errors.New(fmt.Sprint("Token is expired or signature is invalid.")))
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			forbidden(w, errors.New("Token is not valid"))
			return
		}

		for _, v := range specialPrivileges {
			if strings.Contains(requestPath, v) {
				if !tk.Admin {
					forbidden(w, errors.New(fmt.Sprint("You don't have admin privileges!")))
					return
				}
			}
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		//fmt.Sprintf("User %", tk.UserId) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
