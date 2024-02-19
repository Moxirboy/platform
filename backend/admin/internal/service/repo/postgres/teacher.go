package postgres

import (
	"admin/internal/models"
	logger "admin/pkg/log"
	"admin/pkg/utils"
	"context"
	"database/sql"
	"log"
)

type teacherRepo struct {
	db  *sql.DB
	log logger.Logger
}

func NewTeacherRepo(db *sql.DB, log logger.Logger) *teacherRepo {
	return &teacherRepo{
		db:  db,
		log: log,
	}
}

func (r *teacherRepo) Create(ctx context.Context, teacher *models.User) error {
	err := r.db.QueryRowContext(ctx, createTeacher, &teacher.Username, &teacher.Password, &teacher.Role, &teacher.CreatedAt, &teacher.UpdateAt, &teacher.Verified).Scan(&teacher.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *teacherRepo)GetAll(
	ctx context.Context,
	PaginationQuery utils.PaginationQuery,
	) (
	*models.UserList, 
	error,
	){
	count:=0
	err:=r.db.QueryRowContext(ctx,countID).Scan(&count)
	if err!=nil{
		r.log.Error("repo.teacher.GetAll error while counting teachers: ", err.Error())
		return nil,err
	}
	if count==0 {
		return &models.UserList{
			TotalCount: count,
			TotalPages: utils.GetTotalPages(count,PaginationQuery.Size),
			Page: PaginationQuery.GetPage(),
			HasMore: utils.GetHasMore(count,PaginationQuery.GetSize(),PaginationQuery.GetPage()),
			Users: make([]*models.User,0),
		},nil
	}
	rows,err:=r.db.QueryContext(
		ctx,
		GetAll,
		PaginationQuery.GetOffSet(),
		PaginationQuery.GetLimit(),
	)
	if err!=nil{
		r.log.Error("repo.teacher.GetAll error while getting teachers: ", err.Error())
		return nil,err
	}
	defer rows.Close()
	users:=[]*models.User{}
	for rows.Next(){
		var user models.User
		rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdateAt,
			&user.Verified,
		)
		users=append(users, &user)
	}
	log.Println(PaginationQuery.GetSize())
	return &models.UserList{
		TotalCount: count,
		TotalPages: utils.GetTotalPages(count ,PaginationQuery.GetSize()),
		Page: PaginationQuery.GetPage(),
		Size: PaginationQuery.GetSize(),
		HasMore: utils.GetHasMore(PaginationQuery.GetPage(),count,PaginationQuery.GetSize()),
		Users: users,
	},nil
	}
func (r *teacherRepo) GetById(ctx context.Context, id int) (*models.User, error) {
	var teacher *models.User
	err := r.db.QueryRowContext(
		ctx,
		 GetById,
		  id,
		  ).Scan(
			 &teacher.ID,
			 &teacher.Username, 
			 &teacher.Password, 
			 &teacher.Role, 
			 &teacher.CreatedAt, 
			 &teacher.UpdateAt, 
			 &teacher.Verified,
			)
	if err != nil {
		return &models.User{}, err
	}
	return teacher, nil
}

func (r *teacherRepo) Delete(ctx context.Context, username string) error {
	tx,err:=r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
	)
	if err!=nil{
		r.log.Error("repo.teacher.update error while transaction begin: ", err.Error())
		return err
	}
	res,execErr:=tx.ExecContext(
		ctx, 
		DeleteTeacher, 
		username,
	)
	if execErr!=nil{
		r.log.Error("repo.teacher.update error while deleting teacher: ",
		 execErr.Error(),
		)
		_=tx.Rollback()
		return execErr
	}
	if count,_ :=res.RowsAffected();count==1{
		r.log.Error("repo.teacher.update error while deleting teacher: ",
		 "no rows affected",
		)
		_=tx.Rollback()
		return execErr
	}
	return nil
}