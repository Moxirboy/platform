package handler

import (
	"gateway/internal/controller/http/v1/dto"
	pb "gateway/proto"
	"time"
)


func FromReqToTeacher(
	teacher dto.Teacher,
	) pb.User {
	return pb.User{
		ID: teacher.ID,
		Username: teacher.Username,
		Password: teacher.Password,
		Role: "teacher",
		CreatedAt: time.Now().Format("2006-01-02"),
		UpdateAt: time.Now().Format("2006-01-02"),
		Verified: true,
	}
}