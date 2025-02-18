# LearnGolang
Some examples on Golang

# Golang exercises

# algorithm

# data structure

## Go项目目录规范
Go项目的基本布局，不是Go开发团队定义的官方标准，但是Go生态系统中一组常见的老项目和新项目的布局模式。



### Go目录

#### /cmd

本项目的主干。

每个应用程序的目录名应该与你想要的可执行文件的名称相匹配（`/cmd/myapp`）；

不要在这个目录下放太多代码，如果你认为代码可以导入并在其他项目中使用，那么它应该位于`/pkg`目录中；如果代码不是可重用，或者你不希望其他人重用它，请将代码放到`/internal`目录中

#### /internal

私有应用程序和库代码，这是你不希望其他人在其他应用程序或库中导入代码。实际应用程序代码可以放在`internal/app/myapp`目录下

#### /pkg

外部应用程序可以使用的库代码(`pkg/mypubliclib`)

####  /vendor

应用程序依赖项

#### 服务应用程序目录

#### /api

OpenAPI/Swagger规范，json文件、协议定义文件

### Web应用程序目录

#### /web

特定与Web应用程序的组件：静态Web资源、服务端模板和SPAs

#### 通用应用目录

#### /configs

配置文件模板或默认配置

#### /init

System init (systemd ,upstart,sysv)和process manager/supervisor(runit . supervisor)配置

#### /scripts

执行各种构建、安装、分析等操作的脚本

#### /build

打包和持续集成

1. 将云、容器（Docker）、操作系统(deb、rpm、pkg)的包配置和脚本放在`/build/package`目录下
2. 将CI(travis、circle、drone)配置和脚本放在`/build/ci`目录中。

#### /deployments

IaaS,PaaS、系统和容器编排部署配置和模板（docker-compose、kubernetes/helm、mesos、terraform、bosh），这个目录被称为`/deploy`

#### /test

外部测试应用程序和测试数据。



### 其他目录

#### /docs

设计和用户文档

#### /tools

这个项目支持的工具，这些工具可以从`/pkg`和`/internal`目录导入代码

#### /examples

程序实例

#### /third_party

外部辅助工具、分叉代码和带三方工具（swagger ui）

#### /githooks

git hooks

#### /assets

和存储库一起使用的其他资产

#### /website

如果你不使用Github页面，则这里放置项目的网站数据



### 不应该有的目录

#### /src

### Update