package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"blog-system/config"
	"blog-system/routes"
)

var testToken string

func init() {
	config.InitDB()
	config.InitLogger()
}

func TestRegister(t *testing.T) {
	router := routes.SetupRouter()

	registerData := map[string]interface{}{
		"username": "testuser",
		"password": "123456",
		"email":    "test@example.com",
	}
	jsonData, _ := json.Marshal(registerData)

	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("期望状态码 201，得到 %d", w.Code)
	}
}

func TestLogin(t *testing.T) {
	router := routes.SetupRouter()

	loginData := map[string]interface{}{
		"username": "testuser",
		"password": "123456",
	}
	jsonData, _ := json.Marshal(loginData)

	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if data, ok := response["data"].(map[string]interface{}); ok {
		testToken = data["token"].(string)
	}
}

func TestCreatePost(t *testing.T) {
	router := routes.SetupRouter()

	postData := map[string]interface{}{
		"title":   "测试文章",
		"content": "这是测试内容",
	}
	jsonData, _ := json.Marshal(postData)

	req, _ := http.NewRequest("POST", "/api/posts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("期望状态码 201，得到 %d", w.Code)
	}
}

func TestGetPosts(t *testing.T) {
	router := routes.SetupRouter()

	req, _ := http.NewRequest("GET", "/api/posts", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}
}
