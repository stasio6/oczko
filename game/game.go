package game

import (
	"fmt"
	"os"
	"strconv"
)

// Player architecture

type Player struct {
	cards []Card
	passed bool
	name string
	id int
}

func newPlayer(id int, name string) Player {
	return Player{[]Card{}, false, name, id}
}

func (player Player) countPoints() int {
	points := 0
	for _, card := range player.cards {
		cardPoints := card.number
		if cardPoints > 10 {
			cardPoints = 10
		}
		points += cardPoints
	}
	return points
}

func (player Player) showHandToYou() string {
	result := ""
	for _, card := range player.cards {
		result += card.printCard() + " "
	}
	return result
}

func (player Player) showHandToOthers() string {
	result := ""
	for i, card := range player.cards {
		if i == 0 {
			result += card.printCard() + " "
		} else {
			result += "? ";
		}
	}
	return result
}

// Game architecture

type Game struct {
	deck             Deck
	player1, player2 Player
	player1Turn      bool
}

func NewGame() Game {
	game := Game{newDeck(), newPlayer(1, "Player 1"), newPlayer(2, "Player 2"), true}
	drawCard(&game.player1, &game.deck)
	drawCard(&game.player2, &game.deck)
	drawCard(&game.player1, &game.deck)
	drawCard(&game.player2, &game.deck)
	return game
}

func drawCard(player *Player, deck *Deck) {
	player.cards = append(player.cards, deck.draw())
}

func (game *Game) belongsToGame(playerId int) bool {
	return playerId == game.player1.id || playerId == game.player2.id
}

func (game *Game) isYourTurn(playerId int) bool {
	return (game.player1Turn && playerId == game.player1.id) ||
		(!game.player1Turn && playerId == game.player2.id)
}

func (game *Game) nextTurn() {
	game.player1Turn = !game.player1Turn
}

// status types
// 0 - your turn
// 1 - opponent's turn
// 2 - game ended
// 3 - error
func (game *Game) GetStatus(playerId int) (string, int) {
	if !game.belongsToGame(playerId) {
		panic("ERRORRO")
		return "Player doesn't belong to this game\n", 3
	}
	if game.player1.passed && game.player2.passed {
		return game.endGame(), 2
	}
	if !game.isYourTurn(playerId) {
		// send wait response
		return "Waiting for opponent's turn", 1
	}
	var playerMoving *Player
	var secondPlayer *Player
	response := ""
	if game.player1Turn {
		playerMoving = &game.player1
		secondPlayer = &game.player2
		response += game.player1.name + " turn:\n"
	} else {
		playerMoving = &game.player2
		secondPlayer = &game.player1
		response += game.player2.name + " turn:\n"
	}
	if playerMoving.passed {
		response += "You already passed!\n"
		game.nextTurn()
		return response, 1
	}
	response += "Your hand: "
	response += playerMoving.showHandToYou()
	response += "\nOpposite hand: "
	response += secondPlayer.showHandToOthers()
	response += "\n"
	return response, 0
}

func (game *Game) MakeMove(playerId int, move string) string {
	if !game.belongsToGame(playerId) {
		panic("ERRORRO")
		return "ERRORRO"
	}
	if !game.isYourTurn(playerId) {
		panic("Not your turn")
		return "NOT YOUR TURN"
	}
	playerMoving := &game.player1
	if playerId == game.player2.id {
		playerMoving = &game.player2
	}
	response := ""
	if move == "draw" {
		drawCard(playerMoving, &game.deck)
		response = "You drawn a card\n"
		if playerMoving.countPoints() > 21 {
			response += "Fail! You have " + strconv.Itoa(playerMoving.countPoints()) + " points!\n"
			response += "From now you are forced to pass\n"
			playerMoving.passed = true
		}
		game.nextTurn()
	} else if move == "pass" {
		response = "You passed"
		playerMoving.passed = true
		game.nextTurn()
	} else {
		panic("Bad response")
	}
	return response
}

func (game *Game) makeLocalMove() {
	if game.player1.passed && game.player2.passed {
		fmt.Print(game.endGame())
		os.Exit(0)
	}
	var playerMoving *Player
	var secondPlayer *Player
	if game.player1Turn {
		game.player1Turn = false
		playerMoving = &game.player1
		secondPlayer = &game.player2
		fmt.Println(game.player1.name, " turn:")
	} else {
		game.player1Turn = true
		playerMoving = &game.player2
		secondPlayer = &game.player1
		fmt.Println(game.player2.name, " turn:")
	}
	if playerMoving.passed {
		fmt.Print("You already passed!\n")
		return
	}
	fmt.Print("Your hand: ")
	fmt.Println(playerMoving.showHandToYou())
	fmt.Print("\nOpposite hand: ")
	fmt.Println(secondPlayer.showHandToOthers())
	fmt.Println("\nType 'draw' to draw or 'pass' to pass:")
	var response string
	for {
		fmt.Scanln(&response)
		if response == "draw" {
			drawCard(playerMoving, &game.deck)
			fmt.Println("You drawn a card")
			if playerMoving.countPoints() > 21 {
				fmt.Println("Fail! You have ", playerMoving.countPoints(), " points!")
				playerMoving.passed = true
			}
			break
		} else if response == "pass" {
			fmt.Println("You passed")
			playerMoving.passed = true
			break
		} else {
			fmt.Println("Bad response. Type again!")
		}
	}
}

func (game *Game) endGame() string {
	endStatus := ""
	endStatus += "Game has ended!\n"
	p1Points := game.player1.countPoints()
	p2Points := game.player2.countPoints()
	endStatus += "Player 1 has " + strconv.Itoa(p1Points) + " points!\n"
	endStatus += "Player 2 has " + strconv.Itoa(p2Points) + " points!\n"
	if p1Points > 21 && p2Points > 21 {
		endStatus += "You both failed and lost!\n"
	} else if p1Points > 21 {
		endStatus += "Player 1 failed, so Player 2 won!\n"
	} else if p2Points > 21 {
		endStatus += "Player 2 failed, so Player 1 won!\n"
	} else if p1Points > p2Points {
		endStatus += "Player 1 is closer to 21, he won!\n"
	} else if p2Points > p1Points {
		endStatus += "Player 2 is closer to 21, he won!\n"
	} else {
		endStatus += "It's a draw!\n"
	}
	return endStatus
}

