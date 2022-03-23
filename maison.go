package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://portfolio.florentriou.me/Commandes-d90773cb8ae54515a6d3de300a88481e")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)

	//data := doc.Find("div", "spellcheck", "false")

	//fmt.Println(data.Text())
	fmt.Println(doc.HTML())

	/*	links := doc.Find("pre", "id", "prettyprint").FindAll("a")
		for _, link := range links {
			fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
		}*/
}
