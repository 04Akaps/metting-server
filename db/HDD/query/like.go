package query

const (
	LikeToOther = "" +
		"INSERT IGNORE INTO user_like " +
		"(from_user, to_user, status) SELECT ?, ?, 'send' FROM dual " + // 가상 테이블 선언
		"WHERE EXISTS ( SELECT 2 FROM user WHERE user_name = ? OR user_name = ?);"

	RefuseRequest = "" +
		"UPDATE user_like SET status = 'refuse', updated_time = ? " +
		"WHERE from_user = ? AND to_user = ?;"

	AcceptRequest = "" +
		"UPDATE user_like SET status = 'accepted', updated_time = ? " +
		"WHERE from_user = ? AND to_user = ?;"
)

const (
	GetILikedList = "" +
		"SELECT u.user_name , u.image FROM user AS u " +
		"INNER JOIN user_like AS ul ON u.user_name = ul.from_user " +
		"WHERE ul.from_user = ?;"
)
