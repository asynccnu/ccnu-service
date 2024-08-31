package service

import (
	"context"
	"fmt"
	v1 "github.com/asynccnu/ccnu-service/api/ccnu_service/v1"
	"github.com/asynccnu/ccnu-service/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
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
	err := s.uc.SaveUser(ctx, &biz.User{
		UserID:   req.User.Userid,
		Password: req.User.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SaveUserResponse{Message: "User saved successfully"}, nil
}

func (s *CCNUService) GetCookie(ctx context.Context, req *v1.GetCookieRequest) (*v1.GetCookieResponse, error) {
	user, err := s.uc.GetUserByIDFromDB(ctx, req.Userid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

	// 获取Set-Cookie头的内容
	cookieHeader := resp.Header.Get("Set-Cookie")

	// 找到JSESSIONID的开始位置
	cookieStart := strings.Index(cookieHeader, "JSESSIONID=")
	if cookieStart == -1 {
		return "", fmt.Errorf("JSESSIONID not found in cookies")
	}

	// 找到JSESSIONID值的结束位置
	cookieEnd := strings.Index(cookieHeader[cookieStart:], ";")
	if cookieEnd == -1 {
		cookieEnd = len(cookieHeader) // 如果没有;，则直接到字符串末尾
	} else {
		cookieEnd += cookieStart
	}

	// 提取JSESSIONID值
	jsessionID := cookieHeader[cookieStart:cookieEnd]

	return jsessionID, nil

}
