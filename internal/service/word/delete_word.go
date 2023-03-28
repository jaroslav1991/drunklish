package word

import (
	"drunklish/internal/model"
	"log"
)

const (
	deleteWordQuery = `delete from words where id=$1 and user_id=$2`
)

func (w *Word) DeleteWordById(id, userId int64) error {
	var word model.Word

	if err := w.db.QueryRowx(deleteWordQuery, id, userId).Scan(&word.Id); err != nil {
		return err
	}

	log.Printf("deleting word %s success", word.Word)
	return nil
}
