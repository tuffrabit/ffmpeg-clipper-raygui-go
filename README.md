# ffmpeg-clipper-raygui-go

## Disclaimer

No AI was used in the making of this. I like to learn things for myself.

## Description
Cross platform desktop GUI tool for making video clips. Download an executable from the releases and run it. It will ask you to open a directory (pick a target folder) full of video files. This tool takes all the complicated parts of using FFmpeg to edit simple video clips and hides it behind a (hopefully) easy to use interface. You should have FFmpeg, FFprobe, and FFplay either in the target directory or installed globally on your system.

## Expectations
This project is mostly a raylib + raygui + Go testbed. Any functionality that actually works is secondary. The primary goal is to see if it's possible and at all ergonomic to build a cross platform desktop GUI using Go as the application glue code and raylib as the rendering engine and raygui as the GUI/layout engine. An ancillary goal was to see if Zig can be used as a drop in C/C++ compiler for non trival Go projects requiring significant CGO.

## Thanks
This project sits on the shoulders of giants. Huge thanks to these projects:

- https://ffmpeg.org
- https://github.com/raysan5/raylib
- https://github.com/raysan5/raygui
- https://github.com/gen2brain/raylib-go
- https://ziglang.org
