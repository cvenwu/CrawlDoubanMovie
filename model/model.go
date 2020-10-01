package model

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/10/1 9:42 上午
 * @Desc:
 */

type Movie struct {
	Name     string //中文名
	Time     string //时间
	Country  string //国家
	Category string //分类
	Rating   string //评分
	People   int    //有多少人评价
	Quote    string //对应的一句话
}
