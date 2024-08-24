package docker

import (
	"reflect"
	"testing"
)

func TestGetLabels(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name       string
		args       args
		wantLabels map[string]string
		wantErr    bool
	}{
		{"docker_compose",
			args{filename: "docker-compose.yaml"},
			map[string]string{
				"de.zalando.gridRole": "this label will appear on the container",
				"gsops.app":           "This label will appear on the web service",
			},
			false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLabels, err := GetLabels(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLabels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLabels, tt.wantLabels) {
				t.Errorf("GetLabels() = %v, want %v", gotLabels, tt.wantLabels)
			}
		})
	}
}

func TestSetLabels(t *testing.T) {
	type args struct {
		fileName       string
		outputFileName string
		newLabels      map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test",
			args{
				"docker-compose.yaml",
				"new-docker-compose.yaml",
				map[string]string{
					"gsops.app":           "Updated label",
					"new.label":           "New label",
					"de.zalando.gridRole": ""},
			},
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetLabels(tt.args.fileName, tt.args.outputFileName, tt.args.newLabels); (err != nil) != tt.wantErr {
				t.Errorf("SetLabels() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
