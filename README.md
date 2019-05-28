# DVDTS - DVD Term Screensaver

DVD Screensaver style screensaver for terminal.
Instead of the logo of DVD, it uses the distro name/os name.

## Preview

![preview gif](readme_assets/dvdts.gif)

The preview shows use with color cyclying enabled and one with single color.<br><br>

if the gif is laggy, see [video](readme_assets/dvdts.mp4)

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
        starting/only color for the bouncing text (default "blue")
-s int
        speed of text [more is slower] (default 10)
```

### In use bindings
|Key(s)     |Action    |
|-----------|----------|
|q or Ctrl+c|Quit      |
|a|Toggle color cycling|

## Dependencies:
- to compile: [Go](https://golang.org/)
- `lsb_release`

## Build
```
go build

./dvdts
```

## Uses
[termui](https://github.com/gizak/termui)