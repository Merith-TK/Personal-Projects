# Personal-Projects
 * Code.go
    A utility to launch VSCode portable from a PortableApps platform while setting the path enviroment for custom extensions
 * Home.go
    A tool to find the root directory that mingw is installed
    Uses the `mount` command to find the install path, 
 * Wallpaper Tool
    This utility is used to change the user's wallpaper. Originally designed for students who's
    computer IT department have locked down the Wallpaper settings. This program bypasses the 
    blocked settings by directly calling (to my knowledge) `user32.dll` to apply the settings
    * Libraries used
      * [Wallpaper, Used to set the users wallpaper](https://github.com/reujab/wallpaper), Used in V1, and V2
      * [Walk, Used to make a proper window'd GUI](github.com/lxn/walk), used in V2
    * TODO

     - [ ] Set default window size
     - [ ] Provide prebuilt binaries    

 * Compile
    This utility makes compiling some exe programs easier, especially when embedding a icon. 
    Default Golang icon comes with the source because i needed a icon to test it. Also lets you test it yourself
    * DEPENDANCIES
      * [rsrc, used as the core of this entire program](https://github.com/akavel/rsrc)
    * How to install!
      * `go get github.com/Merith-TK/Personal-Projects/compile `

      * `go get https://github.com/akavel/rsrc`

      FUN FACT! I actually build the binary of this program using `go run` due to how it behaves! 