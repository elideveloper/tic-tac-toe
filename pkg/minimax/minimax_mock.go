package minimax

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type StateMock struct {
	mock.Mock
}

func (m *StateMock) Eval() float64 {
	args := m.Called()
	fmt.Println("eval ", args.Get(0))
	return args.Get(0).(float64)
}

func (m *StateMock) GetChildren(isMaximizer bool) []State {
	args := m.Called(isMaximizer)
	return args.Get(0).([]State)
}

func (m *StateMock) Print() {
	m.Called()
}
