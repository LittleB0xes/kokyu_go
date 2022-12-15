package soundBox

import (
	"bytes"
	"embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
//"github.com/hajimehoshi/ebiten/v2/audio/wav"

type SoundAsset int64

const (
	Sword1 SoundAsset = iota
	Sword2
	Heavy
	Huh1
	Huh2
	Huh3
	Death
	Ambiance
	Beat
)

type SoundBox struct {
	context        *audio.Context
	ambiancePlayer *audio.Player
	beatPlayer     *audio.Player
	allSound       [7]*audio.Player
}

func NewSoundBox(f embed.FS) *SoundBox {
	sampleRate := 44100
	context := audio.NewContext(sampleRate)

	assetsPath := [7]string{
		"assets/sounds/sword1.wav",
		"assets/sounds/sword2.wav",
		"assets/sounds/sword_heavy.wav",
		"assets/sounds/huh_1.wav",
		"assets/sounds/huh_2.wav",
		"assets/sounds/huh_3.wav",
		"assets/sounds/death.wav",
	}

	var allSound [7]*audio.Player
	for i, path := range assetsPath {

		d, err := f.ReadFile(path)
		if err != nil {
			panic(err)
		}
		read, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(d))
		if err != nil {
			panic(err)
		}
		s, err := context.NewPlayer(read)
		allSound[i] = s
	}

	data, err := f.ReadFile("assets/sounds/heart_beat.ogg")
	if err != nil {
		panic(err)
	}
	reader, err := vorbis.DecodeWithSampleRate(sampleRate, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	loop := audio.NewInfiniteLoop(reader, reader.Length())

	beatPlayer, err := context.NewPlayer(loop)

	if err != nil {
		panic(err)
	}

	data, err = f.ReadFile("assets/sounds/amb_intro.ogg")
	if err != nil {
		panic(err)
	}
	reader, err = vorbis.DecodeWithSampleRate(sampleRate, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	loop = audio.NewInfiniteLoop(reader, reader.Length())

	ambiancePlayer, err := context.NewPlayer(loop)
	if err != nil {
		panic(err)
	}

	return &SoundBox{
		context,
		ambiancePlayer,
		beatPlayer,
		allSound,
	}
}

func (s *SoundBox) SBPlay(sound SoundAsset) {
	switch sound {
	case Beat:
		s.beatPlayer.Play()
	case Ambiance:
		s.ambiancePlayer.Play()
	default:
		s.allSound[sound].Rewind()
		s.allSound[sound].Play()
	}

}

func (s *SoundBox) SBStop(sound SoundAsset) {
	switch sound {
	case Beat:
		s.beatPlayer.Pause()
		s.beatPlayer.Rewind()
	case Ambiance:
		s.ambiancePlayer.Pause()
		s.ambiancePlayer.Rewind()

	}
}
