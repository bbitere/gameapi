package utils

import (
	"fmt"
	"log"
	"os"
	debug "runtime/debug"
	"strings"
	"time"
)

// utility method: string: substring
func Str_SubString(s string, index int, len int) string {
	return s[index : index+len]
}

// utility method: clone map[string][T]
func Util_CloneMapString[T any](dict *map[string]T) map[string]T {

	var ret = map[string]T{}
	for key, val := range *dict {
		ret[key] = val
	}
	return ret
}
type Aaa struct {m []string }
// utility method: convert from map to array. (values only)
func Util_FromMapToArray[T any](dict *map[string]T) []T {

	var ret = []T{}
	for _, val := range *dict {
		Arr_Append(&ret, val)
	}
	
	return ret
}

// utility method: convert from map to array. (keys only)
func Util_FromMapKeysToArray[T any](dict *map[string]T) []string {

	var ret = []string{}
	for key := range *dict {
		Arr_Append(&ret, key)
	}
	return ret
}

// Used to replace operator ?:
func IFF[T any](b bool, s1 T, s2 T) T {
	if b {
		return s1
	} else {
		return s2
	}
}

func Util_Nop(){

}


func StartLogPanic() {

	currentDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	fmt.Printf("Current Dir: %s", currentDir);

	file, err1 := os.OpenFile("logfile.txt", os.O_APPEND, os.ModeAppend)
	if err1 != nil {
		log.Fatal("Cannot create log file: ", err1)
	}
	defer file.Close()
	log.SetOutput(file)

	var currentTime = time.Now()
	var formated    = currentTime.Format( "2025-01-02 15:04:05" );

	log.Printf("Start log at : %s", formated);
}

func LogPanic() {

	file, err1 := os.OpenFile("logfile.txt", os.O_APPEND, os.ModeAppend)
	if err1 != nil {
		log.Fatal("Cannot create log file: ", err1)
	}
	defer file.Close()
	log.SetOutput(file)

	if err := recover(); err != nil {
		if err != nil {

			var linesText = string(debug.Stack())
			var lines = strings.Split(linesText, "\n")

			for i := 0; i < len(lines) && i <= 6; i++ {
				Arr_RemoveAt(&lines, 0)
			}
			var linesCleaned = strings.Join(lines, "\n")
			log.Printf("panic occurred: %v %s", err, linesCleaned)
		}
		log.SetOutput(os.Stdout)
		log.Printf("panic occurred: write error in logfile.txt")
	}
}