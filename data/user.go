package data

import (
	"context"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct{
	ID int
	Email string
	FirstName string
	LastName string
	Password string
	Active int
	IsAdmin int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan *Plan
}

func(u *User) GetAll()([]*User, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select * from users order by last_name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Active,
			&user.IsAdmin,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err!= nil {
			log.Println("Error scanning", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *User) GetUserByEmail(email string) (*User, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select * from users where email = $1`

	var user User
	row := db.QueryRowContext(ctx, query, email)
	
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Active,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetOne(id int) (*User, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select * from users where id = $1`

	var user User
	row := db.QueryRowContext(ctx, query,id)
	
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Active,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	query = `select p.id, p.plan_name, p.plan_amount, p.created_at, p.updated_at from
			 user_plans up
			 left join plans p on (p.id = up.plan_id)
			 where up.user_id = $1`
	
	var plan Plan
	row = db.QueryRowContext(ctx, query,user.ID)
	if err := row.Scan(
		&plan.ID,
		&plan.PlanName,
		&plan.PlanAmount,
		&plan.CreatedAt,
		&plan.UpdatedAt,
	); err == nil {
		user.Plan = &plan
	}

	return &user, nil
}

func(u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update users set
			 email = $1
			 first_name = $2
			 last_name = $3
			 user_active = $4
			 updated_at = $5
			 where id = $6`
	
	_, err := db.ExecContext(ctx, stmt, u.Email, u.FirstName, u.LastName, u.Active, time.Now(), u.ID)

	if err != nil {
		return err
	}

	return nil
}

func(u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := db.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		return err
	}
	return nil
}


func(u *User) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func(u *User) Insert(user User)(int, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hp, err := bcrypt.GenerateFromPassword([]byte(user.Password),12)
	if err != nil {
		return 0, err
	}

	var nID int
	stmt := `insert into users (email, first_name. last_name, password, user_active, created_at, updated_at)
			values($1,$2,$3,$4,$5,$6,$7) returning id`
	
	if err = db.QueryRowContext(ctx, stmt, user.Email, user.FirstName, user.LastName, hp, user.Active, time.Now(), time.Now()).Scan(&nID); err != nil {
		return 0, err
	}
	
	return nID, nil

}

func(u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hp, err := bcrypt.GenerateFromPassword([]byte(password),12)
	if err != nil {
		return err
	}
	 stmt := `updated users set password = $1 where id = $2`

	 _, err = db.ExecContext(ctx, stmt, hp, u.ID)
	 if err != nil {
		return err
	 }
	 return nil
}

func (u *User) PasswordMatches(pt string) (bool, error){
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pt)); err != nil {
		switch{
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	} 
	return true, nil
}

