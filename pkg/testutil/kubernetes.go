// Copyright 2017 uSwitch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package testutil

import (
	"fmt"
	"time"

	"github.com/uswitch/kiam/pkg/k8s"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	PhaseRunning   = "Running"
	PhaseSucceeded = "Succeeded"
)

func NewNamespace(name, roleRegexp string) *v1.Namespace {
	n := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	n.Annotations = map[string]string{k8s.AnnotationPermittedKey: roleRegexp}
	return n
}

func NewPod(namespace, name, ip, phase string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:         namespace,
			Name:              name,
			ResourceVersion:   fmt.Sprintf("%d", time.Now().UnixNano()),
			CreationTimestamp: metav1.Now(),
		},
		Status: v1.PodStatus{
			PodIP: ip,
			Phase: v1.PodPhase(phase),
		},
		Spec: v1.PodSpec{
			HostNetwork: false,
		},
	}
}

func NewPodWithRole(namespace, name, ip, phase, role string) *v1.Pod {
	pod := NewPod(namespace, name, ip, phase)
	pod.ObjectMeta.Annotations = map[string]string{k8s.AnnotationIAMRoleKey: role}
	return pod
}
