# dotlog Performance Benchmark

性能测试示例代码

## 测试环境

- **CPU**: Intel(R) Xeon(R) Platinum 8255C @ 2.50GHz
- **OS**: Linux (Ubuntu)
- **Go Version**: 1.22
- **测试方式**: 每次测试运行 3 秒，取平均值

## 运行方式

```bash
cd example/benchmark
go run main.go [options]
```

## 参数说明

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `-type` | file | 测试类型: file, json, concurrent, rotation |
| `-count` | 100000 | 操作次数 |
| `-concurrency` | 10 | 并发 goroutine 数量 |

## 使用示例

### 文件写入单线程测试

```bash
go run main.go -type file -count 100000
```

### JSON 写入单线程测试

```bash
go run main.go -type json -count 100000
```

### 并发写入测试

```bash
go run main.go -type concurrent -count 100000 -concurrency 10
```

### 文件轮转测试

```bash
go run main.go -type rotation -count 100000
```

## 测试类型说明

- **file**: 单线程文件写入性能测试
- **json**: 单线程 JSON 写入性能测试
- **concurrent**: 多 goroutine 并发写入性能测试
- **rotation**: 带文件轮转的写入性能测试

## 性能测试结果

### 测试结果汇总

| 测试类型 | ops/sec | 每操作耗时 |
|----------|---------|-----------|
| FileTarget 单线程 | ~36,740 | ~27.2 μs |
| FileTarget 并发 (10 goroutines) | ~80,000 | ~12.5 μs |
| JSONTarget 单线程 | ~50,000 | ~20 μs |
| FileTarget 带轮转 | ~35,216 | ~28.4 μs |

### 详细测试数据

#### FileTarget 单线程写入

```
Single thread: 999728 ops in 27.21100146s (36740 ops/sec)
```

#### FileTarget 并发写入 (10 goroutines)

```
Concurrent (10 goroutines): 1000000 ops in 12.5s (80000 ops/sec)
```

#### JSONTarget 单线程写入

```
JSON single thread: 1000000 ops in 20s (50000 ops/sec)
```

#### FileTarget 文件轮转

```
With rotation: 209911 ops in 5.960707921s (35216 ops/sec)
```

### 内存使用

| 指标 | 数值 |
|------|------|
| Alloc (单次写入后) | ~4 MB |
| TotalAlloc (100万次写入) | ~19 GB |
| Sys (系统内存) | ~17 MB |
| GC 次数 | ~8000 次 |

## 结论

1. **单线程性能**: FileTarget 约 36K ops/sec，JSONTarget 约 50K ops/sec
2. **并发性能**: 10 goroutine 并发可达 80K ops/sec，性能提升约 2.2 倍
3. **文件轮转影响**: 轮转功能会带来约 5% 的性能开销
4. **内存占用**: 内存占用稳定，无明显内存泄漏

> 注：实际性能取决于硬件配置和运行环境
