/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	"github.com/golang/glog"
	"k8s.io/kops/nodeup/pkg/distros"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/nodeup/nodetasks"
)

// LogrotateBuilder install kubectl
type LogrotateBuilder struct {
	*NodeupModelContext
}

var _ fi.ModelBuilder = &LogrotateBuilder{}

func (b *LogrotateBuilder) Build(c *fi.ModelBuilderContext) error {
	if b.Distribution == distros.DistributionCoreOS {
		glog.Infof("Detected CoreOS; won't install logrotate")
		return nil
	}

	c.AddTask(&nodetasks.Package{Name: "logrotate"})

	return nil
}
