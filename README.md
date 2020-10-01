# CrawlDoubanMovie

使用Golang中的正则表达式爬取豆瓣Top250电影并持久化到数据库中

## 说明

爬取豆瓣电影Top250数据：[访问](https://movie.douban.com/top250?start=0&filter=)

使用正则爬取豆瓣电影Top250获得信息后并将信息持久化到数据库中。


## 使用步骤
1. 克隆代码
2. 修改config文件夹下的数据库用户名密码等配置信息，同时确保自己已经建库建表（参照项目目录下的`movie.sql`）
3. 运行main.go即可


