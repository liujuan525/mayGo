# mayGo

## 项目说明

本项目为 Golang 练习项目

> 文中出现的代码路径均以项目路径为根路径

## 项目目录说明

```
./                                  # 项目根目录
├── README.md                       # 项目说明文档
├── conf/                           # 项目配置文件目录
│   ├── app.ini                     # 项目配置文件
│   └── app.ini.bak                 # 项目配置示例文件
├── data/                           # 项目数据目录 (SQL，图示)
├── go.mod                          # 项目依赖管理文件
├── go.sum                          # 项目依赖管理文件
├── pkg/                            # 项目库文件夹 这个目录中存放的就是项目中可以被外部应用使用的代码库 ，其他的项目可以直接通过 import 引入这里的代码
│   ├── e                           # 项目错误定义目录
│   ├── setting                     # 项目配置文件解析
│   ├── response                    # 接口相应文件夹   
│   └── utils                       # 项目中使用到的助手类封装函数 (类似于 laravel 框架中的 helps)
├── main.go                         # 项目入口文件
├── routers/                        # 路由文件夹
├── runtime/                        # 运行时日志等记录文件
│   └──                 
└──vendor/                          # 项目依赖目录  便于 jenkin 编译项目
```
## 开发流程

本项目使用的 Go 版本为 1.14  (因为在此版本下 go module 可以在生产环境中使用，并且 go defer 效率的提升很高， Go 1.* 版本都有很好的兼容性)

* 数据文件(SQL)在项目中的 data 目录同步

本地开发流程：

> 开发阶段在本地开发并调试，因此需要配置[本地的 Go 可执行环境](http://docscn.studygolang.com/doc/install)
>> Linux 可以使用命令 `make golang` 来安装本项目开发环境
>
> 本地 Go 代理配置, 推荐写到 **~/.bashrc** 文件下
>> * export GO111MODULE=on
>> * export GOPROXY=https://goproxy.io

```bash
# 1. 下载项目到本地
$ git clone https://github.com/liujuan525/mayGo.git -b jiahui && cd mayGo

# 2. 创建配置文件
$ cp conf/app.ini.bak conf/app.ini

# 3. 修改配置文件
$ vi conf/app.ini

```