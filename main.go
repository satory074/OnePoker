package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	Suit  string
	Value int
}

type Player struct {
	name     string
	is_human bool
	hand     []Card
	history  []Card
	money    int
}

func main() {
	turns := 5
	deck := shuffleDeck(createDeck())

	player1 := Player{name: "Player1", is_human: true, hand: make([]Card, 0), history: make([]Card, 0), money: 100}
	player2 := Player{name: "Player2", is_human: false, hand: make([]Card, 0), history: make([]Card, 0), money: 100}

	drawCard(&player1, &deck)
	drawCard(&player1, &deck)
	drawCard(&player2, &deck)
	drawCard(&player2, &deck)

	for i := 0; i < turns; i++ {
		fmt.Printf("--- Turn %d / %d ---\n", i+1, turns)
		printHistory(player1)
		printHistory(player2)
		printHand(player1.hand)

		playTurn(&player1, &deck)
		playTurn(&player2, &deck)

		fmt.Print("\n")
	}
}

func playTurn(player *Player, deck *[]Card) {
	drawCard(player, deck)

	// Set
	index := -1
	if player.is_human {
		input := inputValidation([]string{"0", "1"})
		index, _ = strconv.Atoi(input)
	} else {
		index = rand.Intn(2)
	}

	player.history = append(player.history, player.hand[index])
	player.hand = append(player.hand[:index], player.hand[index+1:]...)
}

func inputValidation(validOptions []string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> (%s): ", strings.Join(validOptions, ", "))

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		for _, option := range validOptions {
			if input == option {
				return input
			}
		}
	}
}

func createDeck() []Card {
	suits := []string{"S", "H", "D", "C"}

	values := make([]int, 13)
	for i := 0; i < 13; i++ {
		values[i] = i + 1
	}

	deck := make([]Card, 0, len(suits)*len(values))
	for _, suit := range suits {
		for _, value := range values {
			card := Card{Suit: suit, Value: value}
			deck = append(deck, card)
		}
	}

	return deck
}

func shuffleDeck(deck []Card) []Card {
	shuffledDeck := make([]Card, len(deck))
	copy(shuffledDeck, deck)

	rand.Seed(time.Now().UnixNano())

	for i := len(shuffledDeck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	}

	return shuffledDeck
}

func drawCard(player *Player, deck *[]Card) {
	card := (*deck)[0]
	*deck = (*deck)[1:]

	player.hand = append(player.hand, card)
}

func formatCard(card Card) string {
	suit := ""
	value := ""

	switch card.Suit {
	case "S":
		suit = "♠"
	case "H":
		suit = "♥"
	case "D":
		suit = "♦"
	case "C":
		suit = "♣"
	}

	switch card.Value {
	case 1:
		value = "A"
	case 11:
		value = "J"
	case 12:
		value = "Q"
	case 13:
		value = "K"
	default:
		value = fmt.Sprint(card.Value)
	}

	text := fmt.Sprintf("%s%s", suit, value)

	return text
}

func printHand(hand []Card) {
	for _, h := range hand {
		fmt.Printf("%s\n", formatCard(h))
	}
}

func printHistory(player Player) {
	fmt.Printf("[%s] ", player.name)

	for _, h := range player.history {
		fmt.Printf("%s ", formatCard(h))
	}

	fmt.Print("\n")
}
