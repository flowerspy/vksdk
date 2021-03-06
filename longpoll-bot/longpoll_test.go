package longpoll

import (
	"os"
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/object"
)

func TestLongpoll_Shutdown(t *testing.T) {
	t.Run("Shutdown", func(t *testing.T) {
		lp := &Longpoll{}
		lp.Shutdown()
		if lp.inShutdown != 1 {
			t.Error("inShutdown != 1")
		}
	})
}

func TestLongpoll_Handler(t *testing.T) { // nolint:gocyclo
	lp := &Longpoll{}

	t.Run("FullResponse", func(t *testing.T) {
		lp.FullResponse(func(resp object.LongpollBotResponse) {})
		if len(lp.funcFullResponseList) != 1 {
			t.Error("Want len = 1")
		}
	})
}

func TestInit(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.Init(groupToken)
	badVk := api.Init("")
	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))

	type args struct {
		vk      *api.VK
		groupID int
	}

	tests := []struct {
		name string
		args args
		// wantLp  Longpoll
		wantErr bool
	}{
		{
			name: "Init error",
			args: args{
				vk:      badVk,
				groupID: 0,
			},
			wantErr: true,
		},
		{
			name: "Init good",
			args: args{
				vk:      vk,
				groupID: groupID,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Init(tt.args.vk, tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotLp, tt.wantLp) {
			// 	t.Errorf("Init() = %v, want %v", gotLp, tt.wantLp)
			// }
		})
	}
}

func TestLongpoll_checkResponse(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.Init(groupToken)
	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
	lp, _ := Init(vk, groupID)

	tests := []struct {
		name        string
		argResponse object.LongpollBotResponse
		wantErr     bool
	}{
		{
			name: "ok",
		},
		{
			name:        "failed: 1",
			argResponse: object.LongpollBotResponse{Failed: 1},
		},
		{
			name:        "failed: 2",
			argResponse: object.LongpollBotResponse{Failed: 2},
		},
		{
			name:        "failed: 3",
			argResponse: object.LongpollBotResponse{Failed: 3},
		},
		{
			name:        "failed: 4",
			argResponse: object.LongpollBotResponse{Failed: 4},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := lp.checkResponse(tt.argResponse); (err != nil) != tt.wantErr {
				t.Errorf("Longpoll.checkResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO: write TestLongpoll_Run
// func TestLongpoll_Run(t *testing.T) {
// 	groupToken := os.Getenv("GROUP_TOKEN")
// 	if groupToken == "" {
// 		t.Skip("GROUP_TOKEN empty")
// 	}
// 	vk := api.Init(groupToken)
// 	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
// 	lp, _ := Init(&vk, groupID)
// 	lp.Wait = 1

// 	t.Run("Run", func(t *testing.T) {
//
// 	})
// }

func TestLongpoll_RunError(t *testing.T) {
	vk := api.Init("")
	lp, _ := Init(vk, 0)
	lp.Wait = 1

	t.Run("Run client error", func(t *testing.T) {
		if err := lp.Run(); err == nil {
			t.Error(err)
		}
	})

	lp.Server = "http://example.com"

	t.Run("Run json error", func(t *testing.T) {
		if err := lp.Run(); err == nil {
			t.Error(err)
		}
	})
}
