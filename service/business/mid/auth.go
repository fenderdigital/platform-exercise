package mid

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"platform-exercise/service/business/auth"
	"platform-exercise/service/foundation/web"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
)

// ErrForbidden is returned when an authenticated user does not have a
// sufficient role for an action.
var ErrForbidden = web.NewRequestError(
	errors.New("you are not authorized for that action"),
	http.StatusForbidden,
)

// Authenticate validates a JWT from the `Authorization` header.
func Authenticate(a *auth.Auth) web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(after web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "internal.mid.authenticate")
			defer span.End()

			// Parse the authorization header. Expected header is of
			// the format `Bearer <token>` or cookie.
			parts := strings.Split(r.Header.Get("Authorization"), " ")
			c, err := r.Cookie("session_token")

			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" && err != nil {
				err := errors.New("expected authorization header format: Bearer <token> or cookie is not set")
				return web.NewRequestError(err, http.StatusUnauthorized)
			}

			// Start a span to measure just the time spent in ParseClaims.
			var claims auth.Claims
			if len(parts) == 2 {
				claims, err = a.ValidateToken(parts[1])
				fmt.Println(claims, err)
				if err != nil {
					return web.NewRequestError(err, http.StatusUnauthorized)
				}
			} else {
				claims, err = a.ValidateToken(c.Value)
				if err != nil {
					return web.NewRequestError(err, http.StatusUnauthorized)
				}
			}

			/*


				c, err := r.Cookie("session_token")
				if err != nil {
					// If the cookie is not set, return an unauthorized status
					err := fmt.Errorf("cookie is not set :%s", err.Error())
					return web.NewRequestError(err, http.StatusUnauthorized)
				}

					//TODO :
					// Parse the authorization header. Expected header is of
					// the format `Bearer <token>`.
					parts := strings.Split(r.Header.Get("Authorization"), " ")
					if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
						err := errors.New("expected authorization header format: Bearer <token>")
						return web.NewRequestError(err, http.StatusUnauthorized)
					}

					// Start a span to measure just the time spent in ParseClaims.
					claims, err := a.ValidateToken(parts[1])
					if err != nil {
						return web.NewRequestError(err, http.StatusUnauthorized)
					}
			*/

			// Add claims to the context so they can be retrieved later.
			ctx = context.WithValue(ctx, auth.Key, claims)

			return after(ctx, w, r)
		}

		return h
	}

	return m
}

// HasRole validates that an authenticated user has at least one role from a
// specified list. This method constructs the actual function that is used.
func Authorized(roles ...string) web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(after web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "internal.mid.hasrole")
			defer span.End()

			claims, ok := ctx.Value(auth.Key).(auth.Claims)
			if !ok {
				return errors.New("claims missing from context: HasRole called without/before Authenticate")
			}

			if !claims.Authorized(roles...) {
				return ErrForbidden
			}

			return after(ctx, w, r)
		}

		return h
	}

	return m
}
