package main

import (
	"fmt"
	"go-crawl-douban/db"
	"go-crawl-douban/model"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/9/29 8:54 下午
 * @Desc: 爬取豆瓣电影Top250数据
 */

/*
页面规则：
	https://movie.douban.com/top250?start=0&filter=
start表示从多少开始显示当前页面，

爬取一下电影信息：
电影名
导演
时间
国家
分类
评分
多少人评价
每个电影下面的最经典的一句话
*/

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"
	//每个电影对应的信息解析器
	contentRegex = `<div class="hd">[^<]*<a[^>]*>[^<]*<span[^>]*>([^<]*)</span>[^<]*(<span[^>]*>[^<]*</span>)?[^<]*<span[^>]*>[^<]*</span>[^<]*</a>[^<]*(<span[^>]*>[^<]*</span>)?[^<]*</div>[^<]*<div[^>]*>[^<]*<p[^>]*>([^<]*)<br>([^<]*)</p>[^<]*<div[^>]*>[^<]*<span[^>]*></span>[^<]*<span[^>]*>([^<]*)</span>[^<]*<span[^>]*></span>[^<]*<span>([^人]*)人评价</span>[^<]*</div>[^<]*(<p[^>]*>[^<]*<span[^>]*>([^<]*)</span>[^<]*</p>)?[^<]*</div>`
	//提取导演名字
	directorRegex = `导演: ([^a-zA-Z]*)`
)

//获取第几页对应的页面URL
func UrlConfig(page int) string {
	return fmt.Sprintf("https://movie.douban.com/top250?start=%d&filter=", (page-1)*25)
}

func contentParser(moviesList *[]model.Movie) {
	client := http.Client{}
	regMovieContent := regexp.MustCompile(contentRegex)

	for i := 1; i <= 10; i++ {

		fmt.Println("请求的url为：", UrlConfig(i))
		//模拟一个新的http请求
		req, _ := http.NewRequest("GET", UrlConfig(i), nil)
		//给请求添加伪装的请求头部
		req.Header.Add("User-Agent", UserAgent)
		//发起请求
		resp, _ := client.Do(req)
		//打印一下爬取页面的状态码
		fmt.Printf("--------------------------------请求第%d页，状态码：%d--------------------------------\n", i, resp.StatusCode)
		//读取页面爬取的数据
		body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))

		//使用正则匹配进行匹配
		retMovie := regMovieContent.FindAllStringSubmatch(string(body), -1)
		for _, v := range retMovie {
			//v[0]为匹配到的全部信息
			//v[1]为中文名字，v[2]为电影的别名或英文名字,v[3]为电影是否可以播放
			//v[4]为导演主演信息
			//v[5]为时间+国家+分类
			//v[6]为评分
			//v[7]为多少人评价
			//v[8]为分组捕获的一句话的p标签
			//v[9]为对应的一句话
			//director, star := getDirectorAndStar(v[4])
			time, country, category := getMovieOthers(v[5])
			//给电影结构体赋值
			tempMovie := model.Movie{}
			tempMovie.Name = v[1]
			tempMovie.Category = category
			tempMovie.Country = country
			tempMovie.Quote = v[9]
			tempMovie.Rating = v[6]
			tempMovie.Time = time
			tempMovie.People, _ = strconv.Atoi(v[7])

			*moviesList = append(*moviesList, tempMovie)
		}

		//fmt.Printf("--------------------------------第%d页抓取到：%d条--------------------------------\n", i, len(ret))
	}

	fmt.Printf("--------------------------------爬取成功，共爬取到信息：%d条--------------------------------\n", len(*moviesList))
}

//获取电影的导演以及主演名字
//TODO：提取电影的导演以及主演名字
func getDirectorAndStar(content string) (director string, star string) {
	//reg := regexp.MustCompile(directorRegex)
	//fmt.Println(content)
	//ret := reg.FindAllStringSubmatch(strings.Trim(content, " "), -1)
	//
	//ret2 := strings.Split(strings.Trim(ret[0][0], " "), " ")
	//fmt.Println(len(ret))
	//fmt.Println(ret2)
	return
}

//得到电影的时间，国家以及分类
func getMovieOthers(content string) (time string, country string, category string) {
	ret := strings.Split(strings.Trim(content, " "), "&nbsp;/&nbsp;")
	time, country, category = ret[0], ret[1], ret[2]
	time = time[len(time)-4:]
	return
}

func main() {
	moviesList := []model.Movie{}
	contentParser(&moviesList)

	for index, movie := range moviesList {
		opType := db.AddMovieData(movie)
		if opType {
			fmt.Printf("插入第%d条数据成功\n", index+1)
		} else {
			fmt.Printf("插入第%d条数据失败\n", index+1)
		}
	}
}
