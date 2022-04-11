/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/securisec/enumerate/cmd"

func main() {
	// TODO check for sts get-caller identity here
	// If caller identity fails, ask for optional flag to override
	cmd.Execute()
}
