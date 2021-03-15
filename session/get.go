package session


func Get(sess string) (string,error) {
	tx,err := _db.Begin(false)
	if err != nil {return "",err}
	var val string
	val,err = tx.Get(sess)
	if err != nil {
		tx.Rollback()
		return "",err
	}

	tx.Commit()
	return val,nil
}