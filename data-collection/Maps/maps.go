package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":"https://google.com",
		"Amazon Web Services": "https://aws.com",
	}


	websites["Youtube"] = "https://youtube.com"

	fmt.Println(websites)
	fmt.Println(websites["Google"])


	delete(websites, "Google")
	fmt.Println(websites)


}
