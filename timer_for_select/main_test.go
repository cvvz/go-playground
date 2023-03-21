package main

import (
	"testing"
	"time"
)

func Test_waitForMount(t *testing.T) {
	type args struct {
		path     string
		intervel time.Duration
		timeout  time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := waitForMount(tt.args.path, tt.args.intervel, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("waitForMount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
