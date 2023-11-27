package wrapper

import (
	"fmt"
	"testing"
)

type MockedWingsFlyBehaviour struct {
	FlyFunc func()
}

func (m MockedWingsFlyBehaviour) Fly() {
	m.FlyFunc()
}

func Test_run(t *testing.T) {
	tests := []struct {
		name         string
		flyBehaviour FlyBehaviour
	}{
		{"it should fly", MockedWingsFlyBehaviour{
			FlyFunc: func() {
				fmt.Println("I fly with wings")
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(tt.flyBehaviour)
		})
	}
}
