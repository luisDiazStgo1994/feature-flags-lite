package api

import "net/http"

type Handlers struct {
	GetFeatureFlagsByOrgId     http.HandlerFunc
	UpsertFeatureFlag          http.HandlerFunc
	GetFeatureFlagsByOrgUserId http.HandlerFunc
}

func NewHandlers() *Handlers {
	return &Handlers{
		GetFeatureFlagsByOrgId:     GetFeatureFlagsByOrgId,
		UpsertFeatureFlag:          UpsertFeatureFlag,
		GetFeatureFlagsByOrgUserId: GetFeatureFlagsByOrgUserId,
	}
}

func GetFeatureFlagsByOrgId(w http.ResponseWriter, r *http.Request) {
	orgId := r.PathValue("orgIds")
	// mock org service validation
	if !IsValidOrg(orgId) {
		return
	}
}

func UpsertFeatureFlag(w http.ResponseWriter, r *http.Request) {
	orgId := r.PathValue("orgIds")
	// mock org service validation
	if !IsValidOrg(orgId) {
		return
	}
}

func GetFeatureFlagsByOrgUserId(w http.ResponseWriter, r *http.Request) {
	orgId := r.PathValue("orgIds")
	userId := r.PathValue("userId")

	// mock org service validation
	if !IsValidOrg(orgId) {
		return
	}

	//mock user service validation
	if !IsValidUser(userId) {
		return
	}

}

func IsValidOrg(orgId string) bool {
	return true
}

func IsValidUser(userId string) bool {
	return true
}
