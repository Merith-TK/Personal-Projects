# Personal-Projects
 * Code.go
    A utility to launch VSCode portable from a PortableApps platform while setting the path enviroment for custom extensions
 * Home.go
    A tool to find the root directory that mingw is installed

 * Wallpaper Tool
    This utility is used to change the user's wallpaper. Originally designed for students who's
    computer IT department have locked down the Wallpaper settings. This program bypasses the 
    blocked settings by directly calling (to my knowledge) `user32.dll` to apply the settings
    * Libraries used
      * [Wallpaper, Used to set the users wallpaper](https://github.com/reujab/wallpaper), Used in V1, and V2
      * [Walk, Used to make a proper window'd GUI](github.com/lxn/walk), used in V2
    * TODO
      [ ] Set default window size
      [ ] Provide prebuilt binaries    
