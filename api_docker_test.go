package portainer

import (
	"context"
	"fmt"
	"github.com/eleztian/portainer/model"
	"io/ioutil"
	"testing"
)

func TestDockerApiService_DockerInfo(t *testing.T) {
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

	res, _, err := client.DockerApi.Info(context.TODO(), 2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.Swarm.NodeAddr)
	fmt.Println(sl[0].Name, sl[0].Id)
}

func TestDockerApiService_Nodes(t *testing.T) {
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

	res, _, err := client.DockerApi.Nodes(context.TODO(), sl[0].Id)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res[0].ID)
}
