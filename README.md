# Personal-Projects
 ### Latest information
 * reformated repo layout to support `go get` each program
 * removed old wallpaper code
 * fixed bug with launcher where you could only send one argument
 * made a folder to store functions and resources to import with `github.com/Merith-TK/Personal-Projects/functions`
 * Updated Readme

 ### Tools in this Repo 
 ~~* Code.go
    A utility to launch VSCode portable from a PortableApps platform while setting the path enviroment for custom extensions
    * Has gone through two rewrites~~
   * Launcher.go
    A utility to run any executable with specific arguments and environment variables.
    `example.json`
```json
{
  "application":"echo",
  "applicationArgs":"| grep hello=world",
  "environment": {
    "hello":"world!"
  }
}
```
  * Replaces Code.go
 * `functions/home.go`
    * A tool to find the root directory that mingw is installed
    Uses the `mount` command to find the install path,
    I intend on making this program a library for msys2 projects using go
      * This has gone through ONE complete rewrite
 * Wallpaper Tool
    * This utility is used to change the user's wallpaper. Originally designed for students who's
    computer IT department have locked down the Wallpaper settings. This program bypasses the 
    blocked settings by directly calling (to my knowledge) `user32.dll` to apply the settings
    * Libraries used
      * [Wallpaper, Used to set the users wallpaper](https://github.com/reujab/wallpaper), Used in V1, and V2
      * [Walk, Used to make a proper window'd GUI](github.com/lxn/walk), used in V2
    * TODO

     - [ ] Set default window size
     - [ ] Provide prebuilt binaries    

 * Compile
    ### DISCLAIMER
      * This tool was made to help me personally, i highly reccomend downloading and tweaking the tool to suit your needs. only tested on windows. it is in no way meant to to be used in production.
    * This utility makes compiling some exe programs easier, especially when embedding a icon. 
    Default Golang icon comes with the source because i needed a icon to test it. Also lets you test it yourself!
    * USAGE!
      * Just go into the root of your project, make sure there is a `manifest.xml` file, and run `compile`
      * OR! if you want to compile a specific go file. `compile main.go`
    * DEPENDANCIES
      * [rsrc, used as the core of this entire program](https://github.com/akavel/rsrc)
    * How to install!
      * `go get github.com/Merith-TK/Personal-Projects/compile `
      * `go get github.com/akavel/rsrc`

    * How to test? 
      * This program is only validated to work on windows 7 and windows 10. How you can check to see if this program works. go to your GOPATH/bin folder and look for `compile.exe`, it should have a go-gopher icon.
      This is how you test this program because i bundled the syso file, which was BUILT by this program. crazy i know. it was built by using `go run` on `compile.go`
      

      FUN FACT! I actually built the binary of this program as i was making it due to how it behaves! 
