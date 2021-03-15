package session

func Pop(sess string) error {
	tx,err := _db.Begin(true)
	if err != nil {return err}

	_,err = tx.Delete(sess)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}