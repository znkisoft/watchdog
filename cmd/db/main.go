package main

import (
	"os"
	"path/filepath"

	"github.com/znkisoft/watchdog/pkg/db"
	"github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

const fileName = "watchdog.schema.sqlite"

func main() {
	targetDir := os.Args[1]
	fs := filepath.Join(targetDir, fileName)
	if _, err := os.Stat(fs); err == nil || os.IsExist(err) {
		return
	}

	query := db.CreateDBQuery([]proto.Message{
		&schema.Userver{},
	})

	_ = os.WriteFile(fs, []byte(query), 0644)
}
