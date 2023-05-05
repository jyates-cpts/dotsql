// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package dotsql

import (
	"context"
	"database/sql"
	"sync"
)

var (
	lockPreparerMockPrepare sync.RWMutex
)

// Ensure, that PreparerMock does implement Preparer.
// If this is not the case, regenerate this file with moq.
var _ Preparer = &PreparerMock{}

// PreparerMock is a mock implementation of Preparer.
//
//     func TestSomethingThatUsesPreparer(t *testing.T) {
//
//         // make and configure a mocked Preparer
//         mockedPreparer := &PreparerMock{
//             PrepareFunc: func(query string) (*sql.Stmt, error) {
// 	               panic("mock out the Prepare method")
//             },
//         }
//
//         // use mockedPreparer in code that requires Preparer
//         // and then make assertions.
//
//     }
type PreparerMock struct {
	// PrepareFunc mocks the Prepare method.
	PrepareFunc func(query string) (*sql.Stmt, error)

	// calls tracks calls to the methods.
	calls struct {
		// Prepare holds details about calls to the Prepare method.
		Prepare []struct {
			// Query is the query argument value.
			Query string
		}
	}
}

// Prepare calls PrepareFunc.
func (mock *PreparerMock) Prepare(query string) (*sql.Stmt, error) {
	if mock.PrepareFunc == nil {
		panic("PreparerMock.PrepareFunc: method is nil but Preparer.Prepare was just called")
	}
	callInfo := struct {
		Query string
	}{
		Query: query,
	}
	lockPreparerMockPrepare.Lock()
	mock.calls.Prepare = append(mock.calls.Prepare, callInfo)
	lockPreparerMockPrepare.Unlock()
	return mock.PrepareFunc(query)
}

// PrepareCalls gets all the calls that were made to Prepare.
// Check the length with:
//     len(mockedPreparer.PrepareCalls())
func (mock *PreparerMock) PrepareCalls() []struct {
	Query string
} {
	var calls []struct {
		Query string
	}
	lockPreparerMockPrepare.RLock()
	calls = mock.calls.Prepare
	lockPreparerMockPrepare.RUnlock()
	return calls
}

var (
	lockPreparerContextMockPrepareContext sync.RWMutex
)

// Ensure, that PreparerContextMock does implement PreparerContext.
// If this is not the case, regenerate this file with moq.
var _ PreparerContext = &PreparerContextMock{}

// PreparerContextMock is a mock implementation of PreparerContext.
//
//     func TestSomethingThatUsesPreparerContext(t *testing.T) {
//
//         // make and configure a mocked PreparerContext
//         mockedPreparerContext := &PreparerContextMock{
//             PrepareContextFunc: func(ctx context.Context, query string) (*sql.Stmt, error) {
// 	               panic("mock out the PrepareContext method")
//             },
//         }
//
//         // use mockedPreparerContext in code that requires PreparerContext
//         // and then make assertions.
//
//     }
type PreparerContextMock struct {
	// PrepareContextFunc mocks the PrepareContext method.
	PrepareContextFunc func(ctx context.Context, query string) (*sql.Stmt, error)

	// calls tracks calls to the methods.
	calls struct {
		// PrepareContext holds details about calls to the PrepareContext method.
		PrepareContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
		}
	}
}

// PrepareContext calls PrepareContextFunc.
func (mock *PreparerContextMock) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if mock.PrepareContextFunc == nil {
		panic("PreparerContextMock.PrepareContextFunc: method is nil but PreparerContext.PrepareContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
	}{
		Ctx:   ctx,
		Query: query,
	}
	lockPreparerContextMockPrepareContext.Lock()
	mock.calls.PrepareContext = append(mock.calls.PrepareContext, callInfo)
	lockPreparerContextMockPrepareContext.Unlock()
	return mock.PrepareContextFunc(ctx, query)
}

