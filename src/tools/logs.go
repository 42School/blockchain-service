package tools

import (
	"log"
	"runtime"
)

func getFrame(skipFrames int) runtime.Frame {
	targetFrameIndex := skipFrames + 2
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)
	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}
	return frame
}

func LogsDev(msg string) {
	if Env == "dev" || Env == "DEV" || Env == "Dev" {
		log.Println("Developper mode:", msg)
	}
}

func LogsError(_err error) {
	caller := getFrame(2).Function[39:]
	string := "Error - " + caller + ":"
	log.Println(string, _err)
}

func LogsMsg(msg string) {
	log.Println(msg)
}
