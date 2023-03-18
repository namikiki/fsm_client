package types

type AddWatch struct {
	SyncID string
	Path   string
	IsAdd  bool
	Ignore bool
}
