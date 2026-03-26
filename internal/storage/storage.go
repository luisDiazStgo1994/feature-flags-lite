package storage

type Storage interface {
	UpsertFeatureFlag() error
	GetFeatureFlagsByOrgId() error
	GetFeatureFlagsByOrgUserId() error
}
