package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "king", "queen", "jack":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	// Check if both cards are "ace"
	if card1 == "ace" && card2 == "ace" {
		return "P"   // Return "P" (for "Split")
	}
	// Check if player has a score of 21 and dealer doesn't have a blackjack (10 or 11)
	if playerScore == 21 && (dealerScore != 10 && dealerScore != 11) {
		return "W"   // Return "W" (for "Win")
	}
	// Check if player score is between 17 and 21 (inclusive)
	if playerScore >= 17 && playerScore <= 21 {
		return "S"   // Return "S" (for "Stand")
	}
	// Check if player score is between 12 and 16 (inclusive)
	if playerScore >= 12 && playerScore <= 16 {
		// Check if dealer's score is less than 7
		if dealerScore < 7 {
			return "S"   // Return "S" (for "Stand")
		}
	}
	return "H"   // Return "H" (for "Hit") for all other cases
}
