/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by counterfeiter. DO NOT EDIT.
package anagofakes

import (
	"sync"

	"k8s.io/release/pkg/build"
	"k8s.io/release/pkg/release"
)

type FakeReleaseImpl struct {
	CheckReleaseBucketStub        func(*build.Options) error
	checkReleaseBucketMutex       sync.RWMutex
	checkReleaseBucketArgsForCall []struct {
		arg1 *build.Options
	}
	checkReleaseBucketReturns struct {
		result1 error
	}
	checkReleaseBucketReturnsOnCall map[int]struct {
		result1 error
	}
	CopyStagedFromGCSStub        func(*build.Options, string, string) error
	copyStagedFromGCSMutex       sync.RWMutex
	copyStagedFromGCSArgsForCall []struct {
		arg1 *build.Options
		arg2 string
		arg3 string
	}
	copyStagedFromGCSReturns struct {
		result1 error
	}
	copyStagedFromGCSReturnsOnCall map[int]struct {
		result1 error
	}
	GenerateReleaseVersionStub        func(string, string, string, bool) (*release.Versions, error)
	generateReleaseVersionMutex       sync.RWMutex
	generateReleaseVersionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 bool
	}
	generateReleaseVersionReturns struct {
		result1 *release.Versions
		result2 error
	}
	generateReleaseVersionReturnsOnCall map[int]struct {
		result1 *release.Versions
		result2 error
	}
	PrepareWorkspaceReleaseStub        func(string, string) error
	prepareWorkspaceReleaseMutex       sync.RWMutex
	prepareWorkspaceReleaseArgsForCall []struct {
		arg1 string
		arg2 string
	}
	prepareWorkspaceReleaseReturns struct {
		result1 error
	}
	prepareWorkspaceReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	PublishVersionStub        func(string, string, string, string, string, []string, bool, bool) error
	publishVersionMutex       sync.RWMutex
	publishVersionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 []string
		arg7 bool
		arg8 bool
	}
	publishVersionReturns struct {
		result1 error
	}
	publishVersionReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateImagesStub        func(string, string, string) error
	validateImagesMutex       sync.RWMutex
	validateImagesArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	validateImagesReturns struct {
		result1 error
	}
	validateImagesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseImpl) CheckReleaseBucket(arg1 *build.Options) error {
	fake.checkReleaseBucketMutex.Lock()
	ret, specificReturn := fake.checkReleaseBucketReturnsOnCall[len(fake.checkReleaseBucketArgsForCall)]
	fake.checkReleaseBucketArgsForCall = append(fake.checkReleaseBucketArgsForCall, struct {
		arg1 *build.Options
	}{arg1})
	stub := fake.CheckReleaseBucketStub
	fakeReturns := fake.checkReleaseBucketReturns
	fake.recordInvocation("CheckReleaseBucket", []interface{}{arg1})
	fake.checkReleaseBucketMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) CheckReleaseBucketCallCount() int {
	fake.checkReleaseBucketMutex.RLock()
	defer fake.checkReleaseBucketMutex.RUnlock()
	return len(fake.checkReleaseBucketArgsForCall)
}

func (fake *FakeReleaseImpl) CheckReleaseBucketCalls(stub func(*build.Options) error) {
	fake.checkReleaseBucketMutex.Lock()
	defer fake.checkReleaseBucketMutex.Unlock()
	fake.CheckReleaseBucketStub = stub
}

