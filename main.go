package main

import (
	"github.com/halilylm/utils/logger"
	"github.com/halilylm/utils/logger/sugar"
	"log"
)

func main() {
	z, err := sugar.NewApiLogger(&sugar.Options{
		Level:           logger.WarnLevel,
		DevelopmentMode: true,
		InitialFields: map[string]any{
			"release": "v1.0.0",
		},
	})
	defer func() {
		_ = z.Sync()
	}()
	if err != nil {
		log.Fatalln(err)
	}
	z.Fatal("testing")
}
