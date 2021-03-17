package handlers

import (
	"context"
	"net/http"

	"platform-exercise/service/foundation/database"
	"platform-exercise/service/foundation/web"

	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/trace"
)

type check struct {
	db    *sqlx.DB
}

func (c *check) health(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.check.health")
	defer span.End()

	status := "ok"
	statusCode := http.StatusOK
	if err := database.StatusCheck(ctx, c.db); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
	}

	health := struct {
		Status  string `json:"status"`
	}{
		Status:  status,
	}
	return web.RespondJSON(ctx, w, health, statusCode)
}

func (c *check) landing(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.check.landing")
	defer span.End()

	statusCode := http.StatusOK

	landing := struct {
		Welcome string `json:"welcome"`
	}{
		Welcome: "Fender Musical Instruments Corporation",
	}

	return web.RespondJSON(ctx, w, landing, statusCode)
}
