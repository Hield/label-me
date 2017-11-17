package models

import (
	"errors"

	"github.com/Hield/label-me/backend/global"
)

// Sentence represents a sentence needed to label
type Sentence struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

// LabeledSentece represents a labeled sentence
type LabeledSentece struct {
	Sentence
	Count int     `json:"count"`
	Score float32 `json:"score"`
}

// AddSentence add a sentence into the sentences table
func AddSentence(content string, hasLabel bool, label int) (err error) {
	db := global.GetDB()
	stmt, err := db.Prepare(`INSERT INTO sentences (
		content,
		ftr_label
	) VALUES (
		$1,
		$2
	);`)
	if err != nil {
		return
	}
	if hasLabel {
		_, err = stmt.Exec(content, label)
	} else {
		_, err = stmt.Exec(content, nil)
	}
	return
}

// GetRandomSentences get a number of random sentences from the database
func GetRandomSentences(n int) (sentences []Sentence, err error) {
	db := global.GetDB()
	var m int // The number of sentences in the database
	err = db.QueryRow(`SELECT COUNT(*) FROM sentences`).Scan(&m)
	if err != nil {
		return
	}
	if m < n {
		err = errors.New("Not enough sentences")
		return
	}
	stmt, err := db.Prepare(`SELECT id, content FROM sentences ORDER BY random() LIMIT $1`)
	if err != nil {
		return
	}
	rows, err := stmt.Query(n)
	if err != nil {
		return
	}
	for rows.Next() {
		var sentence Sentence
		err = rows.Scan(
			&sentence.ID,
			&sentence.Content,
		)
		if err != nil {
			return
		}
		sentences = append(sentences, sentence)
	}
	return
}

// GetMostPositiveSentences return top 10 positive sentences
func GetMostPositiveSentences() (sentences []LabeledSentece, err error) {
	db := global.GetDB()
	rows, err := db.Query(`SELECT 
		id,
		content,
		count,
		score
	FROM labeled_sentences
	ORDER BY score DESC
	LIMIT 10`)
	if err != nil {
		return
	}
	for rows.Next() {
		var sentence LabeledSentece
		err = rows.Scan(
			&sentence.ID,
			&sentence.Content,
			&sentence.Count,
			&sentence.Score,
		)
		if err != nil {
			return
		}
		sentences = append(sentences, sentence)
	}
	return
}
