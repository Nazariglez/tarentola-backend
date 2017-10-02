// Created by nazarigonzalez on 2/10/17.

package logger

import (
	_logger "github.com/nazariglez/logger"
)

var Log = func() *_logger.Logger {
	l := _logger.New()
	l.SetLevel(_logger.TRACE)

	err := l.EnableFileOutput("tarentola", "./logs", _logger.TRACE)
	if err != nil {
		l.DisableFileOutput()
		l.Error(err.Error())
	}

	return l
}()
