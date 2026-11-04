package sentences

import (
	"embed"
	"encoding/json"
	"math/rand/v2"

	"github.com/WeatherGod3218/mlc-project-template/internal/logging"
	"github.com/WeatherGod3218/mlc-project-template/internal/models"
)

var SentencesInitialized bool = false
var AllSentences []models.SentenceJSON

func GetAllSentences() []models.SentenceJSON {

	if !SentencesInitialized {
		logging.Logger.Warning("Attempted to grab sentences when none were loaded!")
		return make([]models.SentenceJSON, 0)
	}

	rand.Shuffle(len(AllSentences), func(i, j int) {
		AllSentences[i], AllSentences[j] = AllSentences[j], AllSentences[i]
	})

	return AllSentences
}

func LoadSentences(sentences embed.FS) error {
	rawSentences, err := sentences.ReadFile("sentences.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rawSentences, &AllSentences); err != nil {
		return err
	}

	SentencesInitialized = true
	return nil
}
