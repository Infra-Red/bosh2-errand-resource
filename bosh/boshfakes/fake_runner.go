// This file was generated by counterfeiter
package boshfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-deployment-resource/bosh"
)

type FakeRunner struct {
	ExecuteStub        func(commandOpts interface{}) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		commandOpts interface{}
	}
	executeReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRunner) Execute(commandOpts interface{}) error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		commandOpts interface{}
	}{commandOpts})
	fake.recordInvocation("Execute", []interface{}{commandOpts})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(commandOpts)
	}
	return fake.executeReturns.result1
}

func (fake *FakeRunner) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeRunner) ExecuteArgsForCall(i int) interface{} {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].commandOpts
}

func (fake *FakeRunner) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRunner) recordInvocation(key string, args []interface{}) {
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

var _ bosh.Runner = new(FakeRunner)