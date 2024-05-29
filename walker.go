package main

import "os"

// work into any directory

func walker(dirName string) error {
	err := os.Chdir(dirName)
	if err != nil {
		return err
	}

	return nil
}
