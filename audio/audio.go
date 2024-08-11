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

type AudioController struct {
	context   *audio.Context
	streams   map[string]*audio.Player
	autoplay  map[string]bool
	volume128 int
	sfxQueue  []string
}

func (ac *AudioController) Update() {
}

func (ac *AudioController) PlaySFX(id string, auto bool){
    if v, ok := ac.streams[id]; ok {
        if !v.IsPlaying() {
            v.Rewind()
            v.Play()
        }
    } else {
        b, err := assets.AudioAssets.ReadFile(id)
        if err != nil {
            panic(err)
        }

        s,err := wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(b))
        if err != nil {
            panic(err)
        }
        p, err := ac.context.NewPlayer(s) 
         if err != nil {
            panic(err)
        }
        p.Play()
        ac.streams[id] = p
        ac.autoplay[id] = auto
    }
    
}

func NewController(context *audio.Context) (*AudioController, error) {
	controller := &AudioController{
		context:   context,
		streams:   make(map[string]*audio.Player),
        autoplay: make(map[string]bool),
		volume128: 128,
	}
	return controller, nil
}
