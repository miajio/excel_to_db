package main

import (
	"github.com/miajio/excel_to_db/cmd"

	"github.com/miajio/excel_to_db/cfg"
)

func main() {
	if err := cmd.Run(); err != nil {
		cfg.Log.Errorf("excel run fail: %s", err)
	}
}
