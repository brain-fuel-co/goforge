/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"goforge/cmd/goforge/cmd"
	"goforge/pkg/components/inrootdir"
)

func main() {
	if inrootdir.InRootDir() {
		cmd.Execute()
	} else {
		fmt.Println("  You must be in the root directory of a project")
		fmt.Println("  with a go.mod file in order to run goforge.")
	}
}
