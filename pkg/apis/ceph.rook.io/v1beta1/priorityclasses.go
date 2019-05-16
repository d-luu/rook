/*
Copyright 2018 The Rook Authors. All rights reserved.

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
package v1beta1

import (
	rook "github.com/rook/rook/pkg/apis/rook.io/v1alpha2"
)

const (
	PriorityClassesKeyMgr       = "mgr"
	PriorityClassesKeyMon       = "mon"
	PriorityClassesKeyOSD       = "osd"
	PriorityClassesKeyRBDMirror = "rbdmirror"
)

// GetMgrPriorityClass returns the priority class name for the MGR service
func GetMgrPriorityClass(p rook.PriorityClassSpec) string {
	return p[PriorityClassesKeyMgr]
}

// GetMonPriorityClass returns the priority class name for the monitors
func GetMonPriorityClass(p rook.PriorityClassSpec) string {
	return p[PriorityClassesKeyMon]
}

// GetOSDPriorityClass returns the priority class name for the OSDs
func GetOSDPriorityClass(p rook.PriorityClassSpec) string {
	return p[PriorityClassesKeyOSD]
}

// GetRBDMirrorPriorityClass returns the priority class name for the RBD Mirrors
func GetRBDMirrorPriorityClass(p rook.PriorityClassSpec) string {
	return p[PriorityClassesKeyRBDMirror]
}
