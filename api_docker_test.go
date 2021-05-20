package portainer

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/x1nchen/portainer/model"
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

func TestDockerApiService_ListContainer(t *testing.T) {
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
	if err != nil {
		t.Error(err)
	}

	res, _, err := client.DockerApi.ListContainer(context.TODO(), sl[0].Id)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res[0])
}

func TestDockerApiService_CreateImage(t *testing.T) {
	client := NewAPIClient(&Configuration{
		BasePath:      "https://betaportainer.followme-internal.com/api",
		DefaultHeader: map[string]string{},
		UserAgent:     "iast/sdl_config",
	})
	rsp, rspRaw, err := client.AuthApi.AuthenticateUser(context.TODO(), model.AuthenticateUserRequest{
		Username: "xinwei",
		Password: "Zj130802",
	})

	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp.Jwt)
	r, _ := ioutil.ReadAll(rspRaw.Body)
	fmt.Println(string(r))



	_, err = client.DockerApi.CreateImage(
		ContextWithAPIKeySecret(context.TODO(), APIKey{Prefix: "Bearer",Key: rsp.Jwt}),
		"eyJ1c2VybmFtZSI6ImdpdGxhYi1jaSIsInBhc3N3b3JkIjoieTR6WVhOVzk5Z1hoIiwiZW1haWwiOiIiLCJzZXJ2ZXJhZGRyZXNzIjoiaHR0cHM6Ly9kb2NrZXJodWIuZm9sbG93bWUtaW50ZXJuYWwuY29tIn0=",
		17,
		"dockerhub.followme-internal.com/deploy/go-ltl-grpc-srv",
		"v1.4.18",
	)

	if err != nil {
		t.Error(err)
	}
}

