hiowo-go 一个 go 搭建的轻量级mvc框架，方便 go 新手更好的对 go 的 mvc 框架进行理解，我也会使用整套代码应用到我的个人博客上（ hiowo.com ）。



使用说明：

( tip：请先确保有 go 环境 )

1. 本地新建一个名为 blog 的 mysql 数据库

2. 执行根目录的 create.sql 文件里的SQL语句

3. 将 config 目录下的 config.yaml.hi 复制另存为 config.yaml，并修改里面的数据库配置


``` shell
copy config/config.yaml.hi config/config.yaml

vim config/config.yaml
```

4. 在根目录下执行 main.go 文件

``` shell
 go run ./main.go
```
5. 默认访问地址：http://localhost:8080/