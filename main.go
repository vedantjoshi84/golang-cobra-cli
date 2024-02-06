/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

The Cobra-cli package is built on top of the standard library’s flag package,
providing a higher level of abstraction while adding more features.
Cobra-cli offers features such as command hierarchy, flag and argument parsing, and subcommands,
among others, to make developing complex and feature-rich CLI applications easier.

*/

package main

import "cobra/cmd"

func main() {
	// Execute adds all child commands to the root command and sets flags appropriately.
	// This is called by main.main(). It only needs to happen once to the rootCmd.
	cmd.Execute()
}
