package network

import (
	"github.com/04Akaps/metting/module/API/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type like struct {
	n *Network
}

func likeRouter(n *Network) {
	l := &like{n: n}

	// TODO 기회가 된다면, gRPC를 통한 토큰 인증 도입
	// 토큰인 Paseto, JWT 원하는 방식으로 관리
	n.Router(POST, "/like-user", l.likeSomeOne)
	n.Router(POST, "/accept-like", l.acceptLikeRequest)
	n.Router(POST, "/refuse-like", l.refuseLikeRequest)
	n.Router(GET, "/like-list", l.getLikeList)
}

func (l *like) likeSomeOne(c *gin.Context) {
	var req types.LikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if err = l.n.service.LikeSomeOne(req.FromUser, req.ToUser); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
	} else {
		res(c, http.StatusOK, "Success")
	}
}

func (l *like) acceptLikeRequest(c *gin.Context) {
	var req types.AcceptedRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if err = l.n.service.AcceptedLike(req.FromUser, req.ToUser); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
	} else {
		res(c, http.StatusOK, "Success")
	}
}

func (l *like) refuseLikeRequest(c *gin.Context) {
	var req types.RefuseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if err = l.n.service.RefuseLike(req.FromUser, req.ToUser); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
	} else {
		res(c, http.StatusOK, "Success")
	}
}

func (l *like) getLikeList(c *gin.Context) {
	var req types.GetLikedListRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else if result, err := l.n.service.GetLikedList(req.User); err != nil {
		res(c, http.StatusInternalServerError, err.Error())
	} else {
		res(c, http.StatusOK, result)
	}
}
