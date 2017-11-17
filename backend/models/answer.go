package models

import "github.com/Hield/label-me/backend/global"

// AddAnswer add a answer into the answers table
func AddAnswer(sentenceID int, label int, feedback string, isSkipped bool) (err error) {
	db := global.GetDB()
	var dbLabel interface{}
	if !isSkipped {
		dbLabel = label
	}
	var dbFeedback interface{}
	if feedback != "" {
		dbFeedback = feedback
	}
	stmt, err := db.Prepare(`INSERT INTO answers (
		sentence_id,
		label,
		feedback,
		is_skipped
	) VALUES (
		$1,
		$2,
		$3,
		$4
	);`)
	if err != nil {
		return
	}
	_, err = stmt.Exec(sentenceID, dbLabel, dbFeedback, isSkipped)
	return
}
