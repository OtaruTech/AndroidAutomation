package main

import (
	"log"
	"machine/pkg/bridge"
	"machine/pkg/message"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/zserge/lorca"
)

// var getNameHandle message.MethodHandle = func(input message.Message) (output message.Message) {
// 	return message.Message{}
// }

var (
	supportGui bool
)

func main() {
	supportGui = false

	if(supportGui) {
		gui_main()
		return
	}
	// messagedaemon.InitDaemon()
	var host string
	var port int

	log.Println(os.Args)
	if len(os.Args) != 3 {
		log.Println("Usage: ./main.exe <host addr> <port>")
		return
	}
	host = os.Args[1]
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println("Usage: ./main.exe <host addr> <port>")
		return
	}
	bridge.Initialize()
	message.InitMessage(host, port)
	bridge.Start()
	for {
		time.Sleep(1000 * time.Second)
	}
}

func gui_main() {
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// Bind Go functions to JS
	ui.Bind("onConnect", func() {
		host := ui.Eval(`document.getElementById("host").value`)
		port := ui.Eval(`document.getElementById("port").value`)
		log.Println(host, ":", port)
		message.InitMessage(host.String(), port.Int())
		bridge.Initialize()
		bridge.Start()
	})

	// Load HTML after Go functions are bound to JS
	ui.Load("data:text/html," + url.PathEscape(`
	<html>
		<body>
			<div
				<span>Host:</span>
				<input id="host" type="text" placeholder="host addr">
			</div>
			<div>
				<span>Port:</span>
				<input id="port" type="number" placeholder="host addr">
			</div>
			<button onclick="onConnect()">Connect</button>
		</body>
	</html>
	`))

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
