// Copyright 2016 someonegg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gologf can control log-file rotation of "someonegg/golog".
//
// The simplest way to use gologf is:
//   package main
//
//   import (
//   	"github.com/someonegg/goutility/gologf"
//   )
//
//   func main() {
//      ...
//   	err = gologf.SetOutput(logFilePath)
//   	...
//   }
//
// A simple rotation sh is like:
//   nowtime=$(date "+%Y%m%d%H%M%S")
//   mv -f xxx.log ${nowtime}.xxx.log
//   kill -s SIGUSR1 pid_of_xxx
//
package gologf

import (
	"github.com/someonegg/golog"
	"github.com/someonegg/golog/handler/hjson"
	"github.com/someonegg/golog/handler/hmulti"
	"os"
	"sync"
)

var (
	locker sync.Mutex
	logS   string
	logF   *os.File
)

func SetOutput(path string) error {
	file, err := os.OpenFile(path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	golog.RootLogger.SetHandler(hmulti.New(
		golog.NewHandler(os.Stderr), hjson.New(file)))

	locker.Lock()
	defer locker.Unlock()

	if logF != nil {
		logF.Close()
	}

	logS = path
	logF = file

	return nil
}
