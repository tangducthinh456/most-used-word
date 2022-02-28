package services

import (
	"most-used-word/src/model"
	"sort"
	"strings"
)

// TopUsedWord receive paramter content of text, number of top word want to return
// and return slice represent list top used word

func TopUsedWord(cont string, top int) []model.ItemTopWord {

	result := []model.ItemTopWord{}
	if top <= 0 {
		return nil
	}

	// struct for sorting in map
	type KV struct {
		Key   string
		Value int
	}

	var kvs []KV

	// split content to slice words by space
	words := strings.Fields(cont)
	topUsedMap := make(map[string]int)

	// get all words and number they occur in map
	for _, v := range words {
		topUsedMap[v]++
	}

	// convert map to slice KV to sort
	for k, v := range topUsedMap {
		kvs = append(kvs, KV{k, v})
	}

	// sort slice KV to get top words used
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})

	

	// append top words used to slice result
	for i := 0; i < top; i++ {
		result = append(result, model.ItemTopWord{Word: kvs[i].Key, NumberOccur: kvs[i].Value})
	}

	return result
}
