package session


func Get(sess string) (userId string,err error) {
	tx,txErr := _db.Begin(false)
	if txErr != nil {return "",txErr}
	var val string
	val,err = tx.Get(sess)
	if err != nil {
		tx.Rollback()
		return "",err
	}

	tx.Commit()
	return val,nil
}