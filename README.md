# gin-starter

## 项目介绍

基于这个模板，我们可以更快的构建程序

+ 基于gin框架
+ 简单的增删改查应用
+ 提供数库迁移
+ 提供模型验证和 i18n 错误代码转换
+ 为我们的应用领域提供业务错误代码
+ 提供中间件处理全局错误与响应 

## 启动项目

### local build and run

1.在计算机中安装go环境与postgres

2.

```
make local-build
make local-run
```

### build 

```bash
./dockerbuild.sh
```
### test

```bash
./dockerbuild.sh test
```
### codegen

```bash
./dockerbuild.sh codegen
```

### run

```bash
docker-compose up server
```
