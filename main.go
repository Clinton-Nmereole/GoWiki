package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"jaytaylor.com/html2text"
)




func main () {
	var port int
	flag.IntVar(&port, "p", 8080, "port number");
	flag.Parse();
	fmt.Printf("port: %d\n", port);
	response, err := http.Get("https://en.wikipedia.org/wiki/"+ os.Args[1]);
	if err != nil {
		panic(err);
	}
	fmt.Println(response.Status);
	defer response.Body.Close();
	body, err := io.ReadAll(response.Body);
	if err != nil {
		panic(err);
	}
	htmlText := string(body);
	text, err := html2text.FromString(htmlText, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

}



