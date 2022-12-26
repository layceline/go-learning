package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var val int
	switch card {
	case "ace":
		val = 11
	case "two":
		val = 2
	case "three":
		val = 3
	case "four":
		val = 4
	case "five":
		val = 5
	case "six":
		val = 6
	case "seven":
		val = 7
	case "eight":
		val = 8
	case "nine":
		val = 9
	case "ten":
		val = 10
	case "jack":
		val = 10
	case "queen":
		val = 10
	case "king":
		val = 10
	default:
		val = 0
	}
	return val
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myCards := ParseCard(card1) + ParseCard(card2)
	var action string
	switch {
	case myCards == 22:
		action = "P"
	case myCards == 21 && ParseCard(dealerCard) < 10:
		action = "W"
	case myCards == 21 && ParseCard(dealerCard) > 9:
		action = "S"
	case myCards >= 17 && myCards <= 20:
		action = "S"
	case myCards >= 12 && myCards <= 16 && ParseCard(dealerCard) >= 7:
		action = "H"
	case myCards >= 12 && myCards <= 16 && ParseCard(dealerCard) < 7:
		action = "S"
	case myCards <= 11:
		action = "H"
	}
	return action
}
