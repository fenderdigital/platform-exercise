package handlers

import (
	"context"
	"fmt"
	"net/http"

	"platform-exercise/service/business/auth"
	"platform-exercise/service/foundation/web"

	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/trace"
)

type operHandlers struct {
	db   *sqlx.DB
	auth *auth.Auth
}

func (o *operHandlers) register(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.oper.register")
	defer span.End()

	r.ParseForm()
	name := r.Form["name"][0]
	email := r.Form["email"][0]
	password := r.Form["password"][0]
	if name == "" || email == " " || password == "" {
		return web.NewRequestError(fmt.Errorf("must provide name and email and password"), http.StatusPartialContent)
	}

	return nil
}

func (o *operHandlers) login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.oper.login")
	defer span.End()

	return nil
}
