package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

// The content of the response is in the response field 'Body' which is of type io.ReadCloser
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "URL")
		os.Exit(1)
	}
	url := os.Args[1]

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Println(response.Status, response.Proto)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Print("DumpResponse:", string(b))

	// contentTypes := response.Header["Content-Type"]
	// if !acceptableCharset(contentTypes) {
	// 	fmt.Println("Cannot handle", contentTypes)
	// 	os.Exit(4)
	// }

	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print("response body:", string(buf[0:n]))
	}
	os.Exit(0)
}

// func acceptableCharset(contentTypes []string) bool {
// 	// each type is like [text/html; charset=UTF-8]
// 	// we want the UTF-8 only
// 	for _, cType := range contentTypes {
// 		if strings.Index(strings.ToUpper(cType), "UTF-8") != -1 {
// 			return true
// 		}
// 	}
// 	return false
// }
