package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type LoadTest struct {
	URL         string
	Requests    int
	Concurrency int
	Method      string
	Headers     map[string]string
}

func main() {
	baseURL := "http://localhost:8000"

	tests := []LoadTest{
		{
			URL:         baseURL + "/api/health",
			Requests:    1000,
			Concurrency: 50,
			Method:      "GET",
		},
		{
			URL:         baseURL + "/api/products",
			Requests:    500,
			Concurrency: 25,
			Method:      "GET",
		},
		{
			URL:         baseURL + "/api/search",
			Requests:    200,
			Concurrency: 10,
			Method:      "POST",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}

	fmt.Println("Starting load tests for Price Comparison API")
	fmt.Println("============================================")

	for _, test := range tests {
		runLoadTest(test)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("\nAll load tests completed!")
}

func runLoadTest(test LoadTest) {
	fmt.Printf("\nTesting: %s\n", test.URL)
	fmt.Printf("Method: %s, Requests: %d, Concurrency: %d\n",
		test.Method, test.Requests, test.Concurrency)
	fmt.Println("----------------------------------------")

	args := []string{
		"-n", fmt.Sprintf("%d", test.Requests),
		"-c", fmt.Sprintf("%d", test.Concurrency),
		"-m", test.Method,
	}

	for key, value := range test.Headers {
		args = append(args, "-H", fmt.Sprintf("%s: %s", key, value))
	}

	if test.Method == "POST" {
		args = append(args, "-d", `{"keyword":"kotobukiya"}`)
	}

	args = append(args, test.URL)

	cmd := exec.Command("hey", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Error running load test: %v\n", err)
		log.Printf("Output: %s\n", string(output))
		return
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Requests/sec") ||
		   strings.Contains(line, "Average") ||
		   strings.Contains(line, "Fastest") ||
		   strings.Contains(line, "Slowest") ||
		   strings.Contains(line, "Status code") {
			fmt.Println(line)
		}
	}
}

func checkHeyInstalled() bool {
	cmd := exec.Command("which", "hey")
	err := cmd.Run()
	return err == nil
}

func init() {
	if !checkHeyInstalled() {
		fmt.Println("hey is not installed. Installing...")
		cmd := exec.Command("go", "install", "github.com/rakyll/hey@latest")
		if err := cmd.Run(); err != nil {
			log.Fatal("Failed to install hey:", err)
		}
	}
}