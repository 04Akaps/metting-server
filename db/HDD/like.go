package HDD

import (
	"errors"
	. "github.com/04Akaps/metting/db/HDD/query"
	"strconv"
	"time"
)

const (
	Like   = "send"
	Accept = "accepted"
	Refuse = "refuse"
)

func (d *DB) LikeToOther(from, to string) error {

	if result, err := d.db.Exec(LikeToOther, from, to); err != nil {
		return err
	} else if changed, _ := result.RowsAffected(); changed == 0 {
		return errors.New("Affected Raw Zero")
	} else {
		return nil
	}

}

func (d *DB) RefuseRequest(from, to string) error {

	if result, err := d.db.Exec(RefuseRequest, time.Now().Unix(), from, to); err != nil {
		return err
	} else {
		count, _ := result.RowsAffected()
		d.log.InfoLog("Success To Update user_like to refuse", "affected", strconv.Itoa(int(count)))
		return nil
	}

}

func (d *DB) AcceptRequest(from, to string) error {

	if result, err := d.db.Exec(AcceptRequest, time.Now().Unix(), from, to); err != nil {
		return err
	} else {
		count, _ := result.RowsAffected()
		d.log.InfoLog("Success To Update user_like to accepted", "affected", strconv.Itoa(int(count)))
		return nil
	}

}

func (d *DB) GetLikedList(from string) {
	// return 타입을 지정
	//GetILikedList

}