func (fake *FakeReleaseImpl) CheckReleaseBucketArgsForCall(i int) *build.Options {
	fake.checkReleaseBucketMutex.RLock()
	defer fake.checkReleaseBucketMutex.RUnlock()
	argsForCall := fake.checkReleaseBucketArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseImpl) CheckReleaseBucketReturns(result1 error) {
	fake.checkReleaseBucketMutex.Lock()
	defer fake.checkReleaseBucketMutex.Unlock()
	fake.CheckReleaseBucketStub = nil
	fake.checkReleaseBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CheckReleaseBucketReturnsOnCall(i int, result1 error) {
	fake.checkReleaseBucketMutex.Lock()
	defer fake.checkReleaseBucketMutex.Unlock()
	fake.CheckReleaseBucketStub = nil
	if fake.checkReleaseBucketReturnsOnCall == nil {
		fake.checkReleaseBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkReleaseBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CopyStagedFromGCS(arg1 *build.Options, arg2 string, arg3 string) error {
	fake.copyStagedFromGCSMutex.Lock()
	ret, specificReturn := fake.copyStagedFromGCSReturnsOnCall[len(fake.copyStagedFromGCSArgsForCall)]
	fake.copyStagedFromGCSArgsForCall = append(fake.copyStagedFromGCSArgsForCall, struct {
		arg1 *build.Options
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CopyStagedFromGCSStub
	fakeReturns := fake.copyStagedFromGCSReturns
	fake.recordInvocation("CopyStagedFromGCS", []interface{}{arg1, arg2, arg3})
	fake.copyStagedFromGCSMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSCallCount() int {
	fake.copyStagedFromGCSMutex.RLock()
	defer fake.copyStagedFromGCSMutex.RUnlock()
	return len(fake.copyStagedFromGCSArgsForCall)
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSCalls(stub func(*build.Options, string, string) error) {
	fake.copyStagedFromGCSMutex.Lock()
	defer fake.copyStagedFromGCSMutex.Unlock()
	fake.CopyStagedFromGCSStub = stub
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSArgsForCall(i int) (*build.Options, string, string) {
	fake.copyStagedFromGCSMutex.RLock()
	defer fake.copyStagedFromGCSMutex.RUnlock()
	argsForCall := fake.copyStagedFromGCSArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSReturns(result1 error) {
	fake.copyStagedFromGCSMutex.Lock()
	defer fake.copyStagedFromGCSMutex.Unlock()
	fake.CopyStagedFromGCSStub = nil
	fake.copyStagedFromGCSReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSReturnsOnCall(i int, result1 error) {
	fake.copyStagedFromGCSMutex.Lock()
	defer fake.copyStagedFromGCSMutex.Unlock()
	fake.CopyStagedFromGCSStub = nil
	if fake.copyStagedFromGCSReturnsOnCall == nil {
		fake.copyStagedFromGCSReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyStagedFromGCSReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) GenerateReleaseVersion(arg1 string, arg2 string, arg3 string, arg4 bool) (*release.Versions, error) {
	fake.generateReleaseVersionMutex.Lock()
	ret, specificReturn := fake.generateReleaseVersionReturnsOnCall[len(fake.generateReleaseVersionArgsForCall)]
	fake.generateReleaseVersionArgsForCall = append(fake.generateReleaseVersionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.GenerateReleaseVersionStub
	fakeReturns := fake.generateReleaseVersionReturns
	fake.recordInvocation("GenerateReleaseVersion", []interface{}{arg1, arg2, arg3, arg4})
	fake.generateReleaseVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionCallCount() int {
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	return len(fake.generateReleaseVersionArgsForCall)
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionCalls(stub func(string, string, string, bool) (*release.Versions, error)) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = stub
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionArgsForCall(i int) (string, string, string, bool) {
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	argsForCall := fake.generateReleaseVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionReturns(result1 *release.Versions, result2 error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = nil
	fake.generateReleaseVersionReturns = struct {
		result1 *release.Versions
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionReturnsOnCall(i int, result1 *release.Versions, result2 error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = nil
	if fake.generateReleaseVersionReturnsOnCall == nil {
		fake.generateReleaseVersionReturnsOnCall = make(map[int]struct {
			result1 *release.Versions
			result2 error
		})
	}
	fake.generateReleaseVersionReturnsOnCall[i] = struct {
		result1 *release.Versions
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseImpl) PrepareWorkspaceRelease(arg1 string, arg2 string) error {
	fake.prepareWorkspaceReleaseMutex.Lock()
	ret, specificReturn := fake.prepareWorkspaceReleaseReturnsOnCall[len(fake.prepareWorkspaceReleaseArgsForCall)]
	fake.prepareWorkspaceReleaseArgsForCall = append(fake.prepareWorkspaceReleaseArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.PrepareWorkspaceReleaseStub
	fakeReturns := fake.prepareWorkspaceReleaseReturns
	fake.recordInvocation("PrepareWorkspaceRelease", []interface{}{arg1, arg2})
	fake.prepareWorkspaceReleaseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseCallCount() int {
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	return len(fake.prepareWorkspaceReleaseArgsForCall)
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseCalls(stub func(string, string) error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = stub
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseArgsForCall(i int) (string, string) {
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	argsForCall := fake.prepareWorkspaceReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseReturns(result1 error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = nil
	fake.prepareWorkspaceReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseReturnsOnCall(i int, result1 error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = nil
	if fake.prepareWorkspaceReleaseReturnsOnCall == nil {
		fake.prepareWorkspaceReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.prepareWorkspaceReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PublishVersion(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 []string, arg7 bool, arg8 bool) error {
	var arg6Copy []string
	if arg6 != nil {
		arg6Copy = make([]string, len(arg6))
		copy(arg6Copy, arg6)
	}
	fake.publishVersionMutex.Lock()
	ret, specificReturn := fake.publishVersionReturnsOnCall[len(fake.publishVersionArgsForCall)]
	fake.publishVersionArgsForCall = append(fake.publishVersionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 []string
		arg7 bool
		arg8 bool
	}{arg1, arg2, arg3, arg4, arg5, arg6Copy, arg7, arg8})
	stub := fake.PublishVersionStub
	fakeReturns := fake.publishVersionReturns
	fake.recordInvocation("PublishVersion", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6Copy, arg7, arg8})
	fake.publishVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) PublishVersionCallCount() int {
	fake.publishVersionMutex.RLock()
	defer fake.publishVersionMutex.RUnlock()
	return len(fake.publishVersionArgsForCall)
}

func (fake *FakeReleaseImpl) PublishVersionCalls(stub func(string, string, string, string, string, []string, bool, bool) error) {
	fake.publishVersionMutex.Lock()
	defer fake.publishVersionMutex.Unlock()
	fake.PublishVersionStub = stub
}

func (fake *FakeReleaseImpl) PublishVersionArgsForCall(i int) (string, string, string, string, string, []string, bool, bool) {
	fake.publishVersionMutex.RLock()
	defer fake.publishVersionMutex.RUnlock()
	argsForCall := fake.publishVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8
}

func (fake *FakeReleaseImpl) PublishVersionReturns(result1 error) {
	fake.publishVersionMutex.Lock()
	defer fake.publishVersionMutex.Unlock()
	fake.PublishVersionStub = nil
	fake.publishVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PublishVersionReturnsOnCall(i int, result1 error) {
	fake.publishVersionMutex.Lock()
	defer fake.publishVersionMutex.Unlock()
	fake.PublishVersionStub = nil
	if fake.publishVersionReturnsOnCall == nil {
		fake.publishVersionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishVersionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) ValidateImages(arg1 string, arg2 string, arg3 string) error {
	fake.validateImagesMutex.Lock()
	ret, specificReturn := fake.validateImagesReturnsOnCall[len(fake.validateImagesArgsForCall)]
	fake.validateImagesArgsForCall = append(fake.validateImagesArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ValidateImagesStub
	fakeReturns := fake.validateImagesReturns
	fake.recordInvocation("ValidateImages", []interface{}{arg1, arg2, arg3})
	fake.validateImagesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) ValidateImagesCallCount() int {
	fake.validateImagesMutex.RLock()
	defer fake.validateImagesMutex.RUnlock()
	return len(fake.validateImagesArgsForCall)
}

func (fake *FakeReleaseImpl) ValidateImagesCalls(stub func(string, string, string) error) {
	fake.validateImagesMutex.Lock()
	defer fake.validateImagesMutex.Unlock()
	fake.ValidateImagesStub = stub
}

func (fake *FakeReleaseImpl) ValidateImagesArgsForCall(i int) (string, string, string) {
	fake.validateImagesMutex.RLock()
	defer fake.validateImagesMutex.RUnlock()
	argsForCall := fake.validateImagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeReleaseImpl) ValidateImagesReturns(result1 error) {
	fake.validateImagesMutex.Lock()
	defer fake.validateImagesMutex.Unlock()
	fake.ValidateImagesStub = nil
	fake.validateImagesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) ValidateImagesReturnsOnCall(i int, result1 error) {
	fake.validateImagesMutex.Lock()
	defer fake.validateImagesMutex.Unlock()
	fake.ValidateImagesStub = nil
	if fake.validateImagesReturnsOnCall == nil {
		fake.validateImagesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateImagesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkReleaseBucketMutex.RLock()
	defer fake.checkReleaseBucketMutex.RUnlock()
	fake.copyStagedFromGCSMutex.RLock()
	defer fake.copyStagedFromGCSMutex.RUnlock()
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	fake.publishVersionMutex.RLock()
	defer fake.publishVersionMutex.RUnlock()
	fake.validateImagesMutex.RLock()
	defer fake.validateImagesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseImpl) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
