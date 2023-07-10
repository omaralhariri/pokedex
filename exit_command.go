package main

import "os"

func exitCommand(cfg *config, area_name string) error {
	os.Exit(0)
	return nil
}
