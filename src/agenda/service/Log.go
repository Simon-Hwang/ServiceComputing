package service

import "os"

type LogFile struct{
	file string
}


func (Log *LogFile) Write(p []byte) (int, error) {
    f, err := os.OpenFile(Log.file, os.O_CREATE|os.O_APPEND, 0666)
    defer f.Close()

    if err != nil {
        return -1, err
    }
    return f.Write(p)
}
