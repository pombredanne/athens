package rdbms

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gomods/athens/pkg/storage"
	"github.com/gomods/athens/pkg/storage/rdbms/models"
)

// Get a specific version of a module
func (r *ModuleStore) Get(module, vsn string) (*storage.Version, error) {
	result := models.Module{}
	query := r.conn.Where("module = ?", module).Where("version = ?", vsn)
	if err := query.First(&result); err != nil {
		return nil, err
	}
	return &storage.Version{
		RevInfo: storage.RevInfo{
			Version: result.Version,
			Name:    result.Version,
			Short:   result.Version,
			Time:    time.Now(),
		},
		Mod: result.Mod,
		Zip: ioutil.NopCloser(bytes.NewReader(result.Zip)),
	}, nil
}
