package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"tsi.co/go-api/utils"
)

type ExampleStruct struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Text      string `json:"text"`
	Attack    int    `json:"attack"`
	CardClass string `json:"cardClass"`
	Cost      int    `json:"cost"`
	Health    int    `json:"health"`
	Rarity    string `json:"rarity"`
	CardType  string `json:"type"`
	// Set                string `json:"set"`
	// Durability         int    `json:"durability"`
	// Armor              int    `json:"armor"`
	// HideStats          bool   `json:"hideStats"`
	// HowToEarn          string `json:"howToEarn"`
	// HowToEarnGolden    string `json:"howToEarnGolden"`
	// TargetingArrowText string `json:"targetingArrowText"`
	// DbfId              int    `json:"dbfId"`
	// Artist             string `json:"artist"`
	// Collectible        bool   `json:"collectible"`
	// Elite              bool   `json:"elite"`
	// Faction            string `json:"faction"`
	// Mechanics          []string `json:"mechanics"`
	// Flavor    string `json:"flavor"`
}

func main() {

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.hearthstonejson.com/v1/124497/enUS/cards.collectible.json")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var data []ExampleStruct
		err = json.Unmarshal(body, &data)
		if err != nil {
			panic(err)
		}

		render.HTML(w, r, `
		<html>
		<head>
		<style>table, th, td, { border: 1.5px solid black; padding: 5px; text-align: center; }</style>
		</head>
		<body style="background-color:#242424; font-family: helvetica, sans-serif;">
		<table style="padding: 20px; text-align: center;">
		<tr style="color:#e8e9eb; font-size: 18px; font-weight: bold;">
			<th>ID</th>
			<th>Name</th>
			<th>Text</th>
			<th>Type</th>
			<th>Attack</th>
			<th>Health</th>
			<th>Cost</th>
			<th>Card Class</th>
			<th>Rarity</th>
		`)

		for _, card := range data {

			xColor := utils.ClassColorPicker(card.CardClass)
			yColor := utils.RarityColorPicker(card.Rarity)

			render.HTML(w, r, fmt.Sprintf(`
			<tr style="font-size: 14px;">
			<style>
			a:link {
				color: #e8e9eb;
			}
			a:visited {
				color: #7d7d7d
			}
			a:hover {
				color: yellow
			}
			</style>
				<td style="color:#e8e9eb">%s</td> 
				<td style="color:#e8e9eb"><a href= "https://art.hearthstonejson.com/v1/render/latest/enUS/512x/%s.png">%s</a></td>
				<td style="color:#e8e9eb; text-align: left;">%s</td>
				<td style="color:#e8e9eb">%s</td>
				<td style="color:#ffcc05; text-shadow:
				-1px -1px 0 #4a3c07,
				1px -1px 0 #4a3c07,
				-1px 1px 0 #4a3c07,
				1px 1px 0 #4a3c07;"><b>%d</b></td>
				<td style="color:#ff0505; text-shadow:
				-1px -1px 0 #4f0202,
				1px -1px 0 #4f0202,
				-1px 1px 0 #4f0202,
				1px 1px 0 #4f0202;"><b>%d</b></td>
				<td style="color:#4940f5; text-shadow:
				-1px -1px 0 #050330,
				1px -1px 0 #050330,
				-1px 1px 0 #050330,
				1px 1px 0 #050330;"><b>%d</b></td>
				<td style="color:%s">%s</td>
				<td style="color:%s">%s</td>`, card.Id, card.Id, card.Name, card.Text, card.CardType, card.Attack, card.Health, card.Cost, xColor, card.CardClass, yColor, card.Rarity))
		}
	})

	log.Printf("Starting server on port %s.", port)
	err2 := http.ListenAndServe(":"+port, router)
	log.Fatal(err2)

}
