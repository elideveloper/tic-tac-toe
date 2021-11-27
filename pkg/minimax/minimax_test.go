package minimax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiniMax(t *testing.T) {
	root := &StateMock{}
	root.On("GetChildren", true).Return([]State{ // initial state's depth=0
		func() *StateMock {
			m3 := &StateMock{}
			t.Log("called children 3 from 3,-4")
			m3.On("GetChildren", false).Return([]State{ // depth=1
				func() *StateMock {
					m33 := &StateMock{}
					t.Log("called children 3 from 3,5")
					m33.On("GetChildren", true).Return([]State{ // depth=2
						func() *StateMock {
							m := &StateMock{}
							m.On("GetChildren", false).Return([]State{}) // empty for terminal state
							m.On("Eval").Return(-1.0)
							t.Log("called -1")
							return m
						}(),
						func() *StateMock {
							m := &StateMock{}
							m.On("GetChildren", false).Return([]State{}) // empty for terminal state
							m.On("Eval").Return(3.0)
							t.Log("called 3")
							return m
						}(),
					})
					return m33
				}(),
				func() *StateMock {
					m := &StateMock{}
					t.Log("called children 5 from 3,5")
					m.On("GetChildren", true).Return([]State{ // depth=2
						func() *StateMock {
							m := &StateMock{}
							m.On("GetChildren", false).Return([]State{}) // empty for terminal state
							m.On("Eval").Return(5.0)
							t.Log("called 5")
							return m
						}(),
						func() *StateMock {
							m := &StateMock{}
							m.On("GetChildren", false).Return([]State{}) // empty for terminal state
							m.On("Eval").Return(1.0)
							t.Log("called 1")
							return m
						}(),
					})
					return m
				}(),
			})
			return m3
		}(),

		func() *StateMock {
			m := &StateMock{}
			t.Log("called children -4 from 3,-4")
			m.On("GetChildren", false).Return([]State{ // depth=1
				func() *StateMock {
					m := &StateMock{}
					t.Log("called children -4 from -4,9")
					m.On("GetChildren", true).Return([]State{ // depth=2
						func() *StateMock {
							m := &StateMock{}
							m.On("GetChildren", false).Return([]State{}) // empty for terminal state
							m.On("Eval").Return(-6.0)
							t.Log("called -6")
							return m
						}(),
						func() *StateMock {
							m := &StateMock{}
							m.On("GetChildren", false).Return([]State{}) // empty for terminal state
							m.On("Eval").Return(-4.0)
							t.Log("called -4")
							return m
						}(),
					})
					return m
				}(),
				func() *StateMock {
					m := &StateMock{}
					t.Log("called children 9 from -4,9")
					m.On("GetChildren", true).Return([]State{ // depth=2
						func() *StateMock {
							m := &StateMock{}
							m.On("GetChildren", false).Return([]State{}) // empty for terminal state
							m.On("Eval").Return(0.0)
							t.Log("called 0")
							return m
						}(),
						func() *StateMock {
							m := &StateMock{}
							m.On("GetChildren", false).Return([]State{}) // empty for terminal state
							m.On("Eval").Return(9.0)
							t.Log("called 9")
							return m
						}(),
					})
					return m
				}(),
			})
			return m
		}(),
	})

	val, _ := minimax(root, minPossibleVal, maxPossibleVal, true, 1)
	depth := 4
	needVal := 3.0 / float64(depth)
	assert.Equal(t, needVal, val)
}
