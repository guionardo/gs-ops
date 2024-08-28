package docker

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
)

type Label struct {
	Name  string
	Value string
}

func getFileMap(fileName string) (fileMap map[string]interface{}, err error) {
	var content []byte
	if content, err = os.ReadFile(fileName); err != nil {
		return
	}
	jsonDoc, err := yaml.YAMLToJSON(content)
	if err != nil {
		err = fmt.Errorf("YAML conversion error: %s", err.Error())
		return
	}

	if err = json.Unmarshal(jsonDoc, &fileMap); err != nil {
		err = fmt.Errorf("JSON unmarshaling error: %s", err.Error())
	}

	return
}

func GetLabels(fileName string) (labels map[string]string, err error) {
	var fileMap map[string]interface{}
	if fileMap, err = getFileMap(fileName); err != nil {
		return
	}
	services, ok := fileMap["services"]
	if !ok {
		err = fmt.Errorf("missing 'services' key")
		return
	}
	var servicesMap map[string]interface{}
	if servicesMap, ok = services.(map[string]interface{}); !ok {
		err = fmt.Errorf("'services' key is not a dictionary")
		return
	}
	labels = make(map[string]string)
	for _, value := range servicesMap {
		if labelsRaw, ok := value.(map[string]interface{})["labels"]; ok {
			if labelsMap, ok := labelsRaw.(map[string]interface{}); ok {
				for label, labelValue := range labelsMap {
					labels[label] = labelValue.(string)
				}
			}
		}
	}

	return
}

func SetLabels(fileName string, outputFileName string, newLabels map[string]string) (err error) {
	var labels map[string]string
	if labels, err = GetLabels(fileName); err != nil {
		return
	} else {
		for label, value := range newLabels {
			if len(value) == 0 {
				delete(labels, label)
			} else {
				labels[label] = value
			}
		}
	}

	var fileMap map[string]interface{}
	if fileMap, err = getFileMap(fileName); err != nil {
		return
	}
	newMap := make(map[string]interface{})

	for key, value := range fileMap {
		if key != "services" {
			newMap[key] = value
			continue
		}
		services := value.(map[string]interface{})
		for serviceName, service := range services {
			service.(map[string]interface{})["labels"] = labels
			if _, ok := newMap["services"]; !ok {
				newMap["services"] = make(map[string]interface{})
			}
			newMap["services"].(map[string]interface{})[serviceName] = service
		}
	}
	var newContent []byte
	if newContent, err = yaml.Marshal(newMap); err == nil {
		comment := "# Labels dynamically updated by gs-ops\n\n"
		for label, value := range newLabels {
			if len(value) == 0 {
				comment += fmt.Sprintf("#   %s [REMOVED]\n", label)
			} else {
				comment += fmt.Sprintf("#   %s = %s\n", label, value)
			}
		}
		comment += "#########\n\n"

		err = os.WriteFile(outputFileName, append([]byte(comment), newContent...), 0644)
	}
	return
}
