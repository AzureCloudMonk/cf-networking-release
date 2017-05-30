// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/models"
	"sync"
)

type Store struct {
	AllStub        func() ([]models.Policy, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []models.Policy
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []models.Policy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Store) All() ([]models.Policy, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *Store) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *Store) AllReturns(result1 []models.Policy, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *Store) AllReturnsOnCall(i int, result1 []models.Policy, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []models.Policy
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *Store) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Store) recordInvocation(key string, args []interface{}) {
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
