package main

import ("fmt";"net/http")

func CheckLink(link string, c chan string){
	_, err := http.Get(link)
	if err != nil{
		c<-link+" - Might be down i think!"
		return
	}
	c<-link+" - Yep its up!"
}

func main(){
	links := []string{"https://google.com","https://amazon.com","https://a413.com/"}
	c:=make(chan string)

	for _, link := range links{
		go CheckLink(link, c)
	}

	//This will print 3 values that we get from links!
	for i:=0; i<len(links);i++{
		fmt.Println(<-c)
	}
}


