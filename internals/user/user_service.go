package user

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skradiansys/go/utils"
)

const (
	secretKey = "abkrakadabra"
)

type Service struct {
	store UserStore
	timeout time.Duration
}

func NewService(store UserStore) *Service{
	return &Service{
		store,
		time.Duration(2)*time.Second,
	}
}


func (s *Service) createUser(c context.Context, req *CreateUserReq) (*CreateUserRes,error){
	ctx,cancel := context.WithTimeout(c,s.timeout)

	defer cancel()

	hashedPassword, err:=utils.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	r,storeError := s.store.createUser(ctx,u)

	if storeError != nil {
		return nil, err
	}

	res:= &CreateUserRes{
		ID: strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email: r.Email,
	}

	return res,nil
}

type MyJwtClaims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *Service) login(c context.Context, req *LoginUserReq) (*LoginUserRes,error){
	ctx, cancel := context.WithTimeout(c,s.timeout)

	defer cancel()

	u,err := s.store.getUserByEmail(ctx,req.Email)

	if err != nil {
		return nil, err
	}

	passError := utils.CheckPassword(req.Password,u.Password)

if passError != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,MyJwtClaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
		},
	})

	ss,tokenErr := token.SignedString([]byte(secretKey))

	if tokenErr != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{accessToken: ss,ID: strconv.Itoa(int(u.ID)),Username: u.Username}, nil

}