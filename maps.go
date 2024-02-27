package main

import "fmt"

func main() {
	fmt.Println("hasi")
	websites := map[string]string{
		"google.com":          "https://google.com",
		"Amazon web services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google.com"])
	websites[`facebook.com`] = `http://www.facebook.com/haseebtariq`
	// delete element from the map using its key
	delete(websites, "google.com")
	fmt.Println(websites)

	value, ok := websites["facebook.com"] // check if a specific key is present in the map or not
	if !ok {                              // if it's not present then do something else
		fmt.Println("not found!")
	} else {
		fmt.Println(value)
	}

	for key, value := range websites {
		fmt.Println(key, ": ", value)
	}
}
