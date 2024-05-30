# GVBServer 
博客项目后端项目


## 项目介绍

一些和其他博客不一样的点

1. 文章数据直接存储到es数据库，减少同步问题，使用redis缓存点赞数据，浏览量，所以文章随便点赞
2. 文章全文搜索，可精确定位到某一个自然标题处
3. 系统配置项众多，满足定制化需求
4. 日志系统，登录，操作，运行日志，可帮助管理员更好的观察系统运行情况
5. 群聊功能及其群聊配置
6. 八代具有完备的原型设计

## 项目截图

![](https://image.fengfengzhidao.com/pic/20231112144336.png)

![](https://image.fengfengzhidao.com/pic/20231112144425.png)

![](https://image.fengfengzhidao.com/pic/20231112144451.png)

![](https://image.fengfengzhidao.com/pic/20231112144510.png)

![](https://image.fengfengzhidao.com/pic/20231112144559.png)

## 部署


这里在源码里面提供了release的包

可直接供部署使用



先将前端和后端的release指定版本的包下载下来

分别放在不同的位置

例如

我的目录结构

```Bash
opt
  gvb
    server
      init
        ...
      main
    web
      dist
    web_mobile
    nginx.conf

```



## 基础服务部署

mysql

redis

es

我们直接使用release中提供的docker-compose直接一键运行

```Bash
cd init
sh init.sh
```





## 后端部署

导入数据

```Bash
./main -esload init/data/article_index_20231112.json
./main -esload init/data/full_text_index_20231112.json

./main -load init/data/gvb_db_20231112.sql
```



运行

```Bash
./main
```



### supervisor配置

上面的那个命令最多也就是看看效果，真要部署还是得用supervisor来管理我们的进程



```Bash
sh start.sh
```


## 前端部署



修改一下nginx的主配置文件

把gvb的配置文件引入即可

```nginx
http{
  include /opt/gvb/nginx.conf;
}
```


然后重启nginx即可

```Bash
nginx -s reload
```
