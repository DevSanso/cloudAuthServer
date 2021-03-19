package session

import (
	"github.com/tidwall/buntdb"
	"time"
)
func Set(key,userId string) error {
	tx,err := _db.Begin(true)
	if err != nil {return err}
	
	_,_,err=tx.Set(key,userId,nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func SetTimeOut(key,userId string,duration time.Duration) error {
	tx,err := _db.Begin(true)
	if err != nil {return err}
	
	_,_,err=tx.Set(key,userId,&buntdb.SetOptions{true,duration})
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}