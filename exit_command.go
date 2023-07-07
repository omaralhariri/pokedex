package main

import "os"

func exitCommand(_config *Config) error {
	os.Exit(0)
	return nil
}
