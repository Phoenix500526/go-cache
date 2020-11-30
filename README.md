# go-cache
## 项目描述
go-cache 是一个学习型项目，借鉴了 groupcache 的设计思想，实现了一个简易的分布式缓存系统。

## 功能实现
在功能实现上， go-cache 实现了以下几个功能：
* LRU 淘汰算法
* 一致性哈希算法
* 使用 singleflight 来避免缓存击穿
* 基于 http 协议实现了分布式节点
* 使用 protobuf 来提高节点之间的通信效率
* 使用 go test 对部分代码实现了单元测试

## 目录结构
```bash
$ tree
.
├── gocache
│   ├── byteview.go
│   ├── cache.go
│   ├── consistenthash
│   │   ├── consistenthash.go
│   │   └── consistenthash_test.go
│   ├── gocache.go
│   ├── gocachepb
│   │   ├── gocachepb.pb.go
│   │   └── gocachepb.proto
│   ├── gocache_test.go
│   ├── go.mod
│   ├── go.sum
│   ├── http.go
│   ├── lru
│   │   ├── lru.go
│   │   └── lru_test.go
│   ├── peers.go
│   └── singleflight
│       ├── singleflight.go
│       └── singleflight_test.go
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
└── run.sh

5 directories, 22 files
```
其中，gocache 目录为项目的主体实现，其中的几个目录分别对应了：
> * lru : lru 缓存淘汰算法
> * consistenthash: 一致性哈希算法
> * gocachepb: 使用 Google Protobuffer 来编码报文，提高效率
> * singleflight: 防止缓存击穿

在 gocache 中以 gocache.go 文件为核心，
main.go 展示了 gocache 的一些基本用法，起到了功能测试的作用。