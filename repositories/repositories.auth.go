package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"fm.auth/config"
	"fm.auth/entities"
	"fmt"
	"net/http"
)

func RepoLogin(email string, password string) (error, string) {
	configs := config.GetConfig()
	loginReq := entities.Login{Email: email, Password: password}

	// Chuyển đối tượng User thành JSON
	jsonBody, err := json.MarshalIndent(loginReq, "", " ")
	if err != nil {
		fmt.Println("Lỗi khi chuyển đối tượng sang JSON: ", err)
		return err, ""
	}

	fmt.Println(configs)
	resp, err := http.Post(configs.UrlLoginToFM, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Lỗi khi gọi API: ", err)
		return err, ""
	} else if resp.StatusCode != 200 {
		fmt.Println(resp.Body)
		err := errors.New("Tài khoản hoặc mật khẩu không đúng")
		return err, ""
	}
	defer resp.Body.Close()
	// Đọc nội dung phản hồi và giải mã JSON thành một struct Post
	var response entities.ResponseLoginFM
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Lỗi khi giải mã JSON: ", err)
		return err, ""
	} else if response.Ok != true {
		return err, ""
	}

	return nil, response.Data.Token
}
