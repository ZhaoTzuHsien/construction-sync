package badger

import (
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/constant"
	"github.com/dgraph-io/badger/v3"
	"os"
	"path/filepath"
)

func Open() *badger.DB {
	// Get user cache dir
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	// Create default options and disable logger
	opts := badger.DefaultOptions(filepath.Join(cacheDir, constant.APP_NAME, "badger"))
	opts.Logger = nil

	// open database
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return db
}
