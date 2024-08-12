package sound

import (
	"bytes"
	_ "embed"
	"gorpg/assets"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const SampleRate = 48000
const sfxPath = "audio/sfx/WAV Files/SFX/"

var audioContext = audio.NewContext(SampleRate)

type AudioEmitter struct {
	Streams map[string]*audio.Player
	panning float64
}

func NewEmitter() *AudioEmitter {
	return &AudioEmitter{
		make(map[string]*audio.Player),
		0,
	}
}
func (ae *AudioEmitter) NewPlayer(id string, auto bool) {
	if _, ok := ae.Streams[id]; !ok {
		var p *audio.Player
		b, err := assets.AudioAssets.ReadFile(id)
		if err != nil {
			panic(err)
		}
		if auto {
			decoded, err := wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(b))
			if err != nil {
				panic(err)
			}
			stream := audio.NewInfiniteLoop(decoded, decoded.Length())
			p, err = audioContext.NewPlayer(stream)
			if err != nil {
				panic(err)
			}
		} else {
			decoded, err := wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(b))
			if err != nil {
				panic(err)
			}

			p, err = audioContext.NewPlayer(decoded)
			if err != nil {
				panic(err)
			}
		}
		ae.Streams[id] = p
	}
}
