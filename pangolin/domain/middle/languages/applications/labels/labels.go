package labels

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label"
)

type labels struct {
	mp   map[string]label.Label
	list []label.Label
}

func createLabels(
	mp map[string]label.Label,
	list []label.Label,
) Labels {
	out := labels{
		mp:   mp,
		list: list,
	}

	return &out
}

// All returns the labels
func (obj *labels) All() []label.Label {
	return obj.list
}

// Fetch fetches a label by name
func (obj *labels) Fetch(name string) (label.Label, error) {
	if lbl, ok := obj.mp[name]; ok {
		return lbl, nil
	}

	str := fmt.Sprintf("the label (name: %s) is not defined", name)
	return nil, errors.New(str)
}
