package ivona

import (
	ivonago "github.com/jpadilla/ivona-go"
	"log"
	"sync"
)

var (
	ivonaRequestLock = &sync.Mutex{}
	client           *ivonago.Ivona
)

func InitIvona(accessKey, secretKey string) {
	client = ivonago.New(accessKey, secretKey)
}

func GetTTSFromIvona(provider, text, voice string, keepFile bool) {
	go func() {
		//I want only 1 API call at a time right now, later turn this into a semaphore
		ivonaRequestLock.Lock()
		defer ivonaRequestLock.Unlock()
		log.Println("--- GetTTSFromIvona: Got lock:", text, voice)

		options := ivonago.NewSpeechOptions(text)
		r, err := client.CreateSpeech(options)
		if err != nil {
			log.Println("Error getting response from Ivona: text:", err)
			return
		}
		_resp := &TTSResponse{
			Text:     text,
			Audio:    r.Audio,
			Provider: provider,
			KeepFile: keepFile,
		}
		responseQueue.Push(_resp)
	}()
}
