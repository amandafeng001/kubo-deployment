// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/cmd"
)

type FakeLoginStrategy struct {
	TryStub        func() error
	tryMutex       sync.RWMutex
	tryArgsForCall []struct{}
	tryReturns     struct {
		result1 error
	}
	tryReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLoginStrategy) Try() error {
	fake.tryMutex.Lock()
	ret, specificReturn := fake.tryReturnsOnCall[len(fake.tryArgsForCall)]
	fake.tryArgsForCall = append(fake.tryArgsForCall, struct{}{})
	fake.recordInvocation("Try", []interface{}{})
	fake.tryMutex.Unlock()
	if fake.TryStub != nil {
		return fake.TryStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tryReturns.result1
}

func (fake *FakeLoginStrategy) TryCallCount() int {
	fake.tryMutex.RLock()
	defer fake.tryMutex.RUnlock()
	return len(fake.tryArgsForCall)
}

func (fake *FakeLoginStrategy) TryReturns(result1 error) {
	fake.TryStub = nil
	fake.tryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLoginStrategy) TryReturnsOnCall(i int, result1 error) {
	fake.TryStub = nil
	if fake.tryReturnsOnCall == nil {
		fake.tryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.tryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLoginStrategy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tryMutex.RLock()
	defer fake.tryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLoginStrategy) recordInvocation(key string, args []interface{}) {
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

var _ cmd.LoginStrategy = new(FakeLoginStrategy)
