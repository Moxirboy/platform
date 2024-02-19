package postgres

import (
	logger "admin/pkg/log"
	"auth/internal/domain"
	"auth/internal/service/repo"
	"context"
	"database/sql"
)


type authRepository struct {
	db *sql.DB
	log logger.Logger
}

func NewAuthRepository (
	db *sql.DB,
	log logger.Logger,
) repo.AuthRepository {
	return &authRepository{
		db: db,
		log: log,
	}
}

func (r *authRepository) GetExist(
	ctx context.Context,
	username string,
	) (
	bool,
	error,
	){
	var exist int
	err:=r.db.QueryRowContext(
		ctx,
		GetExist,
		username,
	).Scan(&exist)
	if err!=nil{
		r.log.Error("repo.auth.getexist error : ",err)
		return false, err
	}
	if exist!=1{
		return false,nil
	}
	return true,nil
 }

 func (r *authRepository) GetUser(
	ctx context.Context,
	username string,
) (
	string ,
	string,
	string,
	error,
	){
		var id ,password,role string
		err:=r.db.QueryRowContext(
			ctx,
			GetUser,
			username,
		).Scan(
			&id,
			&password,
			&role,
		)
		if err!=nil{
			r.log.Error("repo.auth.GetUser error : ",err.Error())
			return "","","",err
		}
		return id ,password,role,nil
}

func (r *authRepository) CreateUser(
	ctx context.Context,
	auth *domain.Auth,
	)(
	id string,
	error error,
	){
	err:=r.db.QueryRowContext(
		ctx,
		CreateUser,
		auth.Username,
		auth.Password,
		auth.Role,
		auth.Created_at,
		auth.Updated_at,
		).Scan(
		&id,
		)
	if err!=nil{
		return "",nil
	}
	return id , nil
}