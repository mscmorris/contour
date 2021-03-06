// Copyright © 2020 VMware
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

package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Object is any Kubernetes object that has an ObjectMeta.
type Object interface {
	metav1.ObjectMetaAccessor
}

// FullName holds the name and namespace of a Kubernetes object.
type FullName struct {
	Name, Namespace string
}

// ToFullName returns the FullName of any given Kubernetes object.
func ToFullName(obj Object) FullName {
	m := obj.GetObjectMeta()
	return FullName{
		Name:      m.GetName(),
		Namespace: m.GetNamespace(),
	}
}
