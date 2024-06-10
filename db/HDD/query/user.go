package query

const (
	InsertIgnoreUser = "" +
		"INSERT IGNORE user (user_name, description, hobby) VALUES (?, ?, ?);"

	InsertIgnoreUserLocation = "" +
		"INSERT IGNORE user_location (user_name, latitude, hardness, location) VALUES (?, ?, ?, POINT(?, ?));"

	UpdateUserImage = "" +
		"UPDATE user SET " +
		"image = JSON_ARRAY_APPEND(image, '$', ?), " +
		"is_valid = CASE WHEN is_valid = false THEN true ELSE is_valid END " +
		"WHERE user_name = ?"
)

const (
	GetUserByName = "" +
		"SELECT " +
		"u.user_name, u.image, u.description, u.hobby, u.is_valid, " +
		"ul.latitude, ul.hardness, ul.location " +
		"FROM user AS u INNER JOIN user_location AS ul " +
		"ON u.user_name = ul.user_name " +
		"WHERE u.user_name = ?;"

	GetAroundFriends = "" +
		"SELECT u.user_name, u.image, ul.latitude, ul.hardness " +
		"FROM user AS u " +
		"INNER JOIN user_location AS ul ON u.user_name = ul.user_name " +
		"WHERE u.user_name != ? AND ST_Distance_Sphere(POINT(?, ?), POINT(hardness, latitude)) <= ? " +
		"ORDER BY ST_Distance_Sphere(POINT(?, ?), POINT(hardness, latitude)) LIMIT ?"
)
