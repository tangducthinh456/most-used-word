package unit_test

import (
	"log"
	"most-used-word/src/services"
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	log.Println("Start test!")
	exitVal := t.Run()
	log.Println("End test!")

	os.Exit(exitVal)
}

func TestWithTopMinusOne(t *testing.T) {
	// log.Printf("Test with top minus one")

	top := -1
	content := "the lazy fox jumped over the brown dog"

	topUsed := services.TopUsedWord(content, top)
	if topUsed != nil {
		t.Errorf("Expected output nil, got %+v", topUsed)
	}
}

func TestWithTopZero(t *testing.T) {
	// log.Printf("Test with top zero")
	top := 0
	content := "the window is closing"

	topUsed := services.TopUsedWord(content, top)
	if topUsed != nil {
		t.Errorf("Expected output nil, got %+v", topUsed)
	}
}

func TestWithTopOne(t *testing.T) {
	// log.Printf("Test with top one")
	top := 1
	content := "the lazy fox jumped over the brown dog"

	topUsed := services.TopUsedWord(content, top)
	if topUsed[0].Word != "the" || topUsed[0].NumberOccur != 2 {
		t.Errorf("Expected output {\"the\" : 2, got %+v", topUsed)
	}
}

func TestWithTopFour(t *testing.T) {
	// log.Printf("Test with top four")
	top := 4
	content := "four nine nine eight seven eight ten seven seven eleven seven nine eleven six three eight eleven six four four eight nine three eight seven ten seven ten ten ten two nine nine six nine ten five one ten eleven eleven eleven eleven eight ten eleven eight three two ten five six seven six nine eleven five eight eleven five eleven five nine six ten four"

	topUsed := services.TopUsedWord(content, top)
	if (topUsed[0].Word != "eleven" || topUsed[0].NumberOccur != 11) ||
		(topUsed[1].Word != "ten" || topUsed[1].NumberOccur != 10) ||
		(topUsed[2].Word != "nine" || topUsed[2].NumberOccur != 9) ||
		(topUsed[3].Word != "eight" || topUsed[3].NumberOccur != 8) {
		t.Errorf("Expected output [{\"eleven\" : 11}, {\"ten\" : 10}, {\"nine\" : 9}, {\"eight\" : 8}] , got %+v", topUsed)
	}
}

func TestWithTopFive(t *testing.T) {
	// log.Printf("Test with top five")
	top := 5
	content := "four nine nine eight seven eight ten seven seven eleven seven nine eleven six three eight eleven six four four eight nine three eight seven ten seven ten ten ten two nine nine six nine ten five one ten eleven eleven eleven eleven eight ten eleven eight three two ten five six seven six nine eleven five eight eleven five eleven five nine six ten four"

	topUsed := services.TopUsedWord(content, top)
	if (topUsed[0].Word != "eleven" || topUsed[0].NumberOccur != 11) ||
		(topUsed[1].Word != "ten" || topUsed[1].NumberOccur != 10) ||
		(topUsed[2].Word != "nine" || topUsed[2].NumberOccur != 9) ||
		(topUsed[3].Word != "eight" || topUsed[3].NumberOccur != 8) ||
		(topUsed[4].Word != "seven" || topUsed[4].NumberOccur != 7) {
		t.Errorf("Expected output [{\"eleven\" : 11}, {\"ten\" : 10}, {\"nine\" : 9}, {\"eight\" : 8}, {\"seven\" : 7}] , got %+v", topUsed)
	}
}

func TestWithTopSix(t *testing.T) {
	// log.Printf("Test with top five")
	top := 6
	content := "four nine nine eight seven eight ten seven seven eleven seven nine eleven six three eight eleven six four four eight nine three eight seven ten seven ten ten ten two nine nine six nine ten five one ten eleven eleven eleven eleven eight ten eleven eight three two ten five six seven six nine eleven five eight eleven five eleven five nine six ten four"

	topUsed := services.TopUsedWord(content, top)
	if (topUsed[0].Word != "eleven" || topUsed[0].NumberOccur != 11) ||
		(topUsed[1].Word != "ten" || topUsed[1].NumberOccur != 10) ||
		(topUsed[2].Word != "nine" || topUsed[2].NumberOccur != 9) ||
		(topUsed[3].Word != "eight" || topUsed[3].NumberOccur != 8) ||
		(topUsed[4].Word != "seven" || topUsed[4].NumberOccur != 7) ||
		(topUsed[5].Word != "six" || topUsed[5].NumberOccur != 6) {
		t.Errorf("Expected output [{\"eleven\" : 11}, {\"ten\" : 10}, {\"nine\" : 9}, {\"eight\" : 8}, {\"seven\" : 7}] , got %+v", topUsed)
	}
}

// There is more than one word occur one time
func TestWithTopOneWithMoreThanOneWord(t *testing.T) {
	top := 1
	content := "night star fail"
	topUsed := services.TopUsedWord(content, top)

	if !((topUsed[0].Word == "night" || topUsed[0].Word == "star" || topUsed[0].Word == "fail") && topUsed[0].NumberOccur == 1) {
		t.Errorf("Expected output {\"night\" : 1} or {\"star\" : 1} or {\"fail\" : 1}, got %+v", topUsed)
	}
}

// Test with top 2 with more than one word in top 2
func TestWithTopTwoWithMoreThanOneWord(t *testing.T) {
	top := 2
	content := "night star night fail"
	topUsed := services.TopUsedWord(content, top)

	if !((topUsed[0].Word == "night" && topUsed[0].NumberOccur == 2) &&
		((topUsed[1].Word == "star" || topUsed[1].Word == "fail") && topUsed[1].NumberOccur == 1)) {
		t.Errorf("Expected output {\"night\" : 2}, {\"star\" : 1} or {\"fail\" : 1}, got %+v", topUsed)
	}
}
