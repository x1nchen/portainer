package portainer

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/eleztian/portainer/model"
	"testing"
)

func TestStacksApiService_StackList(t *testing.T) {
	client := NewAPIClient(&Configuration{
		BasePath:      "http://192.168.120.71:9010/api",
		DefaultHeader: map[string]string{},
		UserAgent:     "iast/sdl_config",
	})
	rsp, _, err := client.AuthApi.AuthenticateUser(context.TODO(), model.AuthenticateUserRequest{
		Username: "admin",
		Password: "moresec2019",
	})
	if err != nil {
		t.Error(err)
	}
	sl, _, err := client.StacksApi.StackList(ContextWithAPIKey(context.TODO(), rsp.Jwt), &StackListOpts{Filters: optional.NewString(`{"SwarmID":"jq8yqjtr1pkqkfxhsl9racm00"}`)})

	fmt.Println(sl[0].ResourceControl)
}

func TestStacksApiService_StackInspect(t *testing.T) {
	client := NewAPIClient(&Configuration{
		BasePath:      "http://192.168.120.71:9010/api",
		DefaultHeader: map[string]string{},
		UserAgent:     "iast/sdl_config",
	})
	rsp, _, err := client.AuthApi.AuthenticateUser(context.TODO(), model.AuthenticateUserRequest{
		Username: "admin",
		Password: "moresec2019",
	})
	if err != nil {
		t.Error(err)
	}
	sl, _, err := client.StacksApi.StackInspect(ContextWithAPIKey(context.TODO(), rsp.Jwt), 4)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sl.SwarmID)
}

func TestStacksApiService_StackCreate(t *testing.T) {
	client := NewAPIClient(&Configuration{
		BasePath:      "http://192.168.120.71:9010/api",
		DefaultHeader: map[string]string{},
		UserAgent:     "iast/sdl_config",
	})
	rsp, _, err := client.AuthApi.AuthenticateUser(context.TODO(), model.AuthenticateUserRequest{
		Username: "admin",
		Password: "moresec2019",
	})
	if err != nil {
		t.Error(err)
	}
	sl, _, err := client.StacksApi.StackCreate(ContextWithAPIKey(context.TODO(), rsp.Jwt),
		1, "string", 2, StackCreateRequest{
			SwarmID:          "jq8yqjtr1pkqkfxhsl9racm00",
			StackFileContent: `version: "3.7"`,
			Name:             "test2",
			Env:              []model.StackEnv{},
		})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sl)
}

func TestStacksApiService_StackUpdate(t *testing.T) {
	client := NewAPIClient(&Configuration{
		BasePath:      "http://192.168.120.71:9010/api",
		DefaultHeader: map[string]string{},
		UserAgent:     "iast/sdl_config",
	})
	rsp, _, err := client.AuthApi.AuthenticateUser(context.TODO(), model.AuthenticateUserRequest{
		Username: "admin",
		Password: "moresec2019",
	})
	if err != nil {
		t.Error(err)
	}
	sl, _, err := client.StacksApi.StackUpdate(ContextWithAPIKey(context.TODO(), rsp.Jwt),
		4, model.StackUpdateRequest{
			StackFileContent: `version: "3.7"`,
			Env:              []model.StackEnv{},
		}, &StackUpdateOpts{EndpointId: optional.NewInt32(2)})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sl)
}
