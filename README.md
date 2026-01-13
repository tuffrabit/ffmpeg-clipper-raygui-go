# ffmpeg-clipper-raygui-go

## Disclaimer

No AI was used in the making of this. I like to learn things for myself.

## Description
Cross platform desktop GUI tool for making video clips. Download an executable from the releases and run it. It will ask you to open a directory (pick a target folder) full of video files. This tool takes all the complicated parts of using FFmpeg to edit simple video clips and hides it behind a (hopefully) easy to use interface. You should have FFmpeg, FFprobe, and FFplay either in the target directory or installed globally on your system.

## Expectations
This project is mostly a raylib + raygui + Go testbed. Any functionality that actually works is secondary. The primary goal is to see if it's possible and at all ergonomic to build a cross platform desktop GUI using Go as the application glue code and raylib as the rendering engine and raygui as the GUI/layout engine. An ancillary goal was to see if Zig can be used as a drop in C/C++ compiler for non trival Go projects requiring significant CGO.

## Build

Building is a mess. Building for Windows and Linux with Zig works. Linux was a pain because the Zig C compiler operates like CLANG and not GCC and the C GLFW bindings in this project REALLY prefer a GCC compiler. Haven't yet gotten Mac to build with Zig. For now xgo works fine. Your options are to get a Zig compiler setup and run one of the build scripts or just build everything with xgo.

```
xgo --targets=windows/amd64,darwin/* -ldflags "-s -w" -out ffmpeg-clipper-raygui-go-{version} -dest . .
```

## Thanks
This project sits on the shoulders of giants. Huge thanks to these projects:

- https://ffmpeg.org
- https://github.com/raysan5/raylib
- https://github.com/raysan5/raygui
- https://github.com/gen2brain/raylib-go
- https://ziglang.org
- https://github.com/techknowlogick/xgo

## TODO

-- mac build
-- raygui styles
- more video stats (length, filesize, etc)
- choose new video dir
- modal box message wrapping
- bug with modals not closing on X press
- preserve encoding preset selection
- app config settings (framerate, gui style)
-- ffmpeg local dir testing
- new video size helper since ffmpeg is so damn sensitive to resolution
