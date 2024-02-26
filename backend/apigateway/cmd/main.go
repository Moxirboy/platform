package main

// import (
// 	configs "gateway/internal/config"
// 	"gateway/internal/server"
// 	logger "gateway/pkg/log"
// )
import (
	configs "gateway/internal/config"

	"gateway/internal/server"
	logger "gateway/pkg/log"

	"net"
	"fmt"
)

func main() {
	inter()

	var (
		config = configs.Load()
	)
	log := logger.NewLogger(config.Logger.Level, config.Logger.Encoding)
	log.InitLogger()


	log.Infof(

		"AppVersion: %s, LogLevel: %s, Mode: %s",
		config.AppVersion,
		config.Logger.Level,
		config.Server.Environment,
	)
	s := server.NewServer(config, log)


	
        log.Fatal(s.Run(":5005"))
  
	
        
 
}

func inter(){
	interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }


    // Iterate over each network interface
    for _, iface := range interfaces {
        fmt.Println("Interface:", iface.Name)

        // Get a list of unicast IP addresses associated with the interface
        addrs, err := iface.Addrs()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        // Iterate over each IP address
        for _, addr := range addrs {
            fmt.Println("  IP Address:", addr.String())
        }
    }
}

