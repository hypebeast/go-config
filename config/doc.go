// Copyright 2014 Sebastian Ruml <sebastian.ruml@gmail>. All rights reserved.

/*
Configuration file parser for JSON format.

go-config makes it simple to work with cascading configuration files. It allows
you to use cascading configuration files (e.g. one configuration for every environment).
You can define options for every environment in different configuration files and
load them accordingly to the set environment.

Usage

Place your configuration files in a directory (e.g. config). Configuration files
need to be valid JSON files, terminated by .json.

The following example loads and parses options from base.json in the config directory:

    package main

    import (
        "fmt"
        "github.com/hypebeast/go-config/config"
    )

    type BaseConfig struct {
        Host    string
        Port    int
    }

    func main() {
        // Initialize the config system
        config.Init("config", "")

        // Get the base config
        config.Get("base", &baseConfig)
    }

*/

package config
