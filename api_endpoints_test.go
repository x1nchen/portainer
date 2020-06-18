package portainer

import (
	"context"
	"fmt"
	"github.com/x1nchen/portainer/model"
	"io/ioutil"
	"testing"
)

func ContextWithAPIKey(ctx context.Context, jwt string) context.Context {
	return context.WithValue(ctx, ContextAccessToken, jwt)
}
func TestEndpointsApiService_EndpointList(t *testing.T) {
	client := NewAPIClient(&Configuration{
		BasePath:      "http://192.168.120.71:9010/api",
		DefaultHeader: map[string]string{},
		UserAgent:     "iast/sdl_config",
	})
	rsp, rspRaw, err := client.AuthApi.AuthenticateUser(context.TODO(), model.AuthenticateUserRequest{
		Username: "admin",
		Password: "moresec2019",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp.Jwt)
	r, _ := ioutil.ReadAll(rspRaw.Body)
	fmt.Println(string(r))
	sl, _, err := client.EndpointsApi.EndpointList(ContextWithAPIKey(context.TODO(), rsp.Jwt))

	fmt.Println(sl[0].Name, sl[0].Id)
}
