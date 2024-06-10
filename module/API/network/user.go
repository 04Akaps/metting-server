package network

import (
	"github.com/04Akaps/metting/module/API/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	n *Network
}

func userRouter(n *Network) {
	u := &user{n: n}

	n.Router(POST, "/register-user", u.RegisterUser)
	n.Router(POST, "/upload-image", u.UploadImage)
	n.Router(GET, "/around-friends", u.FindAroundFriends)
}

func (u *user) RegisterUser(c *gin.Context) {
	var req types.RegisterUser

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if err = u.n.service.RegisterUser(req); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
	} else {
		res(c, http.StatusOK, "Success")
	}

}

func (u *user) UploadImage(c *gin.Context) {
	name := c.Request.FormValue("userName")
	file, header, err := c.Request.FormFile("image")

	if err != nil || name == "" {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else {
		if err = u.n.service.UploadFile(name, header, file); err != nil {
			res(c, http.StatusInternalServerError, err.Error())
			return
		} else {
			res(c, http.StatusOK, "Success To Upload Image")
			return
		}
	}

}

func (u *user) FindAroundFriends(c *gin.Context) {
	var req types.FindAroundFriendsReq

	if err := c.ShouldBindQuery(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if users, err := u.n.service.FindAroundFriends(req.User, req.Range, req.Limit); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		res(c, http.StatusOK, users)
	}
}
