# 问答社区

Hi,there!👋

## 🍔简介

这是一个用**Golang**实现的**问答社区!!!**
可以在**csa.madeindz.work**网站(暂无前端页面)用ApiPost或者PostMan等**调试工具**进行调试
## ✨功能

### 认证系统

- JWT...登录时颁发`token`,时效2小时
- 众多操作需要先进行对`token`的解析以获取用户信息和进行安全验证

### 用户系统

- 用户的登录与注册
- 对问题进行点赞,或者取消点赞
- 获取问题的点赞数(游客可查看)
- 获取自己的问题和回答

### 问答系统

- 对问题的增删改查
- 对回答的增删改查
- 回答可以被评论



## 🃏技术栈

![img](https://github.com/StellarisW/gohu/raw/master/manifest/image/mysql.svg)

- [mysql](https://www.mysql.com/)

> 一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统关系型数据库管理系统之一，在 WEB 应用方面，MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一

[![img](https://github.com/StellarisW/gohu/raw/master/manifest/image/redis.svg)](https://github.com/StellarisW/gohu/blob/master/manifest/image/redis.svg)

- [redis](https://redis.io/)

> 一个开源的、使用C语言编写的、支持网络交互的、可基于内存也可持久化的Key-Value数据库

![grpc 的图像结果](https://th.bing.com/th/id/OIP.pTzFSebJ00beKGCeu0u76AHaEW?w=294&h=180&c=7&r=0&o=5&dpr=1.5&pid=1.7)



- [grpc](https://grpc.io/)

> gRPC是一个现代的开源高性能远程过程调用（RPC）框架，可以在任何环境中运行。

还有**Viper,Logger**等小工具...
## 🍟调试参数
JWT使用请求头Authorization ===> Bearer `token` 
- **注册/登录**
| 参数名   | 参数值 |
| -------- | ------ |
| username | xxx    |
| password | xxx    |

- **问题创建**

| 参数名  | 参数值 |
| ------- | ------ |
| message | xxx    |

- **回答创建**

| 参数名      | 参数值 |
| ----------- | ------ |
| message     | xxx    |
| question_id | xxx    |

- **修改问题/回答**
- **评论**

| 参数名  | 参数值 |
| ------- | ------ |
| id      | xxx    |
| message | xxx    |

- **删除问题/回答**
- **点赞/取消点赞**

| 参数名 | 参数值 |
| ------ | ------ |
| id     | xxx    |



## ✅CSA完成度

- #### 基础内容（必做）

  - 用户登录/注册✅成功用gRPC微服务实现
  - 用户可以发起一个问题✅成功用gRPC微服务实现
  - 所有用户都可以回答问题✅成功用gRPC微服务实现
  - 用户可以获取自己的所有问题、所有回答（回答需要能定位出处，比如是哪个问题的回答）✅
  - 用户可以删除或修改自己发布的问题或回答✅
  - 一个较为完善的说明文档，文档内容存放在项目 README 中✅我正在做的事情

- #### 加分内容（选做）

  - 用户可以对**问题的回答**评论（类似于贴吧楼中楼）✅
  - 使用 Redis 作为缓存✅
  - 实现点赞功能（推荐使用 Redis 进行实现）✅
  - 实现 RPC 或微服务（推荐 gRPC、go-zero、Kitex）✅gRPC
  - 将服务部署上线，能直接通过接口调试工具访问到✅csa.madeindz.work

