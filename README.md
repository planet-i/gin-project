## 目录结构
```shell
D:.
├─api                  // 路由逻辑处理
│  └─v1
├─config               // 项目配置
├─docs                 // 接口文档
├─middleware           // 中间件
│  └─jwt
├─models               // 数据模型层
├─pkg                  // 工具包
│  ├─app
│  ├─cron
│  ├─e                 // 错误码处理
│  ├─export
│  ├─file
│  ├─gredis
│  ├─logging
│  ├─qrcode
│  ├─setting
│  ├─upload
│  └─util              // 公共函数库
├─routers              // 路由注册
├─runtime              // 应用运行时数据
│  ├─logs
│  ├─qrcode
│  └─upload
│      └─images
└─service
    ├─article_service
    ├─auth_service
    ├─cache_service
    └─tag_service
```