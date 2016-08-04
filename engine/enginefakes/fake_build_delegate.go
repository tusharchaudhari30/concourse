// This file was generated by counterfeiter
package enginefakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/engine"
	"github.com/concourse/atc/event"
	"github.com/concourse/atc/exec"
	"code.cloudfoundry.org/lager"
)

type FakeBuildDelegate struct {
	InputDelegateStub        func(lager.Logger, atc.GetPlan, event.OriginID) exec.GetDelegate
	inputDelegateMutex       sync.RWMutex
	inputDelegateArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.GetPlan
		arg3 event.OriginID
	}
	inputDelegateReturns struct {
		result1 exec.GetDelegate
	}
	ExecutionDelegateStub        func(lager.Logger, atc.TaskPlan, event.OriginID) exec.TaskDelegate
	executionDelegateMutex       sync.RWMutex
	executionDelegateArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.TaskPlan
		arg3 event.OriginID
	}
	executionDelegateReturns struct {
		result1 exec.TaskDelegate
	}
	OutputDelegateStub        func(lager.Logger, atc.PutPlan, event.OriginID) exec.PutDelegate
	outputDelegateMutex       sync.RWMutex
	outputDelegateArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.PutPlan
		arg3 event.OriginID
	}
	outputDelegateReturns struct {
		result1 exec.PutDelegate
	}
	FinishStub        func(lager.Logger, error, exec.Success, bool)
	finishMutex       sync.RWMutex
	finishArgsForCall []struct {
		arg1 lager.Logger
		arg2 error
		arg3 exec.Success
		arg4 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildDelegate) InputDelegate(arg1 lager.Logger, arg2 atc.GetPlan, arg3 event.OriginID) exec.GetDelegate {
	fake.inputDelegateMutex.Lock()
	fake.inputDelegateArgsForCall = append(fake.inputDelegateArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.GetPlan
		arg3 event.OriginID
	}{arg1, arg2, arg3})
	fake.recordInvocation("InputDelegate", []interface{}{arg1, arg2, arg3})
	fake.inputDelegateMutex.Unlock()
	if fake.InputDelegateStub != nil {
		return fake.InputDelegateStub(arg1, arg2, arg3)
	} else {
		return fake.inputDelegateReturns.result1
	}
}

func (fake *FakeBuildDelegate) InputDelegateCallCount() int {
	fake.inputDelegateMutex.RLock()
	defer fake.inputDelegateMutex.RUnlock()
	return len(fake.inputDelegateArgsForCall)
}

func (fake *FakeBuildDelegate) InputDelegateArgsForCall(i int) (lager.Logger, atc.GetPlan, event.OriginID) {
	fake.inputDelegateMutex.RLock()
	defer fake.inputDelegateMutex.RUnlock()
	return fake.inputDelegateArgsForCall[i].arg1, fake.inputDelegateArgsForCall[i].arg2, fake.inputDelegateArgsForCall[i].arg3
}

func (fake *FakeBuildDelegate) InputDelegateReturns(result1 exec.GetDelegate) {
	fake.InputDelegateStub = nil
	fake.inputDelegateReturns = struct {
		result1 exec.GetDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) ExecutionDelegate(arg1 lager.Logger, arg2 atc.TaskPlan, arg3 event.OriginID) exec.TaskDelegate {
	fake.executionDelegateMutex.Lock()
	fake.executionDelegateArgsForCall = append(fake.executionDelegateArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.TaskPlan
		arg3 event.OriginID
	}{arg1, arg2, arg3})
	fake.recordInvocation("ExecutionDelegate", []interface{}{arg1, arg2, arg3})
	fake.executionDelegateMutex.Unlock()
	if fake.ExecutionDelegateStub != nil {
		return fake.ExecutionDelegateStub(arg1, arg2, arg3)
	} else {
		return fake.executionDelegateReturns.result1
	}
}

