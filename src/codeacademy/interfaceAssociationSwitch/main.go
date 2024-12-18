package main

import "fmt"

func main() {
    var mediaPlayer Media = Video{frameRate: 24}
    
    // Perform a type assertion
    if a, ok := mediaPlayer.(Audio); ok {
        fmt.Println("This is an audio media with duration:", a.duration)
    } else {
        fmt.Println("This media is not audio.")
    }
    
    // Use a type switch to handle different types
    switch v := mediaPlayer.(type) {
      case Audio:
          fmt.Println("Handling audio with duration:", v.duration)
      case Video:
          fmt.Println("Handling video with frame rate:", v.frameRate)
      default:
          fmt.Println("Unknown media type")
    }
}

type Media interface {
    Play()
}

type Audio struct {
    duration int
}

func (a Audio) Play() {
    fmt.Println("Playing audio for", a.duration, "seconds")
}

type Video struct {
    frameRate int
}

func (v Video) Play() {
    fmt.Println("Playing video at", v.frameRate, "fps")
}
