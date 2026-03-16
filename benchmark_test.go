package dotlog

import (
	"path/filepath"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/targets"
)

// SetupBenchmark initializes the minimal required state for benchmarks
func SetupBenchmark() {
	// Initialize global config
	internal.InitInnerLogger("./", "utf-8")
	config.GlobalAppConfig = &config.AppConfig{
		Global: &config.GlobalConfig{
			ChanSize: 10000,
		},
	}
}

// BenchmarkFileTarget_SingleThreadWrite 单线程文件写入性能测试
func BenchmarkFileTarget_SingleThreadWrite(b *testing.B) {
	SetupBenchmark()

	tmpDir := b.TempDir()
	logFile := filepath.Join(tmpDir, "bench.log")

	// Create file target directly
	conf := &config.FileTargetConfig{
		Name:        "BenchLogger",
		IsLog:       true,
		Layout:      "{datetime} - {message}",
		Encode:      "utf-8",
		FileName:    logFile,
		FileMaxSize: 0, // No rotation
	}

	target := targets.NewFileTarget(conf)

	// Warmup
	for i := 0; i < 100; i++ {
		target.WriteLog("warmup message", "", "INFO")
	}

	// Benchmark
	b.ResetTimer()
	start := time.Now()

	for i := 0; i < b.N; i++ {
		target.WriteLog("benchmark message", "", "INFO")
	}

	elapsed := time.Since(start)
	opsPerSec := float64(b.N) / elapsed.Seconds()

	b.ReportMetric(opsPerSec, "ops/sec")
	b.Logf("Single thread: %d ops in %v (%.2f ops/sec)", b.N, elapsed, opsPerSec)
}

// BenchmarkFileTarget_ConcurrentWrite 并发文件写入性能测试
func BenchmarkFileTarget_ConcurrentWrite(b *testing.B) {
	SetupBenchmark()

	tmpDir := b.TempDir()
	logFile := filepath.Join(tmpDir, "bench.log")

	// Create file target directly
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
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			target.WriteLog("warmup", "", "INFO")
		}()
	}
	wg.Wait()

	// Benchmark
	concurrency := 10
	opsPerGoroutine := b.N / concurrency

	b.ResetTimer()
	start := time.Now()

	for i := 0; i < concurrency; i++ {
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
	totalOps := concurrency * opsPerGoroutine
	opsPerSec := float64(totalOps) / elapsed.Seconds()

	b.ReportMetric(opsPerSec, "ops/sec")
	b.Logf("Concurrent (%d goroutines): %d ops in %v (%.2f ops/sec)", concurrency, totalOps, elapsed, opsPerSec)
}

// BenchmarkJSONTarget_SingleThreadWrite 单线程 JSON 写入性能测试
func BenchmarkJSONTarget_SingleThreadWrite(b *testing.B) {
	SetupBenchmark()

	tmpDir := b.TempDir()
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
	for i := 0; i < 100; i++ {
		target.WriteLog("warmup message", "", "INFO")
	}

	// Benchmark
	b.ResetTimer()
	start := time.Now()

	for i := 0; i < b.N; i++ {
		target.WriteLog("benchmark message", "", "INFO")
	}

	elapsed := time.Since(start)
	opsPerSec := float64(b.N) / elapsed.Seconds()

	b.ReportMetric(opsPerSec, "ops/sec")
	b.Logf("JSON single thread: %d ops in %v (%.2f ops/sec)", b.N, elapsed, opsPerSec)
}

// BenchmarkJSONTarget_ConcurrentWrite 并发 JSON 写入性能测试
func BenchmarkJSONTarget_ConcurrentWrite(b *testing.B) {
	SetupBenchmark()

	tmpDir := b.TempDir()
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
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			target.WriteLog("warmup", "", "INFO")
		}()
	}
	wg.Wait()

	// Benchmark
	concurrency := 10
	opsPerGoroutine := b.N / concurrency

	b.ResetTimer()
	start := time.Now()

	for i := 0; i < concurrency; i++ {
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
	totalOps := concurrency * opsPerGoroutine
	opsPerSec := float64(totalOps) / elapsed.Seconds()

	b.ReportMetric(opsPerSec, "ops/sec")
	b.Logf("JSON Concurrent (%d goroutines): %d ops in %v (%.2f ops/sec)", concurrency, totalOps, elapsed, opsPerSec)
}

// BenchmarkFileTarget_MemoryUsage 内存使用测试
func BenchmarkFileTarget_MemoryUsage(b *testing.B) {
	SetupBenchmark()

	tmpDir := b.TempDir()
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

	// Force GC before measurement
	b.StopTimer()
	runtime.GC()
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)
	b.StartTimer()

	// Write messages
	for i := 0; i < b.N; i++ {
		target.WriteLog("memory test message", "", "INFO")
	}

	// Read GC stats
	runtime.GC()
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)

	b.Logf("Memory stats after %d writes:", b.N)
	b.Logf("  Alloc: %v KB", memStatsAfter.Alloc/1024)
	b.Logf("  TotalAlloc: %v KB", memStatsAfter.TotalAlloc/1024)
	b.Logf("  Sys: %v KB", memStatsAfter.Sys/1024)
	b.Logf("  NumGC: %d", memStatsAfter.NumGC)
}

// BenchmarkFileTarget_FileRotation 文件轮转性能测试
func BenchmarkFileTarget_FileRotation(b *testing.B) {
	SetupBenchmark()

	tmpDir := b.TempDir()
	logFile := filepath.Join(tmpDir, "bench.log")

	// Create file target with small max size to trigger rotation
	conf := &config.FileTargetConfig{
		Name:        "BenchLogger",
		IsLog:       true,
		Layout:      "{datetime} - {message}",
		Encode:      "utf-8",
		FileName:    logFile,
		FileMaxSize: 10, // 10KB to trigger rotation quickly
		MaxBackups:  3,
	}

	target := targets.NewFileTarget(conf)

	// Warmup - write until first rotation
	for i := 0; i < 2000; i++ {
		target.WriteLog("warmup message", "", "INFO")
	}

	// Benchmark - measure rotation overhead
	b.ResetTimer()
	start := time.Now()

	for i := 0; i < b.N; i++ {
		target.WriteLog("benchmark message", "", "INFO")
	}

	elapsed := time.Since(start)
	opsPerSec := float64(b.N) / elapsed.Seconds()

	b.ReportMetric(opsPerSec, "ops/sec")
	b.Logf("With rotation: %d ops in %v (%.2f ops/sec)", b.N, elapsed, opsPerSec)
}
