package endpoints

import (
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			render.Status(r, http.StatusForbidden)
			render.JSON(w, r, map[string]string{"error": "Request does not contain authorization"})
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/provider")

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "Error to connect to provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
		_, err = verifier.Verify(r.Context(), token)

		if err != nil {
			render.Status(r, http.StatusForbidden)
			render.JSON(w, r, map[string]string{"error": "Request does not contain authorization"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
