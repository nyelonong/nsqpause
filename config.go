package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		NSQAdmin string        `yaml:"nsqadmin"`
		Worker   int           `yaml:"worker"`
		Timeout  time.Duration `yaml:"timeout"`
		Action   string        `yaml:"action"`
		Target   []string      `yaml:"target"`
	}
)

const (
	defaultTimeout = 2 * time.Second
)

func readConfig(file string) (Config, error) {
	var conf Config
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return conf, err
	}

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}

func (c *Config) validate() error {
	if c.NSQAdmin == "" {
		return errors.New("nsqadmin can not be empty.")
	}

	if c.Action == "" {
		return errors.New("action can not be empty.")
	}

	if c.Timeout == 0 {
		c.Timeout = defaultTimeout
	}

	if len(c.Target) == 0 {
		return errors.New("topic / channel can not be empty.")
	}

	if len(c.Target) < c.Worker {
		return errors.New("num of worker cannot exceed the number of target")
	}

	// Unlimited workerpool
	if c.Worker <= 0 {
		c.Worker = len(c.Target)
	}

	_, err := c.nsqInfo(fmt.Sprintf("/api/topics/%s", c.Target[0]))
	if err != nil {
		return err
	}

	fmt.Println("nsqpause is ready.")
	return nil
}
