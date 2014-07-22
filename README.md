# go-config

Configuration file parser for JSON format.

go-config makes it simple to work with cascading configuration files. It allows you to use cascading configuration files (e.g. one configuration for every environment). You can define options for every environment in different configuration files and load them accordingly to the set environment.

**Build Status**: [![Build Status](https://travis-ci.org/hypebeast/go-config.svg?branch=master)](https://travis-ci.org/hypebeast/go-config)

## Usage

Place your configuration files in a directory (e.g. config). Configuration files need to be valid JSON files, terminated by .json.

The following example loads and parses options from *base.json* in the *config* directory:

```go
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
```

### Environment Specific Files and Cascading

Inside the configuration directory you can have configurations for different environments.

**Example**

You can have your basic options in *mongo.json* that is loaded when no environment is specified and a *mongo.stage.json* file that is loaded when the environment variable (the name of the environment variable can be set during initialization) is set to *stage*.

Configuration options in *mongo.json* will be merged with the options from *mongo.stage.json*.

## Example

The follwoing examples shows the content *mongo.js*:

```json
{
    "Host": "localhost",
    "Port": 27017
}
```

And you have a *mongo.stage.json* file with the following content:

```json
{
    "Host": "mongo-stage.host.com"
}
```

The resulting configuration for the *stage* environment will be:

```json
{
    "Host": "mongo-stage.host.com",
    "Port": 27017
}
```

## Credits

This library was inspired by [konphyg](https://github.com/pgte/konphyg).

## License

The MIT License (MIT)

Copyright (c) 2014 Sebastian Ruml

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
