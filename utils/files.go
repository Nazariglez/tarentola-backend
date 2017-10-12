// Created by nazarigonzalez on 12/10/17.

package utils

import (
	"github.com/nazariglez/tarentola-backend/config"
	"io/ioutil"
	"os"
	"path"
)

func CreateStaticFile(filepath, name string, file []byte) error {
	p := path.Join(config.Data.Static.Path, filepath)
	if err := os.MkdirAll(path.Join(config.Data.Static.Path, filepath), os.ModePerm); err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(p, name), file, os.ModePerm)
}
