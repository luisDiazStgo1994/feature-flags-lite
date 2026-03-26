package cache

type Cache interface {
	GetFeatureFlagsByOrgId() error
	GetFeatureFlagsByOrgUserId() error
	UpsertFeatureFlag() error
}
