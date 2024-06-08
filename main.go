package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"runtime"
)

//go:embed bin/mybinary
var embeddedFiles embed.FS

func main() {
    data, err := embeddedFiles.ReadFile("bin/mybinary")
    if err != nil {
        log.Fatalf("Error reading embedded file: %v", err)
    }

    var tmpFile *os.File
    if runtime.GOOS == "windows" {
        tmpFile, err = os.CreateTemp("", "mybinary-*.exe")
    } else {
        tmpFile, err = os.CreateTemp("", "mybinary-*")
    }
    
    if err != nil {
        log.Fatalf("Error creating temp file: %v", err)
    }
    defer tmpFile.Close()

    tempFilePath := tmpFile.Name()

    _, err = tmpFile.Write(data)
    if err != nil {
        log.Fatalf("Error writing to temp file: %v", err)
    }

    err = os.Chmod(tempFilePath, 0755)
    if err != nil {
        log.Fatalf("Error making temp file executable: %v", err)
    }

    fmt.Printf("Embedded binary extracted to: %s\n", tempFilePath)
}
