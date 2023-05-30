package main

import (
	"fmt"
	"log"
	"strings"

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

var (
	ResultMap = map[string][]string{}

	topicNotFoundStr = "topicnotfound"
	failStr          = "fail"
	pausedStr        = "paused"
	unpausedStr      = "unpaused"

	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

func (c *Config) action() {
	var body struct {
		Action string `json:"action"`
	}
	body.Action = c.Action
	successCount := 0

	wp := workerpool.New(c.Worker)
	for _, target := range c.Target {
		target := target
		wp.Submit(func() {
			path := fmt.Sprintf("/api/topics/%s", target)
			if err := c.nsqAction(path, body); err != nil {
				// log.Println(err)
				if strings.Contains(err.Error(), "TOPIC_NOT_FOUND") {
					ResultMap[topicNotFoundStr] = append(ResultMap[topicNotFoundStr], target)
				} else {
					ResultMap[failStr] = append(ResultMap[failStr], fmt.Sprintf("%s: %s", target, err.Error()))
				}
				return
			}
			// fmt.Printf("%s is %sd\n", target, body.Action)
			successCount++
		})
	}

	wp.StopWait()

	// Print Result
	fmt.Printf("\nTotal Topic\n%sd : %d\n", body.Action, successCount)
	fmt.Printf("Not found : %d\n", len(ResultMap[topicNotFoundStr]))
	fmt.Printf("Failed : %d\n", len(ResultMap[failStr]))
	printFailedResult()

}

func makeColorStr(color, str string) string {
	return fmt.Sprintf("%s%s%s", color, str, colorReset)
}

func (c *Config) info() {
	wp := workerpool.New(c.Worker)
	for _, target := range c.Target {
		target := target
		wp.Submit(func() {
			path := fmt.Sprintf("/api/topics/%s", target)
			resp, err := c.nsqInfo(path)
			if err != nil {
				// log.Println(err)
				if strings.Contains(err.Error(), "TOPIC_NOT_FOUND") {
					ResultMap[topicNotFoundStr] = append(ResultMap[topicNotFoundStr], target)
				} else {
					ResultMap[failStr] = append(ResultMap[failStr], fmt.Sprintf("%s: %s", target, err.Error()))
				}
				return
			}
			color := colorGreen
			if resp.Paused {
				color = colorRed
				ResultMap[pausedStr] = append(ResultMap[pausedStr], target)
			} else {
				ResultMap[unpausedStr] = append(ResultMap[unpausedStr], target)

			}
			fmt.Printf("%s status pause is %s and depth is %d\n", target, makeColorStr(color, fmt.Sprintf("%t", resp.Paused)), resp.Depth)
		})
	}

	wp.StopWait()

	// Print Result
	fmt.Printf("\nTotal Topic\n")
	fmt.Printf("Unpaused : %d\n", len(ResultMap[unpausedStr]))
	fmt.Printf("Paused : %d\n", len(ResultMap[pausedStr]))
	fmt.Printf("Not found : %d\n", len(ResultMap[topicNotFoundStr]))
	fmt.Printf("Failed : %d\n", len(ResultMap[failStr]))

	fmt.Printf("\n\n%s:\n", makeColorStr(colorRed, "List Topic paused"))
	for _, data := range ResultMap[pausedStr] {
		fmt.Println(data)
	}

	printFailedResult()

}

func printFailedResult() {
	fmt.Printf("\n\n%s:\n", makeColorStr(colorRed, "List Topic not found"))
	for _, data := range ResultMap[topicNotFoundStr] {
		fmt.Println(data)
	}

	fmt.Printf("\n\n%s:\n", makeColorStr(colorRed, "List Topic failed"))
	for _, data := range ResultMap[failStr] {
		fmt.Println(data)
	}
}
