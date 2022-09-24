package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/znkisoft/watchdog/pkg/db"
	"github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

const fileName = "watchdog.schema.sql"

func main() {
	targetDir := "."
	if len(os.Args) > 1 {
		targetDir = os.Args[1]
	}
	fs := filepath.Join(targetDir, fileName)
	if _, err := os.Stat(fs); err == nil || os.IsExist(err) {
		log.Printf("file %s already exists", fs)
		return
	}

	query := db.CreateDBQuery([]proto.Message{
		&schema.Userver{},
	})

	_ = os.WriteFile(fs, []byte(query), 0644)
}
