package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/massalabs/thyra-ipfs-plugin/pkg/ipfs"
	"github.com/massalabs/thyra-ipfs-plugin/pkg/plugin"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func pushData(c *gin.Context) {
	c.JSON(201, gin.H{"status": ""})
}

func fetchData(c *gin.Context) {

	c.JSON(200, gin.H{"status": ""})
}

func main() {

	//nolint:gomnd
	if len(os.Args) < 2 {
		panic("this program must be run with correlation id argument!")
	}

	pluginID := os.Args[1]

	standaloneMode := false

	if len(os.Args) == 3 && os.Args[2] == "--standalone" {
		standaloneMode = true
	}

	ipfs.Init()

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/fetch", fetchData)
	router.POST("/push", pushData)

	ln, _ := net.Listen("tcp", ":")

	log.Println("Listening on " + ln.Addr().String())
	if !standaloneMode {
		err := plugin.Register(pluginID, "IPFS Plugin", "Massalabs", "push & fetch IPFS data", ln.Addr())
		if err != nil {
			log.Panicln(err)
		}
	}

	err := http.Serve(ln, router)
	if err != nil {
		log.Fatalln(err)
	}
}
