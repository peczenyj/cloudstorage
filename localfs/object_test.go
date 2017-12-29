package localfs_test

import (
	"testing"

	"github.com/lytics/cloudstorage"
	"github.com/lytics/cloudstorage/localfs"
	"github.com/lytics/cloudstorage/testutils"
)

var config = &cloudstorage.Config{
	Type:       "localfs",
	AuthMethod: localfs.AuthFileSystem,
	LocalFS:    "/tmp/mockcloud",
	TmpDir:     "/tmp/localcache",
}

func TestAll(t *testing.T) {
	store, err := cloudstorage.NewStore(config)
	if err != nil {
		t.Fatalf("Could not create store: config=%+v  err=%v", config, err)
		return
	}
	testutils.RunTests(t, store)
}