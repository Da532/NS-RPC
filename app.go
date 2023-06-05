package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/hugolgst/rich-go/client"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type App struct {
	ctx context.Context
}

type Games []struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

type Game struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

type Pins []string

var gamesList Games
var connErr bool = false

const clientID string = "1114647533562646700"
const gamesURL string = "https://raw.githubusercontent.com/Da532/NS-RPC/master/games.json"

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := a.GetGamesData()
	if err != nil {
		panic(err)
	}
	err = client.Login(clientID)
	if err != nil {
		connErr = true
	}
	err = client.SetActivity(client.Activity{
		LargeImage: "home",
		Details:    "Home",
		State:      "Idle",
	})
	if err != nil {
		panic(err)
	}
}

func (a *App) shutdown(ctx context.Context) {
	client.Logout()
}

func (a *App) CheckConn() bool {
	return connErr
}

func (a *App) Reconnect() bool {
	err := client.Login(clientID)
	if err != nil {
		return false
	}
	err = client.SetActivity(client.Activity{
		LargeImage: "home",
		Details:    "Home",
		State:      "Idle",
	})
	if err != nil {
		return false
	}
	connErr = false
	return true
}

func (a *App) GetGamesData() error {
	resp, err := http.Get(gamesURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &gamesList)
	if err != nil {
		return err
	}
	sort.Slice(gamesList, func(i, j int) bool {
		return gamesList[i].Title < gamesList[j].Title
	})
	return nil
}

func (a *App) GetGamesList() string {
	data, err := json.Marshal(gamesList)
	if err != nil {
		a.GetGamesData()
		return err.Error()
	}
	return string(data)
}

func (a *App) SetGame(title string, status string) {
	var selectedGame Game
	for _, game := range gamesList {
		if game.Title == title {
			selectedGame = game
			break
		}
	}
	if selectedGame.Title != "" {
		err := client.SetActivity(client.Activity{
			LargeImage: selectedGame.Img,
			Details:    selectedGame.Title,
			State:      cases.Title(language.English).String(status),
		})
		if err != nil {
			panic(err)
		}
	}
}

func LoadPinJson() Pins {
	var pins Pins
	configDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configDir = filepath.Join(configDir, "NS-RPC")
	_, err = os.Stat(configDir)
	if err != nil {
		err = os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	pinsJson, err := os.Open(filepath.Join(configDir, "pinned.json"))
	if err == nil {
		defer pinsJson.Close()
		bytes, _ := io.ReadAll(pinsJson)
		json.Unmarshal(bytes, &pins)
	}
	return pins
}

func (a *App) PinGame(title string) {
	pins := LoadPinJson()
	configDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	removedPin := false
	for i, pin := range pins {
		if pin == title {
			pins = slices.Delete(pins, i, i+1)
			removedPin = true
			break
		}
	}
	if !removedPin {
		pins = append(pins, title)
	}
	file, _ := json.Marshal(pins)
	err = os.WriteFile(filepath.Join(configDir, "NS-RPC", "pinned.json"), file, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (a *App) GetPins() string {
	pins := LoadPinJson()
	var pinMenu Games
	for _, pin := range pins {
		pinMenu = append(pinMenu, Game{Title: pin, Img: ""})
	}
	if len(pinMenu) == 0 {
		pinMenu = append(pinMenu, Game{Title: "No Pins!", Img: ""})
	}
	data, _ := json.Marshal(pinMenu)
	return string(data)
}

func (a *App) IsMac() bool {
	return runtime.GOOS != "windows"
}
