# go-jy-sso-test
---

## Description
用户测试sso-api统一登录实现go代码demo示例

## Feature

- 提供携程Apollo配置，提供统一配置管理
- 提供Gin Url分组管理，提供统一路由插拔策略
- 提供Mysql，提供统一数据库操作Client
- 提供Redis，提供统一Redis操作，Set Get
- 提供test 模块，熟悉如何完成业务开发

## 使用
```
apollo模块为全局所有配置唯一来源
修改apollo模块中 struct中相关配置(包含：mysql相关配置，redis相关配置，以及未来可能使用到的所有配置)

设置环境变量(正常Prod环境以及Dev Test 已经有统一的环境变量)，仅适用本地开发环境：

export RUNTIME_ENV=dev && export RUNTIME_CLUSTER=default && export RUNTIME_APP_NAME=xxxx && export LOG_BASE=debug

go run cmd/app/main.go 
```

## CI build 参考如下
https://github.com/hashicorp/terraform/tree/main/scripts





