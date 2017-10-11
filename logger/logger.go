// Created by nazarigonzalez on 2/10/17.

package logger

import (
	_logger "github.com/nazariglez/logger"
	"github.com/nazariglez/tarentola-backend/config"
	"os"
	"strings"
)

var Log = func() *_logger.Logger {
	var l *_logger.Logger

	if config.Data.Logger.Async {
		l = _logger.NewAsync()
	} else {
		l = _logger.New()
	}

	l.SetLevel(_logger.LogLevel(config.Data.Logger.Level))

	if config.Data.Logger.File {
		if _, err := os.Stat(config.Data.Logger.Path); os.IsNotExist(err) {
			err := os.Mkdir(config.Data.Logger.Path, 0777)
			if err != nil {
				l.Error(err.Error())
			}

			l.Debugf("Created folder to save logs '%s'", config.Data.Logger.Path)
		}

		err := l.EnableFileOutput(strings.ToLower(config.Data.Name), config.Data.Logger.Path, _logger.LogLevel(config.Data.Logger.FileLevel))
		if err != nil {
			l.DisableFileOutput()
			l.Error(err.Error())
		}
	}

	return l
}()
