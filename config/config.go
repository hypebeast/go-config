// Copyright 2014 Sebastian Ruml <sebastian.ruml@gmail>. All rights reserved.

package config

import (
	"bytes"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io"
	"io/ioutil"
	"os"
	"path"
)

type Config map[string]interface{}
type Configs map[string]interface{}

var basePath string
var env string
var configs Configs

// Init sets the base directory (the directory that contains the configuration files)
// and the name of the environment variable that defines the active environment.
// If no environment variable is specified no cascading files will be loaded.
func Init(baseDir string, environment ...string) {
	basePath = baseDir
	configs = make(Configs)
	if len(environment) == 1 {
		env = os.Getenv(environment[0])
	} else {
		env = ""
	}
}

// Get reads the config options for the given domain and writes the config options
// in the given struct. The rawValu must be a pointer to a structure that represents
// the structure of the JSON.
func Get(domain string, rawVal interface{}) (err error) {
	if _, ok := configs[domain]; ok {
		rawVal = configs[domain]
		return nil
	}

	var baseFile io.Reader
	var domainFile io.Reader
	baseFile, err = load(domain + ".json")
	if err != nil {
		return
	}

	var baseConfig Config
	var domainConfig Config
	baseConfig, err = marshalReader(baseFile)
	if err != nil {
		return
	}

	if env != "" {
		domainConfFound := true
		domainFile, err = load(domain + "." + env + ".json")
		if err != nil {
			domainConfFound = false
		}

		if domainConfFound {
			domainConfig, err = marshalReader(domainFile)
			if err != nil {
				return
			}
		}
	}

	err = marshal(baseConfig, domainConfig, rawVal)
	if err != nil {
		return
	}

	if _, ok := configs[domain]; !ok {
		configs[domain] = rawVal
	}

	return
}

// load loads the config file with the given file name from disk.
func load(filename string) (r io.Reader, err error) {
	if _, err = os.Stat(path.Join(basePath, filename)); err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(path.Join(basePath, filename))
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(file), nil
}

// marshalReader reads the config from the reader, unmarshals it to a map and
// returns it.
func marshalReader(r io.Reader) (confMap Config, err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	confMap = make(Config)
	err = json.Unmarshal(buf.Bytes(), &confMap)
	if err != nil {
		return nil, err
	}
	return confMap, nil
}

// marshal merges config a and b and decodes the merge config into the provided
// structure.
func marshal(a Config, b Config, rawVal interface{}) (err error) {
	err = mapstructure.Decode(&a, rawVal)
	if err != nil {
		return
	}

	if b != nil {
		err = mapstructure.Decode(&b, rawVal)
		if err != nil {
			return
		}
	}

	return
}
