package session

func Exist(key string) (bool,error) {
	tx,err := _db.Begin(false)
	if err != nil {return false,err}
	_,err = tx.Get(key)
	if err != nil {return false,err}
	tx.Commit()
	return true,nil	
}