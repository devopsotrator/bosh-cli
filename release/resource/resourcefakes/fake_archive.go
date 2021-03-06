// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/release/resource"
)

type FakeArchive struct {
	BuildStub        func(string) (string, string, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 string
	}
	buildReturns struct {
		result1 string
		result2 string
		result3 error
	}
	buildReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	FingerprintStub        func() (string, error)
	fingerprintMutex       sync.RWMutex
	fingerprintArgsForCall []struct {
	}
	fingerprintReturns struct {
		result1 string
		result2 error
	}
	fingerprintReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArchive) Build(arg1 string) (string, string, error) {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Build", []interface{}{arg1})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.buildReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeArchive) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeArchive) BuildCalls(stub func(string) (string, string, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *FakeArchive) BuildArgsForCall(i int) string {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeArchive) BuildReturns(result1 string, result2 string, result3 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeArchive) BuildReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeArchive) Fingerprint() (string, error) {
	fake.fingerprintMutex.Lock()
	ret, specificReturn := fake.fingerprintReturnsOnCall[len(fake.fingerprintArgsForCall)]
	fake.fingerprintArgsForCall = append(fake.fingerprintArgsForCall, struct {
	}{})
	fake.recordInvocation("Fingerprint", []interface{}{})
	fake.fingerprintMutex.Unlock()
	if fake.FingerprintStub != nil {
		return fake.FingerprintStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fingerprintReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeArchive) FingerprintCallCount() int {
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	return len(fake.fingerprintArgsForCall)
}

func (fake *FakeArchive) FingerprintCalls(stub func() (string, error)) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = stub
}

func (fake *FakeArchive) FingerprintReturns(result1 string, result2 error) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = nil
	fake.fingerprintReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeArchive) FingerprintReturnsOnCall(i int, result1 string, result2 error) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = nil
	if fake.fingerprintReturnsOnCall == nil {
		fake.fingerprintReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.fingerprintReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeArchive) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeArchive) recordInvocation(key string, args []interface{}) {
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

var _ resource.Archive = new(FakeArchive)
