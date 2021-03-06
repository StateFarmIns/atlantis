// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: ProjectLocker)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
	models "github.com/runatlantis/atlantis/server/events/models"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockProjectLocker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectLocker() *MockProjectLocker {
	return &MockProjectLocker{fail: pegomock.GlobalFailHandler}
}

func (mock *MockProjectLocker) TryLock(log *logging.SimpleLogger, pull models.PullRequest, user models.User, workspace string, project models.Project) (*events.TryLockResponse, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectLocker().")
	}
	params := []pegomock.Param{log, pull, user, workspace, project}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLock", params, []reflect.Type{reflect.TypeOf((**events.TryLockResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *events.TryLockResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*events.TryLockResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockProjectLocker) VerifyWasCalledOnce() *VerifierProjectLocker {
	return &VerifierProjectLocker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectLocker) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierProjectLocker {
	return &VerifierProjectLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectLocker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierProjectLocker {
	return &VerifierProjectLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectLocker) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierProjectLocker {
	return &VerifierProjectLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierProjectLocker struct {
	mock                   *MockProjectLocker
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierProjectLocker) TryLock(log *logging.SimpleLogger, pull models.PullRequest, user models.User, workspace string, project models.Project) *ProjectLocker_TryLock_OngoingVerification {
	params := []pegomock.Param{log, pull, user, workspace, project}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLock", params, verifier.timeout)
	return &ProjectLocker_TryLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectLocker_TryLock_OngoingVerification struct {
	mock              *MockProjectLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectLocker_TryLock_OngoingVerification) GetCapturedArguments() (*logging.SimpleLogger, models.PullRequest, models.User, string, models.Project) {
	log, pull, user, workspace, project := c.GetAllCapturedArguments()
	return log[len(log)-1], pull[len(pull)-1], user[len(user)-1], workspace[len(workspace)-1], project[len(project)-1]
}

func (c *ProjectLocker_TryLock_OngoingVerification) GetAllCapturedArguments() (_param0 []*logging.SimpleLogger, _param1 []models.PullRequest, _param2 []models.User, _param3 []string, _param4 []models.Project) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*logging.SimpleLogger, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*logging.SimpleLogger)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.User, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(models.User)
		}
		_param3 = make([]string, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]models.Project, len(params[4]))
		for u, param := range params[4] {
			_param4[u] = param.(models.Project)
		}
	}
	return
}
