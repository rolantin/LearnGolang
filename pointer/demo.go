package main

import (
	"fmt"
)

type storyPage struct {
	text string
	nextPage *storyPage
}


//func playStory(page *storyPage)  {
func (page *storyPage) playStory() {
	if page == nil{
		return
	}
	fmt.Println(page.text)
	//playStory(page.nextPage)
	page.nextPage.playStory()
}

func main() {

	page1 := storyPage{"it was a dark and stormy night",nil}
	page2 := storyPage{"You are alone,and you need to find the sacred helmet before the bad guys do",nil}
    page3 := storyPage{"You see a troll ahead.",nil}
	fmt.Println("test start -----------------------",&page2)
    page1.nextPage = &page2
    page2.nextPage = &page3

    page1.playStory()
    //playStory(&page1)

	//scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("It was a dark and stormy night.")
	//scanner.Scan() // --> 等待输入
	//fmt.Println("You are alone,and you need to find the sacred helmet before the bad guys do")
	//scanner.Scan()
	//fmt.Println("You see a troll ahead.")
	//scanner.Scan()
}
