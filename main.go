package main

import (
	"fmt"
	"log"

	"github.com/gammazero/workerpool"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conf, err := readConfig("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	if err := conf.validate(); err != nil {
		log.Fatalln(err)
	}

	switch conf.Action {
	case "pause", "unpause", "empty":
		conf.action()
	case "info", "check":
		conf.info()
	}
}

func (c *Config) action() {
	var body struct {
		Action string `json:"action"`
	}
	body.Action = c.Action

	wp := workerpool.New(c.Worker)
	for _, target := range c.Target {
		target := target
		wp.Submit(func() {
			path := fmt.Sprintf("/api/topics/%s", target)
			if err := c.nsqAction(path, body); err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s is %sd\n", target, body.Action)
		})
	}

	wp.StopWait()
}

func (c *Config) info() {
	wp := workerpool.New(c.Worker)
	for _, target := range c.Target {
		target := target
		wp.Submit(func() {
			path := fmt.Sprintf("/api/topics/%s", target)
			resp, err := c.nsqInfo(path)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s status pause is %t and depth is %d\n", target, resp.Paused, resp.Depth)
		})
	}

	wp.StopWait()
}
