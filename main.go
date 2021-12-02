package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/francescomari/wsk-use/internal/config"
	"github.com/francescomari/wsk-use/internal/openwhisk"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	return runArgs(os.Args[1:]...)
}

func runArgs(args ...string) error {
	switch len(args) {
	case 0:
		return printContexts()
	case 1:
		return switchContexts(args[0])
	default:
		return fmt.Errorf("more than one context name is provided")
	}
}

func printContexts() error {
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("read configuration: %v", err)
	}

	if cfg == nil {
		return nil
	}

	var names []string

	for name := range cfg.Contexts {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}

func switchContexts(contextName string) error {
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("read configuration: %v", err)
	}

	if cfg == nil {
		return fmt.Errorf("configuration does not exist")
	}

	for name, context := range cfg.Contexts {
		if name == contextName {
			return switchContext(context)
		}
	}

	return fmt.Errorf("context does not exist")
}

func switchContext(context config.Context) error {
	config := openwhisk.Config{
		Auth:    context.Auth,
		APIHost: context.APIHost,
	}

	if err := openwhisk.WriteConfig(&config); err != nil {
		return fmt.Errorf("write OpenHisk configuration: %v", err)
	}

	return nil
}
