# Proxy Test

A delays tester for [Mihomo](https://github.com/MetaCubeX/mihomo).

## Installation

```bash
git clone https://github.com/MarkIvory2973/proxy-test.git
```

## Usage

```bash
cd proxy-test/src
go run . help
go run . --g SELECT -k 0.7
go run . --group SELECT --weight 0.3
```

## Parameters

|Name|Short|Required|Default|Description|
|-:|:-|:-:|:-|:-|
|--tls|-t|✗|no|Use TLS|
|--addr|-a|✗|127.0.0.1|Mihomo API Address|
|--port|-p|✗|9090|Mihomo API Port|
|--secret|-s|✗|-|Mihomo API Access Key|
|--group|-g|✓|-|Group Name|
|--url|-u|✗|(Google 204)|Test URL|
|--timeout|-m|✗|3000|Maximum Delay|
|--num|-n|✗|10|Number of Times|
|--interval|-i|✗|3|Interval of Each Test|
|--weight|-k|✗|0.5|Weight ∈ [0, 1]|

## The range of *k*

![The range of k](https://raw.githubusercontent.com/MarkIvory2973/proxy-test/main/imgs/k.png)
