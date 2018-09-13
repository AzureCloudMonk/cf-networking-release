// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/db"
	"policy-server/store"
	"sync"
)

type Store struct {
	CreateStub        func([]store.Policy) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 []store.Policy
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	CreateWithTxStub        func(db.Transaction, []store.Policy) error
	createWithTxMutex       sync.RWMutex
	createWithTxArgsForCall []struct {
		arg1 db.Transaction
		arg2 []store.Policy
	}
	createWithTxReturns struct {
		result1 error
	}
	createWithTxReturnsOnCall map[int]struct {
		result1 error
	}
	AllStub        func() ([]store.Policy, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []store.Policy
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []store.Policy
		result2 error
	}
	DeleteStub        func([]store.Policy) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 []store.Policy
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteWithTxStub        func(db.Transaction, []store.Policy) error
	deleteWithTxMutex       sync.RWMutex
	deleteWithTxArgsForCall []struct {
		arg1 db.Transaction
		arg2 []store.Policy
	}
	deleteWithTxReturns struct {
		result1 error
	}
	deleteWithTxReturnsOnCall map[int]struct {
		result1 error
	}
	ByGuidsStub        func([]string, []string, bool) ([]store.Policy, error)
	byGuidsMutex       sync.RWMutex
	byGuidsArgsForCall []struct {
		arg1 []string
		arg2 []string
		arg3 bool
	}
	byGuidsReturns struct {
		result1 []store.Policy
		result2 error
	}
	byGuidsReturnsOnCall map[int]struct {
		result1 []store.Policy
		result2 error
	}
	CheckDatabaseStub        func() error
	checkDatabaseMutex       sync.RWMutex
	checkDatabaseArgsForCall []struct{}
	checkDatabaseReturns     struct {
		result1 error
	}
	checkDatabaseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Store) Create(arg1 []store.Policy) error {
	var arg1Copy []store.Policy
	if arg1 != nil {
		arg1Copy = make([]store.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 []store.Policy
	}{arg1Copy})
	fake.recordInvocation("Create", []interface{}{arg1Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *Store) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Store) CreateArgsForCall(i int) []store.Policy {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *Store) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Store) CreateWithTx(arg1 db.Transaction, arg2 []store.Policy) error {
	var arg2Copy []store.Policy
	if arg2 != nil {
		arg2Copy = make([]store.Policy, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.createWithTxMutex.Lock()
	ret, specificReturn := fake.createWithTxReturnsOnCall[len(fake.createWithTxArgsForCall)]
	fake.createWithTxArgsForCall = append(fake.createWithTxArgsForCall, struct {
		arg1 db.Transaction
		arg2 []store.Policy
	}{arg1, arg2Copy})
	fake.recordInvocation("CreateWithTx", []interface{}{arg1, arg2Copy})
	fake.createWithTxMutex.Unlock()
	if fake.CreateWithTxStub != nil {
		return fake.CreateWithTxStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createWithTxReturns.result1
}

func (fake *Store) CreateWithTxCallCount() int {
	fake.createWithTxMutex.RLock()
	defer fake.createWithTxMutex.RUnlock()
	return len(fake.createWithTxArgsForCall)
}

func (fake *Store) CreateWithTxArgsForCall(i int) (db.Transaction, []store.Policy) {
	fake.createWithTxMutex.RLock()
	defer fake.createWithTxMutex.RUnlock()
	return fake.createWithTxArgsForCall[i].arg1, fake.createWithTxArgsForCall[i].arg2
}

func (fake *Store) CreateWithTxReturns(result1 error) {
	fake.CreateWithTxStub = nil
	fake.createWithTxReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) CreateWithTxReturnsOnCall(i int, result1 error) {
	fake.CreateWithTxStub = nil
	if fake.createWithTxReturnsOnCall == nil {
		fake.createWithTxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createWithTxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Store) All() ([]store.Policy, error) {
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

func (fake *Store) AllReturns(result1 []store.Policy, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []store.Policy
		result2 error
	}{result1, result2}
}

func (fake *Store) AllReturnsOnCall(i int, result1 []store.Policy, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []store.Policy
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []store.Policy
		result2 error
	}{result1, result2}
}

func (fake *Store) Delete(arg1 []store.Policy) error {
	var arg1Copy []store.Policy
	if arg1 != nil {
		arg1Copy = make([]store.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 []store.Policy
	}{arg1Copy})
	fake.recordInvocation("Delete", []interface{}{arg1Copy})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *Store) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *Store) DeleteArgsForCall(i int) []store.Policy {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

func (fake *Store) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Store) DeleteWithTx(arg1 db.Transaction, arg2 []store.Policy) error {
	var arg2Copy []store.Policy
	if arg2 != nil {
		arg2Copy = make([]store.Policy, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deleteWithTxMutex.Lock()
	ret, specificReturn := fake.deleteWithTxReturnsOnCall[len(fake.deleteWithTxArgsForCall)]
	fake.deleteWithTxArgsForCall = append(fake.deleteWithTxArgsForCall, struct {
		arg1 db.Transaction
		arg2 []store.Policy
	}{arg1, arg2Copy})
	fake.recordInvocation("DeleteWithTx", []interface{}{arg1, arg2Copy})
	fake.deleteWithTxMutex.Unlock()
	if fake.DeleteWithTxStub != nil {
		return fake.DeleteWithTxStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteWithTxReturns.result1
}

func (fake *Store) DeleteWithTxCallCount() int {
	fake.deleteWithTxMutex.RLock()
	defer fake.deleteWithTxMutex.RUnlock()
	return len(fake.deleteWithTxArgsForCall)
}

func (fake *Store) DeleteWithTxArgsForCall(i int) (db.Transaction, []store.Policy) {
	fake.deleteWithTxMutex.RLock()
	defer fake.deleteWithTxMutex.RUnlock()
	return fake.deleteWithTxArgsForCall[i].arg1, fake.deleteWithTxArgsForCall[i].arg2
}

func (fake *Store) DeleteWithTxReturns(result1 error) {
	fake.DeleteWithTxStub = nil
	fake.deleteWithTxReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) DeleteWithTxReturnsOnCall(i int, result1 error) {
	fake.DeleteWithTxStub = nil
	if fake.deleteWithTxReturnsOnCall == nil {
		fake.deleteWithTxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteWithTxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Store) ByGuids(arg1 []string, arg2 []string, arg3 bool) ([]store.Policy, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.byGuidsMutex.Lock()
	ret, specificReturn := fake.byGuidsReturnsOnCall[len(fake.byGuidsArgsForCall)]
	fake.byGuidsArgsForCall = append(fake.byGuidsArgsForCall, struct {
		arg1 []string
		arg2 []string
		arg3 bool
	}{arg1Copy, arg2Copy, arg3})
	fake.recordInvocation("ByGuids", []interface{}{arg1Copy, arg2Copy, arg3})
	fake.byGuidsMutex.Unlock()
	if fake.ByGuidsStub != nil {
		return fake.ByGuidsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.byGuidsReturns.result1, fake.byGuidsReturns.result2
}

func (fake *Store) ByGuidsCallCount() int {
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	return len(fake.byGuidsArgsForCall)
}

func (fake *Store) ByGuidsArgsForCall(i int) ([]string, []string, bool) {
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	return fake.byGuidsArgsForCall[i].arg1, fake.byGuidsArgsForCall[i].arg2, fake.byGuidsArgsForCall[i].arg3
}

func (fake *Store) ByGuidsReturns(result1 []store.Policy, result2 error) {
	fake.ByGuidsStub = nil
	fake.byGuidsReturns = struct {
		result1 []store.Policy
		result2 error
	}{result1, result2}
}

func (fake *Store) ByGuidsReturnsOnCall(i int, result1 []store.Policy, result2 error) {
	fake.ByGuidsStub = nil
	if fake.byGuidsReturnsOnCall == nil {
		fake.byGuidsReturnsOnCall = make(map[int]struct {
			result1 []store.Policy
			result2 error
		})
	}
	fake.byGuidsReturnsOnCall[i] = struct {
		result1 []store.Policy
		result2 error
	}{result1, result2}
}

func (fake *Store) CheckDatabase() error {
	fake.checkDatabaseMutex.Lock()
	ret, specificReturn := fake.checkDatabaseReturnsOnCall[len(fake.checkDatabaseArgsForCall)]
	fake.checkDatabaseArgsForCall = append(fake.checkDatabaseArgsForCall, struct{}{})
	fake.recordInvocation("CheckDatabase", []interface{}{})
	fake.checkDatabaseMutex.Unlock()
	if fake.CheckDatabaseStub != nil {
		return fake.CheckDatabaseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkDatabaseReturns.result1
}

func (fake *Store) CheckDatabaseCallCount() int {
	fake.checkDatabaseMutex.RLock()
	defer fake.checkDatabaseMutex.RUnlock()
	return len(fake.checkDatabaseArgsForCall)
}

func (fake *Store) CheckDatabaseReturns(result1 error) {
	fake.CheckDatabaseStub = nil
	fake.checkDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) CheckDatabaseReturnsOnCall(i int, result1 error) {
	fake.CheckDatabaseStub = nil
	if fake.checkDatabaseReturnsOnCall == nil {
		fake.checkDatabaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkDatabaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Store) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.createWithTxMutex.RLock()
	defer fake.createWithTxMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteWithTxMutex.RLock()
	defer fake.deleteWithTxMutex.RUnlock()
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	fake.checkDatabaseMutex.RLock()
	defer fake.checkDatabaseMutex.RUnlock()
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

var _ store.Store = new(Store)
