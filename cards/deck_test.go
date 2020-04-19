package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)


const filename = "_decktesting"

func TestDeck(t *testing.T) {
	t.Run("New Deck returns 52 cards", func(t *testing.T) {
		expectedCount := 52
		expectedFirstCard := "Ace of Spades"
		expectedLastCard := "King of Clubs"

		result := newDeck()

		assert.Equal(t, len(result), expectedCount)
		assert.Equal(t, result[0], expectedFirstCard)
		assert.Equal(t, result[len(result)-1], expectedLastCard)
	})

	t.Run("Save and load files", func(t *testing.T) {
		os.Remove(filename)
		cards := newDeck()
		cards.saveToFile(filename)
		result := newDeckFromFile(filename)

		assert.Equal(t, len(result), 52)
		os.Remove(filename)
	})
}
