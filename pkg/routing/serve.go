package routing

import (
	"fmt"
	"github.com/abbasfisal/blog/pkg/config"
	"log"
)

func Serve() {
	r := GetRouter()
	configs := config.Get()

	if err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)); err != nil {
		log.Fatal(err)
	}
}
