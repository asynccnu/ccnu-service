package service

import (
	v1 "ccnu-service/api/ccnu_service/v1"
	"ccnu-service/internal/biz"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type CCNUService struct {
	v1.UnimplementedCCNUServiceServer
	uc *biz.UserUsecase
}

func NewCCNUService(uc *biz.UserUsecase) *CCNUService {
	return &CCNUService{uc: uc}
}

func (s *CCNUService) SaveUser(ctx context.Context, req *v1.SaveUserRequest) (*v1.SaveUserResponse, error) {
	err := s.uc.Save(ctx, &biz.User{
		UserID:   req.User.Userid,
		Password: req.User.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SaveUserResponse{Message: "User saved successfully"}, nil
}

func (s *CCNUService) GetCookie(ctx context.Context, req *v1.GetCookieRequest) (*v1.GetCookieResponse, error) {
	user, err := s.uc.GetByUserID(ctx, req.Userid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	cookie, err := loginCCNU(user.Username, user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	return &v1.GetCookieResponse{Cookie: cookie}, nil
}

// 模拟登录CCNU并返回Cookie
func loginCCNU(username, password string) (string, error) {
	loginURL := "https://account.ccnu.edu.cn/cas/login" // 真实的登录URL
	data := fmt.Sprintf("username=%s&password=%s", username, password)
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "CASTGC" { // 使用真实的cookie名称
			return cookie.Value, nil
		}
	}

	return "", nil
}
