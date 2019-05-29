# DVDTS - DVD Term Screensaver

DVD Screensaver style screensaver for terminal.
Instead of the logo of DVD, it uses the distro name/os name or custom text.

## Preview

![preview gif](readme_assets/dvdts.gif)

The preview shows one use with color cyclying enabled and one with single color.<br><br>

Preview shows one with custom text and color cycling enabled, and the default run.

If the gif is laggy, see [video](readme_assets/dvdts.mp4)

## Usage

### Example
```
dvdts -a -c green -s 5
```
This will start the colors from green, cycle because of `a` flag, and text will move twice as fast.

### Flags
```
-a      cycle through terminal colors
-c string
        color for the bouncing text (default "blue")
-s int
        speed of text [more is slower] (default 10)
-t string
        text to display (default <distro name>)
```

### In use bindings
|Key(s)     |Action    |
|-----------|----------|
|q or Ctrl+c|Quit      |
|a|Toggle color cycling|

## Build Dependencies:
- [Go](https://golang.org/)
- [termui](https://github.com/gizak/termui)

## Install
```
go install
```
