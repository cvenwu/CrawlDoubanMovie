package db

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/10/1 10:01 上午
 * @Desc: 关于将电影插入mysql的操作
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-crawl-douban/config"
	"go-crawl-douban/model"
)

var conn *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.MysqlUsername, config.MysqlPwd, config.MysqlHost, config.MysqlPort, config.MysqlDatabase)
	conn, _ = sql.Open("mysql", dataSourceName)

	if conn.Ping() != nil {
		fmt.Println("---------------------------初始化mysql连接失败---------------------------")
		return
	}

	fmt.Println("---------------------------初始化mysql连接成功---------------------------")
}

//向电影表中添加电影数据，如果添加成功就返回true，否则返回false
func AddMovieData(movie model.Movie) bool {

	sqlstr := "insert into movie(name, time, country, category, rating, people, quote) values(?, ?, ?, ?, ?, ?, ?)"
	stmt, err := conn.Prepare(sqlstr)
	if err != nil {
		fmt.Println("---------------------------插入电影数据预编译失败，请稍后再试---------------------------", err)
		return false
	}
	ret, err := stmt.Exec(movie.Name, movie.Time, movie.Country, movie.Category, movie.Rating, movie.People, movie.Quote)
	fmt.Println(err)
	if num, err := ret.RowsAffected(); num > 0 && nil == err {
		fmt.Println("---------------------------插入成功---------------------------")
		return true
	}

	fmt.Println(err)

	return false
}
