package handler

import (
	"gateway/internal/controller/http/v1/dto"
	"gateway/pkg/utils"
	pb "gateway/proto"
	"time"
)

func FromReqToTeacher(
	teacher dto.Teacher,
) *pb.User {
	return &pb.User{
		ID:        teacher.ID,
		Firstname:  teacher.Firstname,
		Lastname:  teacher.Lastname,
		Password:  teacher.Password,
		Role:      "teacher",
		CreatedAt: time.Now().Format("2006-01-02"),
		UpdatedAt: time.Now().Format("2006-01-02"),
		Verified:  true,
	}
}

func FromProtoToDtoUser(
	user *pb.User,
) *dto.Teacher {
	return &dto.Teacher{
		ID:        user.GetID(),
		Firstname:  user.GetFirstname(),
		Lastname: user.GetLastname(),
		Password:  user.GetPassword(),
		Role:      user.GetRole(),
		CreatedAt: user.GetCreatedAt(),
		UpdateAt:  user.GetDeletedAt(),
		Verified:  user.GetVerified(),
	}
}

func FromResQueryToQuery(
	query utils.PaginationQuery,
) *pb.PaginationQuery {
	return &pb.PaginationQuery{
		Size:    int32(query.GetSize()),
		Page:    int32(query.GetPage()),
		OrderBy: query.OrderBy,
	}
}

func FromResToProtoRes(
	res *pb.UserList,
) dto.ProtoQuery {
	user := []*dto.Teacher{}
	for _, instance := range res.Users {
		user = append(user, FromProtoToDtoUser(instance))
	}
	return dto.ProtoQuery{
		TotalCount: int(res.TotalCount),
		TotalPages: int(res.TotalPages),
		Page:       int(res.Page),
		Size:       int(res.GetSize()),
		HasMore:    res.GetHasMore(),
		Users:      user,
	}
}
