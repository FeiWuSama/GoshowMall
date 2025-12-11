# GoShowMall - 基于Golang开发的在线商城项目

<p align="center">
  <a href="https://github.com/FeiWuSama/GoshowMall/stargazers">
    <img src="https://img.shields.io/github/stars/FeiWuSama/GoshowMall" alt="GitHub Stars">
  </a>
  <a href="https://github.com/FeiWuSama/GoshowMall/issues">
    <img src="https://img.shields.io/github/issues/FeiWuSama/GoshowMall" alt="GitHub Issues">
  </a>
</p>

>GoShowMall: 一个适合新手学习Go语言的在线商城项目(正在长期开发中)
> 
> 后端技术栈：Go + Gorm + MySQL + Redis + RabbitMQ + ElasticSearch + Docker + Docker Compose
> 
> 前端技术栈：Vue + AntDesign Vue + Pinia + Iframe + Canvas
> 
> 如果你觉得这个项目对你有帮助，请给个Star，谢谢！
>

## 项目介绍：
自己也作为一位新手学习Go语言，想自己写一个项目，有感而发，觉得自己写一个商城项目会比较适合大部分新手入门Golang网站开发，所以就写了一个网上比较通用的项目

> 项目目前正在开发中，功能正在持续更新中，欢迎大家提issue，欢迎大家提PR，我希望能同为Golang初学者的小伙伴一起学习，一起进步

## 项目亮点

> 注：前端大部分的代码都是使用 AI Vibe Coding 生成，界面丑陋以及出现bug请谅解

1）在后端搭建的业务框架中，将路由层抽象脱离出业务层，使路由层能够轻松兼容各种Golang开发框架，避免开发框架与业务逻辑过度耦合，也更适应于Golang的开发思想

2）后端使用了Gorm生成器，拥有实体类对应query结构体，避免了数据库表修改后导致的对应业务层的批量更改

3）后端的登录业务接入了滑块验证码，有更灵敏的接口校验功能

4）后端实现了多种登录方式，包括手机密码登录、手机短信登录和多种app应用登录，大幅提高了登录业务的可拓展性