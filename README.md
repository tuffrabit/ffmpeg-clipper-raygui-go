# ffmpeg-clipper-raygui-go

## Disclaimer

No AI was used in the making of this. I like to learn things for myself.

## Description
Cross platform desktop GUI tool for making video clips. Download an executable from the releases and run it. It will ask you to open a directory (pick a target folder) full of video files. This tool takes all the complicated parts of using FFmpeg to edit simple video clips and hides it behind a (hopefully) easy to use interface. You should have FFmpeg, FFprobe, and FFplay either in the target directory or installed globally on your system.

## Expectations
This project is mostly a raylib + raygui + Go testbed. Any functionality that actually works is secondary. The primary goal is to see if it's possible and at all ergonomic to build a cross platform desktop GUI using Go as the application glue code and raylib as the rendering engine and raygui as the GUI/layout engine. An ancillary goal was to see if Zig can be used as a drop in C/C++ compiler for non trival Go projects requiring significant CGO.

## Learnings after reaching v1.0.0
TL;DR up front: Totally possible to build crossplatform GUI apps this way. Wouldn't personally try a large multi-team member project without some pain points resolved. Great for tools. Performance is bonkers good.

Long version...

There's some very real strengths with an approach like this. Also some cons, most of the cons can be resolved with some more work. Two big cons, that may be deal breakers for some people, likely can't be resolved or resolved easily.

Pros:

- This thing is FAST and small. Even with the FPS set to 60 (its basically a game so yeah there are frames per second to worry over), it wouldn't break 1% CPU or GPU usage on any of my test machines (which vary wildly in hardware specs). Memory usage is stable around 30mb. Binary size is under 15mb depending on the platform. Yes yes, you can get smaller and faster with something else lower level. But then you're writing a crossplatform desktop GUI app in something lower level than Go....... have fun. This is quite a bit better than something web based like Electron or a webview.

- Logic is simple. The entire point of raylib is to provide a straightforward programming interface for drawing stuff on the screen and allow people to escape overly complicated game engine editors. That idea holds here. Anything you want to do, you have to implement it. To me that is a huge pro.

- Goroutines are just fabulous for GUI work. The ease of dropping out of the main drawing thread to a concurrent execution context is SSSSOOO ergonomic. I've done this in Windows Forms and WPF and Python, Go is stupid better in this specific way.

Cons:

- Static window size. As far as I can tell raylib won't let you create a window that is resizable. If I'm wrong please tell me so I can play with it. This is likely a deal breaker for certain types of apps.

- Native GUI/styles. Yeah that's not a thing here. Personally I don't care about that. I get that this may be a big issue for some people and some apps.

- Layout? What layout? As it stands right now, your only option for layout is hard coded rectangle dimensions and x/y absolute positioning. I found an ok way of dealing with that via compile time CONST math. It's my personal opinion that a layout engine is needed before this approach is ready for anything other than small/internal tools. A larger customer facing app would be painful without a layout engine.

Architecture:

I didn't have a plan for this app other than "just make things work and see how it feels". Code architecture is all over the place. I did settle into a sort of global state management approach, which did work well enough. I'm not confident how well that would scale to a larger app. Before I try this again I would want to nail down some conventions around event/message propagation, layout, and state. State gets funny though, there's GUI control state (scroll index, select index, etc) and then there's business logic/entity state. I would really want to be intentional with a lot of that before another try at this.

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

- better builds
- more video stats (length, filesize, etc)
- choose new video dir
- modal box message wrapping
- bug with modals not closing on X press
- preserve encoding preset selection
- app config settings (framerate, gui style)
- new video size helper since ffmpeg is so damn sensitive to resolution
