package ggit

import (
	"strings"
)

// PrevHash Hash Name Email 1455122575 +0900	commit: add makefile
type Log struct {
	PrevHash string
	Hash     string
	Name     string
	Email    string
	At       string // UNIX Time
	TimeZone string
	Action   string // commit
	IsInitial  bool
	Message  string
}

func Parse(line string) *Log {
	tmp := strings.Split(line, "\t")
	idx := strings.Index(tmp[1], ": ")
	res := Log{}
	if strings.Contains(tmp[1][0:idx], "commit") {
		res.Action = "commit"
		if strings.Contains(tmp[1][0:idx], "(initial)") {
			res.IsInitial = true
		}
	}
	res.Message = tmp[1][(idx + 2):]
	info := strings.Split(tmp[0], " ")
	res.PrevHash = info[0]
	res.Hash = info[1]
	res.Name = info[2]
	res.Email = info[3][1 : len(info[3])-1]
	res.At = info[4]
	res.TimeZone = info[5]
	return &res
}
