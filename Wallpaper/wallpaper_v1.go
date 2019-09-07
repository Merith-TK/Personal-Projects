	package main

	import (
		"fmt"
		"os"
		"time"
		"github.com/reujab/wallpaper"
	)

	var URL = "https://mgo.azureedge.net/cloudcache/7/c/2/f/3/4/7c2f345bdfcadb8a3faf483ebaa2e9aea712bbdb.jpg"
	var file = ""

	func main() {
		
		//Check to see if image is provided
		if len(os.Args) >= 2 {
			file = os.Args[1]
		}
		
		//If no image provided, set wallpaper to Default Windows 10 wallpaper
		if file == "" {
			wallpaper.SetFromURL(URL)
			fmt.Println("")
			fmt.Println("To use, drag and drop a image onto the exe, as if you were moving it to a folder.")
			fmt.Println("Default behavior is to reset to the default windows wallpaper")
			
			time.Sleep(30 * time.Second)
		} else {
		// If a image is provided, set that as the wallpaper
			wallpaper.SetFromFile(file)
		}
	}
