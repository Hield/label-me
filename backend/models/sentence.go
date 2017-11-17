package models

import (
	"log"

	"github.com/Hield/label-me/backend/global"
)

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
	log.Println("Content")
	log.Println(content)
	log.Println("Label")
	log.Println(label)
	if hasLabel {
		_, err = stmt.Exec(content, label)
	} else {
		_, err = stmt.Exec(content, nil)
	}
	log.Println("Error")
	log.Println(err)
	return
}
