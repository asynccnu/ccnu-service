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
	var cookie string
	var err error
	user, err := s.uc.GetUserByIDFromDB(ctx, req.Userid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	if CheckIsUndergraduate(user.UserID) {
		cookie, err = BKSloginCCNU(user.UserID, user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to login: %v", err)
		}
	}

	return &v1.GetCookieResponse{Cookie: cookie}, nil
}

// CheckIsUndergraduate 检查该学号是否是本科生
func CheckIsUndergraduate(stuId string) bool {
	return stuId[4] == '2'
	//区分是学号第五位，本科是2，硕士是1，博士是0，工号是6或9
}

// BKSloginCCNU 模拟本科生登录CCNU并返回Cookie
func BKSloginCCNU(username, password string) (string, error) {
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
	var value string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "JSESSIONID" {
			value = cookie.Value
		}
	}
	return fmt.Sprintf("JSESSIONID=%s", value), nil
}

//func YJSloginCCNU(id, mm string) (cookie string, err error) {
//	client := &http.Client{}
//	mp := make(map[string]string)
//	str := fmt.Sprintf("csrftoken=&yhm=%s&mm=%s", id, mm)
//	var data = strings.NewReader(str)
//	timestamp := time.Now().Unix()
//	url := fmt.Sprintf("https://grd.ccnu.edu.cn/yjsxt/xtgl/login_slogin.html?time=%d", timestamp)
//	req, err := http.NewRequest("POST", url, data)
//	if err != nil {
//		return "", err
//	}
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//	for _, v := range resp.Cookies() {
//		mp[v.Name] = v.Value
//	}
//	cookie = fmt.Sprintf("JSESSIONID=%s; route=%s", mp["JSESSIONID"], mp["route"])
//	return cookie, nil
//}
