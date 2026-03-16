# dotlog Performance Benchmark

性能测试示例代码

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

## 性能测试结果参考

| 测试类型 | 性能 |
|----------|------|
| FileTarget 单线程 | ~36,000 ops/sec |
| FileTarget 并发 (10 goroutines) | ~80,000 ops/sec |
| JSONTarget 单线程 | ~50,000 ops/sec |
| FileTarget 带轮转 | ~35,000 ops/sec |

> 注：实际性能取决于硬件配置和运行环境
