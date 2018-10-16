package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var templet string = "# SSR 账号每日更新 \n> 数据来源: [逗比根据地](https://doub.io/sszhfx/) \n----------------------------------------------\n## 更新日期：%s \n***食用方法：复制下面的节点到SSR客户端去重添加即可***\n\n %s"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("doub.io"),
	)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		comments := e.ChildText(".commentlist")
		reg := regexp.MustCompile(`ssr(.*?)[\n]`)
		ssrAddr := reg.FindAllString(comments, -1)
		ssrAddr2Str := strings.Replace(strings.Trim(fmt.Sprint(ssrAddr), "[]"), " ", "\n", -1)
		file, error := os.OpenFile("README.md", os.O_WRONLY|os.O_TRUNC, 0600)
		defer file.Close()
		if error != nil {
			fmt.Println(error)
		}
		t := time.Now().Format("2006-01-02 15:04:05")
		strText := fmt.Sprintf(templet, t, ssrAddr2Str)
		file.WriteString(strText)
		fmt.Print(">>> Git同步···\n")
		autogit("git add .")
		autogit("git commit -m \"ssr节点分享，每日更新\"")
		autogit("git remote add origin git@github.com:BeanWei/freeSSR.git")
		autogit("git pull origin master")
		autogit("git push origin master")
		fmt.Print("\n>>> Done")
	})

	c.Visit("https://doub.io/sszhfx/")
}

func autogit(strCmd string) {
	stout, err := exec.Command("cmd", "/C", strCmd).Output()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(stout))
}
