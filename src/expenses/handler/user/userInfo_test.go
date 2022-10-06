package user

import (

	// "expenses/route"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestGetUserInfo(t *testing.T) {
	expected := `{"id":1,"name":"sample"}`
	r := SetupRouter()
	//将项目中的API注册到测试使用的router
	r.GET("/user", GetUserInfo())
	//向注册的路有发起请求
	req, _ := http.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	//模拟http服务处理请求
	r.ServeHTTP(w, req)
	// var recipes []Recipe
	// fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), expected)
	// assert.Equal(t, 492, len(recipes))

	// json.Unmarshal([]byte(w.Body.String()), &recipes)

	// expected := `{"message":"hello world"}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }

	// fmt.Println(1111)
	// mockUserResp := `{"message":"hello world"}`
	// // テスト用のサーバーを立てる
	// ts := httptest.NewServer(SetupServer())
	// defer ts.Close()
	// // リクエストを送れるか?
	// resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))
	// if err != nil {
	// 	t.Fatalf("Expected no error, got %v", err)
	// }
	// defer resp.Body.Close()
	// // Statusコードは200か?
	// fmt.Println(resp.StatusCode)
	// if resp.StatusCode != http.StatusOK {
	// 	t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	// }
	// responseData, _ := ioutil.ReadAll(resp.Body)
	// if string(responseData) != mockUserResp {
	// 	t.Fatalf("Expected hello world message, got %v", responseData)
	// }
}
