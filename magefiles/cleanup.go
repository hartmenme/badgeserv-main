package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"path"
)

var goCmds []string


// Clean deletes build output and cleans up the working directory.
// Clean deletes build output and cleans up the working directory.
func Clean() error {
	for _, name := range goCmds {
		targetPath := path.Join(binDir, name)
		fmt.Println("Removing file:", targetPath)
		if err := sh.Rm(targetPath); err != nil {
			fmt.Println("Error removing file:", err)
			return err
		}
	}

	for _, name := range outputDirs {
		fmt.Println("Removing directory:", name)
		if err := sh.Rm(name); err != nil {
			fmt.Println("Error removing directory:", err)
			return err
		}
	}
	fmt.Println("Clean completed successfully")
	return nil
}

func CleanGoLandDefaultBuild() error {
	binaryName := "badgeserv" // Update this to the actual binary name
	binaryPath := path.Join(".", binaryName)

	fmt.Println("Removing GoLand default build binary:", binaryPath)
	if err := sh.Rm(binaryPath); err != nil {
		fmt.Println("Error removing binary:", err)
		return err
	}

	fmt.Println("GoLand default build binary removed successfully")
	return nil
}

