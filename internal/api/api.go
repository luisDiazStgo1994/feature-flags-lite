package api

import (
	"feature-flags-lite/internal/cache"
	"feature-flags-lite/internal/storage"
	"net/http"
)

type api struct {
	Mux   http.ServeMux
	repo  *storage.Storage
	cache *cache.Cache
}

func newApi(handlers *Handlers) *api {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /org/{orgId}", handlers.GetFeatureFlagsByOrgId)
	mux.HandleFunc("GET /org/{orgId}/user/{userId}", handlers.GetFeatureFlagsByOrgUserId)
	mux.HandleFunc("POST /org/{orgID}", handlers.UpsertFeatureFlag)

	return &api{}
}
