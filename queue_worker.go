package main

import (
	"log"
	"time"
)

func RunResponseWorker() {
	log.Println("--- RunResponseWorker")
	for {
		ivonaResp := responseQueue.Poll()
		if ivonaResp == nil {
			log.Println("got nil, sleeping ...")
			time.Sleep(10 * time.Second)
			continue
		}
		log.Println("playing:", ivonaResp.Text)
		playAudioSlice(ivonaResp.Audio, ivonaResp.KeepFile)
		time.Sleep(2 * time.Second)
	}
}
