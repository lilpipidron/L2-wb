package pattern

import "fmt"

// Применимость: Хотим разложить подсистему на отдельные слои
// Плюсы: Предоставляет ограничить api
// Минусы: Может стать god классом

type AudioSystem struct {
	filename string
}

func (a *AudioSystem) LoadAudio(filename string) {
	fmt.Println("Loading audio file", filename)
}

func (a *AudioSystem) PlayAudio() {
	if a.filename != "" {
		fmt.Println("Playing audio file", a.filename)
	} else {
		fmt.Println("No audio file found")
	}
}

type VideoSystem struct {
	filename string
}

func (v *VideoSystem) LoadVideo(filename string) {
	fmt.Println("Loading video file", filename)
}

func (v *VideoSystem) PlayVideo() {
	if v.filename != "" {
		fmt.Println("Playing video file", v.filename)
	} else {
		fmt.Println("No video file found")
	}
}

type Facade struct {
	audioSystem *AudioSystem
	videoSystem *VideoSystem
}

func NewFacade() *Facade {
	return &Facade{
		audioSystem: &AudioSystem{},
		videoSystem: &VideoSystem{},
	}
}

func (f *Facade) PlayMedia(filename string) {
	f.audioSystem.LoadAudio(filename)
	f.videoSystem.LoadVideo(filename)
	f.audioSystem.PlayAudio()
	f.videoSystem.PlayVideo()
}
