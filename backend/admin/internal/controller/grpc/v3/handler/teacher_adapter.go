package handler

import (
	"admin/internal/models"
	"admin/pkg/utils"
	pb "admin/proto"
	"time"
)

func FromReqToModel(
	req *pb.User,
) models.User {
	return models.User{
		ID:        req.GetID(),
		Firstname:  req.GetFirstname(),
		Lastname: req.GetLastname(),
		Password:  req.GetPassword(),
		Role:      req.GetRole(),
		CreatedAt: time.Now().Format("2006-01-02"),
		UpdateAt:  time.Now().Format("2006-01-02"),
		DeletedAt: "",
		Verified:  true,
	}
}
func FromModelToResponse(
	err error,
) *pb.ErrorResponse {
	return &pb.ErrorResponse{
		ErrorMessage: err.Error(),
	}
}

func FromReqQueryToModel(
	req *pb.PaginationQuery,
) utils.PaginationQuery {
	return utils.PaginationQuery{
		Size:    int(req.GetSize()),
		Page:    int(req.GetPage()),
		OrderBy: req.GetOrderBy(),
	}
}
func FromModelToResponseQuery(
	userList *models.UserList,
) *pb.UserList {
	return &pb.UserList{
		TotalCount: int32(userList.TotalCount),
		TotalPages: int32(userList.TotalPages),
		Page:       int32(userList.Page),
		Size:       int32(userList.Size),
		HasMore:    userList.HasMore,
		Users:      FromModelToResponseList(userList.Users),
	}
}

func FromModelToResponseList(
	users []*models.User,
) []*pb.User {
	var res []*pb.User
	for _, user := range users {
		res = append(res, &pb.User{
			ID:        user.ID,
			Firstname:  user.Firstname,
			Lastname: user.Lastname,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdateAt,
			DeletedAt: user.DeletedAt,
			Verified:  user.Verified,
		})
	}
	return res
}
