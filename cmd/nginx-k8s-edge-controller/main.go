// Copyright 2023 f5 Inc. All rights reserved.
// Use of this source code is governed by the Apache
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/nginxinc/kubernetes-nginx-ingress/internal/observation"
	"github.com/nginxinc/kubernetes-nginx-ingress/internal/synchronization"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	err := run()
	if err != nil {
		logrus.Error(err)
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)
	<-signalChan
}

func run() error {
	ctx := context.Background()

	synchronizer, err := synchronization.NewSynchronizer()
	if err != nil {
		return fmt.Errorf(`error initializing synchronizer: %w`, err)
	}

	err = synchronizer.Initialize()
	if err != nil {
		return fmt.Errorf(`error initializing synchronizer: %w`, err)
	}

	handler := observation.NewHandler(synchronizer)
	handler.Initialize()

	watcher, err := observation.NewWatcher(ctx, handler)
	if err != nil {
		return fmt.Errorf(`error occurred creating a watcher: %w`, err)
	}

	err = watcher.Initialize()
	if err != nil {
		return fmt.Errorf(`error occurred initializing the watcher: %w`, err)
	}

	err = watcher.Watch()
	if err != nil {
		return fmt.Errorf(`error occurred watching for events: %w`, err)
	}

	return nil
}
