package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

type CommentResource struct {
	HotComments []*Comment `json:"hotComments"`
	Comments []*Comment `json:"comments"`
	Total int32 `json:"total"`
}

type Comment struct {
	Time int64 `json:"time"`
	Content string `json:"content"`
	BeReplied []*Reply `json:"beReplied"`
}

type Reply struct {
	Content string `json:"content"`
}

func commentGet(c *bm.Context) {
	resp := &CommentResource{
		Total: 29452,
	}

	comment := &Comment{
		Time: 1431950889987,
		Content: "十年前我爸跟我妈离婚 那年我听我妈唱了无数遍这首歌",
	}
	reply := &Reply{
		Content: "后来锦毛鼠变成了一撮白毛，可是即使是这轻飘飘的几根白毛，猪八戒也没有办法留住，最后也只能看着她从手上飘走，消散在天涯。虽然菩萨说锦毛鼠会在五百年后转世和他再见，但那时我们就已经知道，锦毛鼠渐渐消散的那个夜晚，她就再也回不来了。",
	}
	comment.BeReplied = append(comment.BeReplied, reply)

	resp.HotComments = append(resp.HotComments, comment)
	resp.Comments = append(resp.Comments, comment)

	c.JSON(resp, nil)
}