// PrepareContextCalls gets all the calls that were made to PrepareContext.
// Check the length with:
//     len(mockedPreparerContext.PrepareContextCalls())
func (mock *PreparerContextMock) PrepareContextCalls() []struct {
	Ctx   context.Context
	Query string
} {
	var calls []struct {
		Ctx   context.Context
		Query string
	}
	lockPreparerContextMockPrepareContext.RLock()
	calls = mock.calls.PrepareContext
	lockPreparerContextMockPrepareContext.RUnlock()
	return calls
}

var (
	lockQueryerMockQuery sync.RWMutex
)

// Ensure, that QueryerMock does implement Queryer.
// If this is not the case, regenerate this file with moq.
var _ Queryer = &QueryerMock{}

// QueryerMock is a mock implementation of Queryer.
//
//     func TestSomethingThatUsesQueryer(t *testing.T) {
//
//         // make and configure a mocked Queryer
//         mockedQueryer := &QueryerMock{
//             QueryFunc: func(query string, args ...interface{}) (*sql.Rows, error) {
// 	               panic("mock out the Query method")
//             },
//         }
//
//         // use mockedQueryer in code that requires Queryer
//         // and then make assertions.
//
//     }
type QueryerMock struct {
	// QueryFunc mocks the Query method.
	QueryFunc func(query string, args ...interface{}) (*sql.Rows, error)

	// calls tracks calls to the methods.
	calls struct {
		// Query holds details about calls to the Query method.
		Query []struct {
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// Query calls QueryFunc.
func (mock *QueryerMock) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if mock.QueryFunc == nil {
		panic("QueryerMock.QueryFunc: method is nil but Queryer.Query was just called")
	}
	callInfo := struct {
		Query string
		Args  []interface{}
	}{
		Query: query,
		Args:  args,
	}
	lockQueryerMockQuery.Lock()
	mock.calls.Query = append(mock.calls.Query, callInfo)
	lockQueryerMockQuery.Unlock()
	return mock.QueryFunc(query, args...)
}

// QueryCalls gets all the calls that were made to Query.
// Check the length with:
//     len(mockedQueryer.QueryCalls())
func (mock *QueryerMock) QueryCalls() []struct {
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Query string
		Args  []interface{}
	}
	lockQueryerMockQuery.RLock()
	calls = mock.calls.Query
	lockQueryerMockQuery.RUnlock()
	return calls
}

var (
	lockQueryerContextMockQueryx sync.RWMutex
)

// Ensure, that QueryerContextMock does implement QueryerContext.
// If this is not the case, regenerate this file with moq.
var _ QueryerContext = &QueryerContextMock{}

// QueryerContextMock is a mock implementation of QueryerContext.
//
//     func TestSomethingThatUsesQueryerContext(t *testing.T) {
//
//         // make and configure a mocked QueryerContext
//         mockedQueryerContext := &QueryerContextMock{
//             QueryxFunc: func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
// 	               panic("mock out the Queryx method")
//             },
//         }
//
//         // use mockedQueryerContext in code that requires QueryerContext
//         // and then make assertions.
//
//     }
type QueryerContextMock struct {
	// QueryxFunc mocks the Queryx method.
	QueryxFunc func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// calls tracks calls to the methods.
	calls struct {
		// Queryx holds details about calls to the Queryx method.
		Queryx []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// Queryx calls QueryxFunc.
func (mock *QueryerContextMock) Queryx(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if mock.QueryxFunc == nil {
		panic("QueryerContextMock.QueryxFunc: method is nil but QueryerContext.Queryx was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []interface{}
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	lockQueryerContextMockQueryx.Lock()
	mock.calls.Queryx = append(mock.calls.Queryx, callInfo)
	lockQueryerContextMockQueryx.Unlock()
	return mock.QueryxFunc(ctx, query, args...)
}

// QueryxCalls gets all the calls that were made to Queryx.
// Check the length with:
//     len(mockedQueryerContext.QueryxCalls())
func (mock *QueryerContextMock) QueryxCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []interface{}
	}
	lockQueryerContextMockQueryx.RLock()
	calls = mock.calls.Queryx
	lockQueryerContextMockQueryx.RUnlock()
	return calls
}

var (
	lockQueryRowerMockQueryRow sync.RWMutex
)

// Ensure, that QueryRowerMock does implement QueryRower.
// If this is not the case, regenerate this file with moq.
var _ QueryRower = &QueryRowerMock{}

// QueryRowerMock is a mock implementation of QueryRower.
//
//     func TestSomethingThatUsesQueryRower(t *testing.T) {
//
//         // make and configure a mocked QueryRower
//         mockedQueryRower := &QueryRowerMock{
//             QueryRowFunc: func(query string, args ...interface{}) *sql.Row {
// 	               panic("mock out the QueryRow method")
//             },
//         }
//
//         // use mockedQueryRower in code that requires QueryRower
//         // and then make assertions.
//
//     }
type QueryRowerMock struct {
	// QueryRowFunc mocks the QueryRow method.
	QueryRowFunc func(query string, args ...interface{}) *sql.Row

	// calls tracks calls to the methods.
	calls struct {
		// QueryRow holds details about calls to the QueryRow method.
		QueryRow []struct {
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// QueryRow calls QueryRowFunc.
func (mock *QueryRowerMock) QueryRow(query string, args ...interface{}) *sql.Row {
	if mock.QueryRowFunc == nil {
		panic("QueryRowerMock.QueryRowFunc: method is nil but QueryRower.QueryRow was just called")
	}
	callInfo := struct {
		Query string
		Args  []interface{}
	}{
		Query: query,
		Args:  args,
	}
	lockQueryRowerMockQueryRow.Lock()
	mock.calls.QueryRow = append(mock.calls.QueryRow, callInfo)
	lockQueryRowerMockQueryRow.Unlock()
	return mock.QueryRowFunc(query, args...)
}

// QueryRowCalls gets all the calls that were made to QueryRow.
// Check the length with:
//     len(mockedQueryRower.QueryRowCalls())
func (mock *QueryRowerMock) QueryRowCalls() []struct {
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Query string
		Args  []interface{}
	}
	lockQueryRowerMockQueryRow.RLock()
	calls = mock.calls.QueryRow
	lockQueryRowerMockQueryRow.RUnlock()
	return calls
}

var (
	lockQueryRowerContextMockQueryRowContext sync.RWMutex
)

// Ensure, that QueryRowerContextMock does implement QueryRowerContext.
// If this is not the case, regenerate this file with moq.
var _ QueryRowerContext = &QueryRowerContextMock{}

// QueryRowerContextMock is a mock implementation of QueryRowerContext.
//
//     func TestSomethingThatUsesQueryRowerContext(t *testing.T) {
//
//         // make and configure a mocked QueryRowerContext
//         mockedQueryRowerContext := &QueryRowerContextMock{
//             QueryRowContextFunc: func(ctx context.Context, query string, args ...interface{}) *sql.Row {
// 	               panic("mock out the QueryRowContext method")
//             },
//         }
//
//         // use mockedQueryRowerContext in code that requires QueryRowerContext
//         // and then make assertions.
//
//     }
type QueryRowerContextMock struct {
	// QueryRowContextFunc mocks the QueryRowContext method.
	QueryRowContextFunc func(ctx context.Context, query string, args ...interface{}) *sql.Row

	// calls tracks calls to the methods.
	calls struct {
		// QueryRowContext holds details about calls to the QueryRowContext method.
		QueryRowContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// QueryRowContext calls QueryRowContextFunc.
func (mock *QueryRowerContextMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if mock.QueryRowContextFunc == nil {
		panic("QueryRowerContextMock.QueryRowContextFunc: method is nil but QueryRowerContext.QueryRowContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []interface{}
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	lockQueryRowerContextMockQueryRowContext.Lock()
	mock.calls.QueryRowContext = append(mock.calls.QueryRowContext, callInfo)
	lockQueryRowerContextMockQueryRowContext.Unlock()
	return mock.QueryRowContextFunc(ctx, query, args...)
}

// QueryRowContextCalls gets all the calls that were made to QueryRowContext.
// Check the length with:
//     len(mockedQueryRowerContext.QueryRowContextCalls())
func (mock *QueryRowerContextMock) QueryRowContextCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []interface{}
	}
	lockQueryRowerContextMockQueryRowContext.RLock()
	calls = mock.calls.QueryRowContext
	lockQueryRowerContextMockQueryRowContext.RUnlock()
	return calls
}

var (
	lockExecerMockExec sync.RWMutex
)

// Ensure, that ExecerMock does implement Execer.
// If this is not the case, regenerate this file with moq.
var _ Execer = &ExecerMock{}

// ExecerMock is a mock implementation of Execer.
//
//     func TestSomethingThatUsesExecer(t *testing.T) {
//
//         // make and configure a mocked Execer
//         mockedExecer := &ExecerMock{
//             ExecFunc: func(query string, args ...interface{}) (sql.Result, error) {
// 	               panic("mock out the Exec method")
//             },
//         }
//
//         // use mockedExecer in code that requires Execer
//         // and then make assertions.
//
//     }
type ExecerMock struct {
	// ExecFunc mocks the Exec method.
	ExecFunc func(query string, args ...interface{}) (sql.Result, error)

	// calls tracks calls to the methods.
	calls struct {
		// Exec holds details about calls to the Exec method.
		Exec []struct {
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// Exec calls ExecFunc.
func (mock *ExecerMock) Exec(query string, args ...interface{}) (sql.Result, error) {
	if mock.ExecFunc == nil {
		panic("ExecerMock.ExecFunc: method is nil but Execer.Exec was just called")
	}
	callInfo := struct {
		Query string
		Args  []interface{}
	}{
		Query: query,
		Args:  args,
	}
	lockExecerMockExec.Lock()
	mock.calls.Exec = append(mock.calls.Exec, callInfo)
	lockExecerMockExec.Unlock()
	return mock.ExecFunc(query, args...)
}

// ExecCalls gets all the calls that were made to Exec.
// Check the length with:
//     len(mockedExecer.ExecCalls())
func (mock *ExecerMock) ExecCalls() []struct {
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Query string
		Args  []interface{}
	}
	lockExecerMockExec.RLock()
	calls = mock.calls.Exec
	lockExecerMockExec.RUnlock()
	return calls
}

var (
	lockExecerContextMockExecContext sync.RWMutex
)

// Ensure, that ExecerContextMock does implement ExecerContext.
// If this is not the case, regenerate this file with moq.
var _ ExecerContext = &ExecerContextMock{}

// ExecerContextMock is a mock implementation of ExecerContext.
//
//     func TestSomethingThatUsesExecerContext(t *testing.T) {
//
//         // make and configure a mocked ExecerContext
//         mockedExecerContext := &ExecerContextMock{
//             ExecContextFunc: func(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
// 	               panic("mock out the ExecContext method")
//             },
//         }
//
//         // use mockedExecerContext in code that requires ExecerContext
//         // and then make assertions.
//
//     }
type ExecerContextMock struct {
	// ExecContextFunc mocks the ExecContext method.
	ExecContextFunc func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// calls tracks calls to the methods.
	calls struct {
		// ExecContext holds details about calls to the ExecContext method.
		ExecContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// ExecContext calls ExecContextFunc.
func (mock *ExecerContextMock) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if mock.ExecContextFunc == nil {
		panic("ExecerContextMock.ExecContextFunc: method is nil but ExecerContext.ExecContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []interface{}
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	lockExecerContextMockExecContext.Lock()
	mock.calls.ExecContext = append(mock.calls.ExecContext, callInfo)
	lockExecerContextMockExecContext.Unlock()
	return mock.ExecContextFunc(ctx, query, args...)
}

// ExecContextCalls gets all the calls that were made to ExecContext.
// Check the length with:
//     len(mockedExecerContext.ExecContextCalls())
func (mock *ExecerContextMock) ExecContextCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []interface{}
	}
	lockExecerContextMockExecContext.RLock()
	calls = mock.calls.ExecContext
	lockExecerContextMockExecContext.RUnlock()
	return calls
}
