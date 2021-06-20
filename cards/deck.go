package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// range(inclusiveIndex, exclusiveIndex), will return 2 new reference without modify existing one
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) // 0666 default permission: anyone can read and write this file
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File does not exist with error:", err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ",")
	return deck(stringSlice)
}

func (d deck) shuffleOutdated() {
	// this buit-in randomizer (rand.Intn) is depend on seed that Go will always have the same seed value everytime we run go file
	// so we have to create new randomizer with our unique (everytime we run go file) seed
	seed := rand.NewSource(time.Now().UnixNano()) // we will create new seed from a time at moment in int form
	r := rand.New(seed)
	for i := range d {
		newIndex := r.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		newIndex := rand.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
