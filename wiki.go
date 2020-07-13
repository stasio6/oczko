package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Webpage architecture

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "{data: \"Witaj Swiecie\"}")
}

func nieBudzKojota(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ostrzezenie: \"Nie waz sie budzic groznego kojota!!!\"}")
}

func runWebServer() {
	// var number int
	// fmt.Scanf("%d", &number)
	// fmt.Print(number)
	/*runServer()
	game := newGame()
	for {
		game.makeLocalMove()
	}
	card := newCard(9, "pik")
	card.printCard()
	return*/
	http.HandleFunc("/", handler)
	http.HandleFunc("/kojot", nieBudzKojota)
	log.Fatal(http.ListenAndServe(":8080", nil))
}