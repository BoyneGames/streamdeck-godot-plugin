package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
)

type Plugin struct {
	Port          string
	PluginUUID    string
	RegisterEvent string
	Info          PluginInfo
}

type PluginInfo struct {
	Application struct {
		Font            string
		Language        string
		Platform        string
		PlatformVersion string
		Version         string
	}
	Colors struct {
		ButtonPressedBackgroundColor string
		ButtonPressedBorderColor     string
		ButtonPressedTextColor       string
		DisabledColor                string
		HighlightColor               string
		MouseDownColor               string
	}
	DevicePixelRatio string
	Devices          []struct {
		Id   string
		Name string
		Size struct {
			Columns int
			Rows    int
		}
		Type int
	}
	Plugin struct {
		Uuid    string
		Version string
	}
}

var (
	plugin               Plugin
	streamDeckConnection *websocket.Conn
	godotConnection      *websocket.Conn
	config               *ini.File
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	writePluginFile()
	registerEvent()
	setupWebSocketEndpoint()
}

func writePluginFile() {
	var pluginInfoJson string

	flag.StringVar(&plugin.Port, "port", "", "Plugin port")
	flag.StringVar(&plugin.PluginUUID, "pluginUUID", "", "Plugin UUID")
	flag.StringVar(&plugin.RegisterEvent, "registerEvent", "", "Plugin register event")
	flag.StringVar(&pluginInfoJson, "info", "", "Plugin info")
	flag.Parse()

	json.Unmarshal([]byte(pluginInfoJson), &plugin.Info)

	file, _ := ini.Load("plugin.ini")
	config = file
}

func registerEvent() {
	client, _, _ := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://127.0.0.1:%s", plugin.Port), nil)

	streamDeckConnection = client

	go listenEvents()

	streamDeckConnection.WriteJSON(&struct {
		Event string `json:"event"`
		UUID  string `json:"uuid"`
	}{
		Event: plugin.RegisterEvent,
		UUID:  plugin.PluginUUID,
	})
}

func listenEvents() {
	var event interface{}

	for {
		err := streamDeckConnection.ReadJSON(&event)

		if err != nil {
			logMessage(fmt.Sprint(err))
		}

		if godotConnection != nil {
			godotConnection.WriteJSON(event)
		}
	}
}

func setupWebSocketEndpoint() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, _ := upgrader.Upgrade(w, r, nil)
		godotConnection = ws
	})

	http.ListenAndServe(fmt.Sprintf(":%s", config.Section("bridge").Key("port").String()), nil)
}

func logMessage(message string) {
	socketFile, _ := os.Create("log")

	defer socketFile.Close()

	socketFile.WriteString(message)
}
