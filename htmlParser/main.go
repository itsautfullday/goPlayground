package main
import (
    "bufio"
    "fmt"
	"os"
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func main(){
	//Returns a file pointer!
	f, err := os.Open("ex3.html")
	if(err != nil){
		fmt.Println("Error in reading file");
	}
	var result []Link
	fileReader := bufio.NewReader(f)
	z := html.NewTokenizer(fileReader)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		token := z.Token()
		
		if(token.Type == html.StartTagToken && token.Data == "a"){
			data := ""
			link := ""
			attributes := token.Attr
			for _, val := range attributes {
				if(val.Key == "href"){
					link = val.Val
				}
			}
			for {
				z.Next()
				nextToken := z.Token()
				if(nextToken.Type == html.EndTagToken && nextToken.Data == "a" ){
					break
				}
				if(nextToken.Type == html.TextToken){
					cleanedData := strings.TrimSpace(nextToken.Data)
					data += cleanedData + " "
				}	
			}
			ln := Link {Href :link , Text:strings.TrimSpace(data)}
			result = append(result,ln)
		}
	}

	fmt.Println("Lenght of result ", len(result))
	for _,val := range result{
		printLink(val)
	}
}

func printLink(ln Link){
	fmt.Println("Link ",ln.Href, "\nData ",ln.Text)
}

