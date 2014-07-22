// Copyright 2014 Sebastian Ruml <sebastian.ruml@gmail>. All rights reserved.

package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

type BasicConfig struct {
	Host    string
	Port    int
	ApiKey  string
	Twitter struct {
		Username string
		Password string
	}
	Facebook struct {
		Username string
		Password string
	}
}

type RedisConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	OptionInBase string
}

func TestInitialization(t *testing.T) {
	Convey("Initializing the library", t, func() {
		os.Setenv("GOENV", "stage")
		Init("../data", "GOENV")

		Convey("When setting GOENV", func() {
			Convey("to stage", func() {
				Convey("env should be stage", func() {
					So(env, ShouldEqual, "stage")
				})
			})
		})
	})
}

func TestGetDomainBaseConfig(t *testing.T) {
	var basicConfig BasicConfig
	err := Get("basic", &basicConfig)
	Convey("When reading options from base.json", t, func() {
		Convey("No error should happen", func() {
			So(err, ShouldEqual, nil)
		})

		Convey("When getting first levele options", func() {
			Convey("The properties should have the right values", func() {
				So(basicConfig.Host, ShouldEqual, "localhost")
				So(basicConfig.Port, ShouldEqual, 3030)
				So(basicConfig.ApiKey, ShouldEqual, "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8")
			})
		})

		Convey("When reading second level options", func() {
			Convey("The properties should have the right values", func() {
				So(basicConfig.Twitter.Username, ShouldEqual, "mytwittername")
				So(basicConfig.Twitter.Password, ShouldEqual, "mytwitterpassword")
				So(basicConfig.Facebook.Username, ShouldEqual, "myfacebookname")
				So(basicConfig.Facebook.Password, ShouldEqual, "myfacebookpassword")
			})
		})
	})
}

func TestGetDomainRedisConfig(t *testing.T) {
	Convey("Initializing with empty environment", t, func() {
		os.Setenv("GOENV", "")
		Init("../data", "GOENV")
		var redisConf RedisConfig
		err := Get("redis", &redisConf)

		Convey("When reading options from redis.json", func() {
			Convey("No errors should happen", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("The right properties should be loaded", func() {
				So(redisConf.Host, ShouldEqual, "redis.host.com")
				So(redisConf.Port, ShouldEqual, 6379)
				So(redisConf.Username, ShouldEqual, "test")
				So(redisConf.Password, ShouldEqual, "password")
			})
		})
	})

	Convey("Initializing with environment set to ci", t, func() {
		os.Setenv("GOENV", "ci")
		Init("../data", "GOENV")
		var redisConf RedisConfig
		err := Get("redis", &redisConf)

		Convey("When reading options from redis.ci.json", func() {
			Convey("No errors should happen", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("The right properties should be loaded", func() {
				So(redisConf.Host, ShouldEqual, "redis.ci.host.com")
				So(redisConf.Port, ShouldEqual, 6379)
				So(redisConf.Username, ShouldEqual, "ciuser")
				So(redisConf.Password, ShouldEqual, "cipassword")
			})
		})
	})

	Convey("Initializing with environment set to stage", t, func() {
		os.Setenv("GOENV", "stage")
		Init("../data", "GOENV")
		var redisConf RedisConfig
		err := Get("redis", &redisConf)

		Convey("When reading options from redis.stage.json", func() {
			Convey("No errors should happen", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("The right properties should be loaded", func() {
				So(redisConf.Host, ShouldEqual, "redis.stage.host.com")
				So(redisConf.Port, ShouldEqual, 8080)
				So(redisConf.Username, ShouldEqual, "stageuser")
				So(redisConf.Password, ShouldEqual, "stagepassword")
			})
		})
	})

	Convey("Initializing with environment set to prod", t, func() {
		os.Setenv("GOENV", "prod")
		Init("../data", "GOENV")
		var redisConf RedisConfig
		err := Get("redis", &redisConf)

		Convey("When reading options from redis.prod.json", func() {
			Convey("No errors should happen", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("The right properties should be loaded", func() {
				So(redisConf.Host, ShouldEqual, "redis.prod.host.com")
				So(redisConf.Port, ShouldEqual, 9990)
				So(redisConf.Username, ShouldEqual, "produser")
				So(redisConf.Password, ShouldEqual, "prodpassword")
			})
		})
	})
}
