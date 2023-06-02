package blackjack

const (
	Stand = "S"
	Hit   = "H"
	Split = "P"
	Win   = "W"
)

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

// Check if the player's cards sum up to 21 (Blackjack)
func IsBlackJack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// Check if the player's cards are a pair of aces
func PairOfAces(card1, card2 string) bool {
	return card1 == "ace" && card2 == "ace"
}

// Using the helper functions, decide what the outcome of the First Turn is
func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2) // Calculate the total score of the player's cards
	dealerScore := ParseCard(dealerCard)             // Calculate the score of the dealer's card

	// Check if the player has a pair of aces
	if PairOfAces(card1, card2) {
		return Split // Always split a pair of aces
	} else if IsBlackJack(card1, card2) { // Check if the player has a Blackjack
		if dealerScore != 1 && dealerScore != 10 && dealerScore < 10 {
			// If the dealer does not have an ace, a figure, or a ten, the player automatically wins
			return Win
		} else {
			// If the dealer has any of those cards, the player has to stand and wait for the reveal of the other card
			return Stand
		}
	}

	switch {
	case IsBlackJack(card1, card2) && (dealerScore == 1 || dealerScore == 10):
		// If the player has a Blackjack and the dealer's card is an ace or a ten, the player stands
		return Stand
	case playerScore >= 17 && playerScore <= 20:
		// If the player's cards sum up to a value within the range [17, 20], the player should stand
		return Stand
	case playerScore >= 12 && playerScore <= 16 && dealerScore >= 7:
		// If the player's cards sum up to a value within the range [12, 16] and the dealer has a 7 or higher, the player should hit
		return Hit
	case playerScore <= 11:
		// If the player's cards sum up to 11 or lower, the player should always hit
		return Hit
	default:
		// For any other case, the player should stand
		return Stand
	}
}





/*
The switch statement is used to evaluate the player's score.

The first case playerScore < 12 represents the condition for a small hand. 
If the player's score is less than 12, the DecideSmallHand() function is called and its result is returned.

The second case playerScore <= 20 represents the condition for a medium hand. 

If the player's score is between 12 and 20 (inclusive), the DecideMediumHand(dealerScore) 
function is called with the dealer's score as an argument, and its result is returned.

If none of the above cases match, the default case is executed, representing a large hand. 
The DecideLargeHand() function is called, and its result is returned.
*/


// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
//FIRST ATTEMPT - MONOLITHIC/BLOBBY
// 	func FirstTurn(card1, card2, dealerCard string) string {
// 		playerScore := ParseCard(card1) + ParseCard(card2)
// 		dealerScore := ParseCard(dealerCard)

// 		// Check if both cards are "ace"
// 		if card1 == "ace" && card2 == "ace" {
// 			return "P"   // Return "P" (for "Split")
// 		}
// 		// Check if player has a score of 21 and dealer doesn't have a blackjack (10 or 11)
// 		if playerScore == 21 && (dealerScore != 10 && dealerScore != 11) {
// 			return "W"   // Return "W" (for "Win")
// 		}
// 		// Check if player score is between 17 and 21 (inclusive)
// 		if playerScore >= 17 && playerScore <= 21 {
// 			return "S"   // Return "S" (for "Stand")
// 		}
// 		// Check if player score is between 12 and 16 (inclusive)
// 		if playerScore >= 12 && playerScore <= 16 {
// 			// Check if dealer's score is less than 7
// 			if dealerScore < 7 {
// 				return "S"   // Return "S" (for "Stand")
// 			}
// 		}
// 		return "H"   // Return "H" (for "Hit") for all other cases
// }