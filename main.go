package main

import (
	"encoding/json"
	"log"
	"net/http"

	sr "igor.com.br/lol/lib/summonersRift"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")

	players := sr.NewSummonerRift().Login("BANOFFE").Queue().FindMatch().GetPlayers()
	json.NewEncoder(w).Encode(players)
}
func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9595", nil))
}
