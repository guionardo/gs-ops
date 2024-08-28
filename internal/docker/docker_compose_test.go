package docker

import (
	"path"
	"reflect"
	"testing"
)

func TestGetSetLabels(t *testing.T) {

	outputFileName := path.Join(t.TempDir(), "docker-compose-new.yaml")
	t.Run("set_labels", func(t *testing.T) {
		if err := SetLabels("docker-compose.yaml", outputFileName, map[string]string{
			"gsops.app":           "Updated label",
			"new.label":           "New label",
			"de.zalando.gridRole": ""},
		); err != nil {
			t.Errorf("SetLabels() error = %v", err)
		}
	})
	t.Run("check_labels", func(t *testing.T) {
		gotLabels, err := GetLabels(outputFileName)
		if err != nil {
			t.Errorf("GetLabels() error = %v", err)
			return
		}
		wantedLabels := map[string]string{
			"gsops.app": "Updated label",
			"new.label": "New label",
		}
		if !reflect.DeepEqual(gotLabels, wantedLabels) {
			t.Errorf("GetLabels() = %v, want %v", gotLabels, wantedLabels)
		}
	})
}
