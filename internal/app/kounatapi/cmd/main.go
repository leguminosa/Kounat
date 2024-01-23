package cmd

import (
	"context"
	"log"

	"github.com/leguminosa/kounat/internal/app/kounatapi"
	"github.com/leguminosa/kounat/internal/tools/config"
)

func Main() {
	ctx := context.Background()
	cfg := config.New()

	server, err := kounatapi.InitServer(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(server.Start())
}
