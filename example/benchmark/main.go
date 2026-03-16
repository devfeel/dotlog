package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/targets"
)

var (
	benchType   = flag.String("type", "file", "Benchmark type: file, json, concurrent, rotation")
	benchCount  = flag.Int("count", 100000, "Number of operations")
	concurrency = flag.Int("concurrency", 10, "Number of goroutines for concurrent test")
)

func init() {
	// Initialize global config
	internal.InitInnerLogger("./", "utf-8")
	config.GlobalAppConfig = &config.AppConfig{
		Global: &config.GlobalConfig{
			ChanSize: 100000,
		},
	}
}

func main() {
	flag.Parse()

	fmt.Println("=== dotlog Performance Benchmark ===")
	fmt.Printf("Type: %s, Count: %d\n", *benchType, *benchCount)
	fmt.Println()

	switch *benchType {
	case "file":
		runFileTargetBenchmark()
	case "json":
		runJSONTargetBenchmark()
	case "concurrent":
		runConcurrentBenchmark()
	case "rotation":
		runRotationBenchmark()
	default:
		fmt.Println("Unknown benchmark type")
	}
}

func runFileTargetBenchmark() {
	tmpDir := "/tmp/dotlog-bench"
	logFile := filepath.Join(tmpDir, "bench.log")

	conf := &config.FileTargetConfig{
		Name:        "BenchLogger",
		IsLog:       true,
		Layout:      "{datetime} - {message}",
		Encode:      "utf-8",
		FileName:    logFile,
		FileMaxSize: 0,
	}

	target := targets.NewFileTarget(conf)

	// Warmup
	for i := 0; i < 1000; i++ {
		target.WriteLog("warmup message", "", "INFO")
	}

	// Benchmark
	start := time.Now()
	for i := 0; i < *benchCount; i++ {
		target.WriteLog("benchmark message", "", "INFO")
	}
	elapsed := time.Since(start)

	opsPerSec := float64(*benchCount) / elapsed.Seconds()
	fmt.Printf("FileTarget Single Thread:\n")
	fmt.Printf("  Total: %d ops in %v\n", *benchCount, elapsed)
	fmt.Printf("  Performance: %.2f ops/sec\n", opsPerSec)
}

func runJSONTargetBenchmark() {
	tmpDir := "/tmp/dotlog-bench"
	logFile := filepath.Join(tmpDir, "bench.json")

	prettyPrint := false
	conf := &config.JSONTargetConfig{
		Name:        "JSONLogger",
		IsLog:       true,
		Layout:      "{datetime} - {message}",
		Encode:      "utf-8",
		FileName:    logFile,
		FileMaxSize: 0,
		PrettyPrint: &prettyPrint,
	}

	target := targets.NewJSONTarget(conf)

	// Warmup
	for i := 0; i < 1000; i++ {
		target.WriteLog("warmup message", "", "INFO")
	}

	// Benchmark
	start := time.Now()
	for i := 0; i < *benchCount; i++ {
		target.WriteLog("benchmark message", "", "INFO")
	}
	elapsed := time.Since(start)

	opsPerSec := float64(*benchCount) / elapsed.Seconds()
	fmt.Printf("JSONTarget Single Thread:\n")
	fmt.Printf("  Total: %d ops in %v\n", *benchCount, elapsed)
	fmt.Printf("  Performance: %.2f ops/sec\n", opsPerSec)
}

func runConcurrentBenchmark() {
	tmpDir := "/tmp/dotlog-bench"
	logFile := filepath.Join(tmpDir, "bench.log")

	conf := &config.FileTargetConfig{
		Name:        "BenchLogger",
		IsLog:       true,
		Layout:      "{datetime} - {message}",
		Encode:      "utf-8",
		FileName:    logFile,
		FileMaxSize: 0,
	}

	target := targets.NewFileTarget(conf)

	// Warmup
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			target.WriteLog("warmup", "", "INFO")
		}()
	}
	wg.Wait()

	// Benchmark
	opsPerGoroutine := *benchCount / *concurrency
	start := time.Now()

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				target.WriteLog("benchmark message", "", "INFO")
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	totalOps := *concurrency * opsPerGoroutine

	opsPerSec := float64(totalOps) / elapsed.Seconds()
	fmt.Printf("FileTarget Concurrent (%d goroutines):\n", *concurrency)
	fmt.Printf("  Total: %d ops in %v\n", totalOps, elapsed)
	fmt.Printf("  Performance: %.2f ops/sec\n", opsPerSec)
}

func runRotationBenchmark() {
	tmpDir := "/tmp/dotlog-bench-rotation"
	logFile := filepath.Join(tmpDir, "bench.log")

	conf := &config.FileTargetConfig{
		Name:        "BenchLogger",
		IsLog:       true,
		Layout:      "{datetime} - {message}",
		Encode:      "utf-8",
		FileName:    logFile,
		FileMaxSize: 10, // 10KB to trigger rotation
		MaxBackups:  3,
	}

	target := targets.NewFileTarget(conf)

	// Warmup - write until first rotation
	for i := 0; i < 5000; i++ {
		target.WriteLog("warmup message", "", "INFO")
	}

	// Benchmark
	start := time.Now()
	for i := 0; i < *benchCount; i++ {
		target.WriteLog("benchmark message", "", "INFO")
	}
	elapsed := time.Since(start)

	opsPerSec := float64(*benchCount) / elapsed.Seconds()
	fmt.Printf("FileTarget with Rotation:\n")
	fmt.Printf("  Total: %d ops in %v\n", *benchCount, elapsed)
	fmt.Printf("  Performance: %.2f ops/sec\n", opsPerSec)
}

// Memory benchmark
func runMemoryBenchmark() {
	tmpDir := "/tmp/dotlog-bench"
	logFile := filepath.Join(tmpDir, "bench.log")

	conf := &config.FileTargetConfig{
		Name:        "BenchLogger",
		IsLog:       true,
		Layout:      "{datetime} - {message}",
		Encode:      "utf-8",
		FileName:    logFile,
		FileMaxSize: 0,
	}

	target := targets.NewFileTarget(conf)

	// Warmup
	for i := 0; i < 10000; i++ {
		target.WriteLog("warmup message", "", "INFO")
	}

	// Force GC before measurement
	runtime.GC()
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	// Write messages
	for i := 0; i < *benchCount; i++ {
		target.WriteLog("memory test message", "", "INFO")
	}

	// Read GC stats
	runtime.GC()
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)

	fmt.Printf("Memory stats after %d writes:\n", *benchCount)
	fmt.Printf("  Alloc: %v KB\n", memStatsAfter.Alloc/1024)
	fmt.Printf("  TotalAlloc: %v KB\n", memStatsAfter.TotalAlloc/1024)
	fmt.Printf("  Sys: %v KB\n", memStatsAfter.Sys/1024)
	fmt.Printf("  NumGC: %d\n", memStatsAfter.NumGC)
}
