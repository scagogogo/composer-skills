package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/scagogogo/composer-crawler/pkg/client"
)

func main() {
	// 定义命令行参数
	packageName := flag.String("package", "", "Package name to fetch (e.g. symfony/console)")
	stats := flag.Bool("stats", false, "Get repository statistics")
	advisories := flag.Bool("advisories", false, "Get security advisories")
	outputFile := flag.String("output", "", "Output file for results (default: stdout)")
	timeout := flag.Duration("timeout", 30*time.Second, "HTTP request timeout")
	flag.Parse()

	// 创建客户端
	composerClient := client.NewComposerClient(*timeout)

	var data interface{}
	var err error

	// 根据参数执行相应操作
	switch {
	case *packageName != "":
		fmt.Printf("Fetching package information for %s...\n", *packageName)
		data, err = composerClient.GetPackage(*packageName)
		if err != nil {
			log.Fatalf("Error fetching package information: %v", err)
		}

	case *stats:
		fmt.Println("Fetching repository statistics...")
		data, err = composerClient.GetStatistics()
		if err != nil {
			log.Fatalf("Error fetching statistics: %v", err)
		}

	case *advisories:
		fmt.Println("Fetching security advisories...")
		data, err = composerClient.GetSecurityAdvisories()
		if err != nil {
			log.Fatalf("Error fetching security advisories: %v", err)
		}

	default:
		fmt.Println("No action specified. Use -package, -stats, or -advisories.")
		flag.Usage()
		os.Exit(1)
	}

	// 将结果转为 JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling result to JSON: %v", err)
	}

	// 输出结果
	if *outputFile == "" {
		// 输出到标准输出
		fmt.Println(string(jsonData))
	} else {
		// 输出到文件
		err = os.WriteFile(*outputFile, jsonData, 0644)
		if err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
		fmt.Printf("Results written to %s\n", *outputFile)
	}
}