func (fake *FakeBuildDelegate) ExecutionDelegateCallCount() int {
	fake.executionDelegateMutex.RLock()
	defer fake.executionDelegateMutex.RUnlock()
	return len(fake.executionDelegateArgsForCall)
}

func (fake *FakeBuildDelegate) ExecutionDelegateArgsForCall(i int) (lager.Logger, atc.TaskPlan, event.OriginID) {
	fake.executionDelegateMutex.RLock()
	defer fake.executionDelegateMutex.RUnlock()
	return fake.executionDelegateArgsForCall[i].arg1, fake.executionDelegateArgsForCall[i].arg2, fake.executionDelegateArgsForCall[i].arg3
}

func (fake *FakeBuildDelegate) ExecutionDelegateReturns(result1 exec.TaskDelegate) {
	fake.ExecutionDelegateStub = nil
	fake.executionDelegateReturns = struct {
		result1 exec.TaskDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) OutputDelegate(arg1 lager.Logger, arg2 atc.PutPlan, arg3 event.OriginID) exec.PutDelegate {
	fake.outputDelegateMutex.Lock()
	fake.outputDelegateArgsForCall = append(fake.outputDelegateArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.PutPlan
		arg3 event.OriginID
	}{arg1, arg2, arg3})
	fake.recordInvocation("OutputDelegate", []interface{}{arg1, arg2, arg3})
	fake.outputDelegateMutex.Unlock()
	if fake.OutputDelegateStub != nil {
		return fake.OutputDelegateStub(arg1, arg2, arg3)
	} else {
		return fake.outputDelegateReturns.result1
	}
}

func (fake *FakeBuildDelegate) OutputDelegateCallCount() int {
	fake.outputDelegateMutex.RLock()
	defer fake.outputDelegateMutex.RUnlock()
	return len(fake.outputDelegateArgsForCall)
}

func (fake *FakeBuildDelegate) OutputDelegateArgsForCall(i int) (lager.Logger, atc.PutPlan, event.OriginID) {
	fake.outputDelegateMutex.RLock()
	defer fake.outputDelegateMutex.RUnlock()
	return fake.outputDelegateArgsForCall[i].arg1, fake.outputDelegateArgsForCall[i].arg2, fake.outputDelegateArgsForCall[i].arg3
}

func (fake *FakeBuildDelegate) OutputDelegateReturns(result1 exec.PutDelegate) {
	fake.OutputDelegateStub = nil
	fake.outputDelegateReturns = struct {
		result1 exec.PutDelegate
	}{result1}
}

func (fake *FakeBuildDelegate) Finish(arg1 lager.Logger, arg2 error, arg3 exec.Success, arg4 bool) {
	fake.finishMutex.Lock()
	fake.finishArgsForCall = append(fake.finishArgsForCall, struct {
		arg1 lager.Logger
		arg2 error
		arg3 exec.Success
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Finish", []interface{}{arg1, arg2, arg3, arg4})
	fake.finishMutex.Unlock()
	if fake.FinishStub != nil {
		fake.FinishStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *FakeBuildDelegate) FinishCallCount() int {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return len(fake.finishArgsForCall)
}

func (fake *FakeBuildDelegate) FinishArgsForCall(i int) (lager.Logger, error, exec.Success, bool) {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return fake.finishArgsForCall[i].arg1, fake.finishArgsForCall[i].arg2, fake.finishArgsForCall[i].arg3, fake.finishArgsForCall[i].arg4
}

func (fake *FakeBuildDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.inputDelegateMutex.RLock()
	defer fake.inputDelegateMutex.RUnlock()
	fake.executionDelegateMutex.RLock()
	defer fake.executionDelegateMutex.RUnlock()
	fake.outputDelegateMutex.RLock()
	defer fake.outputDelegateMutex.RUnlock()
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBuildDelegate) recordInvocation(key string, args []interface{}) {
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

var _ engine.BuildDelegate = new(FakeBuildDelegate)
