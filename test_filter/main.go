/**
 * @author zhangyuehao
 * @date 2019-01-21 10:01
 */

package main

import "github.com/astaxie/beego"

func main() {

	beego.Run()
	beego.InsertFilter("/*", beego.BeforeRouter, UrlManager)
}
