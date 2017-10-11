// Created by nazarigonzalez on 11/10/17.

package content

import (
	"fmt"
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/logger"
	"net/http"
	"os"
)

func Serve() {
	if !config.Data.Static.Enabled {
		return
	}

	if _, err := os.Stat(config.Data.Static.Path); os.IsNotExist(err) {
		err := os.Mkdir(config.Data.Static.Path, 0777)
		if err != nil {
			logger.Log.Error(err.Error())
		}

		logger.Log.Debugf("Created folder to serve statics '%s'", config.Data.Static.Path)
	}

	fs := http.FileServer(http.Dir(config.Data.Static.Path))
	http.Handle("/", fs)

	port := fmt.Sprintf(":%d", config.Data.Static.Port)

	logger.Log.Logf("Serving Static Files on 127.0.0.1%s...", port)
	logger.Log.Fatal(http.ListenAndServe(port, nil))
}
