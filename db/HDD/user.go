package HDD

import (
	"context"
	"encoding/json"
	. "github.com/04Akaps/metting/db/HDD/query"
	"github.com/04Akaps/metting/db/HDD/types"
	"strconv"
)

func (d *DB) RegisterUser(user, description string, hobby []string, latitude, hardness float64) error {
	// 유저 등록, session 만들어서 생성
	ctx := context.Background()

	if tx, err := d.db.BeginTx(ctx, nil); err != nil {
		return err
	} else {

		if json, err := json.Marshal(hobby); err != nil {
			return err
		} else {
			if result, err := tx.Exec(InsertIgnoreUser, user, description, json); err != nil {
				tx.Rollback()
				return err
			} else {
				count, _ := result.RowsAffected()
				d.log.InfoLog("Success To Insert user", "affected", strconv.Itoa(int(count)))
			}

			if result, err := tx.Exec(InsertIgnoreUserLocation, user, latitude, hardness, latitude, hardness); err != nil {
				tx.Rollback()
				return err
			} else {
				count, _ := result.RowsAffected()
				d.log.InfoLog("Success To Insert user_location", "affected", strconv.Itoa(int(count)))
			}

			tx.Commit()
		}

	}

	return nil
}

func (d *DB) GetUser(user string) (*types.User, error) {
	var res types.User

	// ORM이 아닌 dataBase/SQL 패키지에서는 내부 Row의 Scan 리시버가 많은 타입을 지원하지 않는다.
	// 그래서 JSON 값에 대해서 가져오려면 interface로 받아줘야 한다.
	var image interface{}
	var hobby interface{}

	if err := d.db.QueryRow(GetUserByName, user).Scan(
		&res.UserName,
		&image,
		&res.Description,
		&hobby,
		&res.IsValid,
		&res.Latitude,
		&res.Hardness,
		&res.Location,
	); err != nil {
		return nil, err
	} else if err = unMarshalToField(
		[]interface{}{image, hobby},
		&res.Image, &res.Hobby,
	); err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

func (d *DB) UpdateUserImage(user string, imageLink string) error {
	// S3 업로드 진행 후 실행

	if result, err := d.db.Exec(UpdateUserImage, imageLink, user, imageLink); err != nil {
		return err
	} else {
		count, _ := result.RowsAffected()
		d.log.InfoLog("Success To Update image", "affected", strconv.Itoa(int(count)))
		return nil
	}

}

func (d *DB) GetAroundFriends(userName string, latitude, hardness float64, searchRange, limit int64) ([]*types.AroundUser, error) {

	if rows, err := d.db.Query(GetAroundFriends, userName, hardness, latitude, searchRange, hardness, latitude, limit); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var result []*types.AroundUser

		for rows.Next() {
			var res types.AroundUser
			var image interface{}

			if err = rows.Scan(
				&res.UserName,
				&image,
				&res.Latitude,
				&res.Hardness,
			); err != nil {
				return nil, err
			} else if err = unMarshalToField([]interface{}{image}, &res.Image); err != nil {
				return nil, err
			} else {
				result = append(result, &res)
			}
		}

		return result, nil
	}

}
