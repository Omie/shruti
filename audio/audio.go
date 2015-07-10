package audio

import (
	"crypto/md5"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	_CMD_AUDIO_PLAY  = "ffplay"
	_AUDIO_FILES_DIR = "./audiofiles"
)

func PlayAudioSlice(input []byte, keepFile bool) (err error) {
	log.Println("--- playAudioSlice")

	uuid := md5.Sum(uuid.NewV1().Bytes())
	uuidHex := fmt.Sprintf("/%x.mp3", uuid)
	filename := _AUDIO_FILES_DIR + uuidHex

	err = ioutil.WriteFile(filename, input, 0644)
	if err != nil {
		return err
	}
	if !keepFile {
		defer os.Remove(filename)
	}

	err = PlayAudioFile(filename)
	if err != nil {
		return err
	}

	return nil
}

func PlayAudioFile(inputFile string) (err error) {
	command := exec.Command(_CMD_AUDIO_PLAY, "-nodisp", "-autoexit", "-i", inputFile)
	return command.Run()
}

func PlayGreeting() {
	PlayAudioFile(_AUDIO_FILES_DIR + "/greeting.mp3")
}
