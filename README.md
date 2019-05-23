# go_runner

[![Build Status](https://www.travis-ci.org/26huitailang/go-runner.svg?branch=master)](https://www.travis-ci.org/26huitailang/go-runner)

run multiple concurrency tasks

测试运行

travis-ci没有打开没有反应

## Benchmark

    go test -bench . -cpuprofile profile
    go tool pprof profile
