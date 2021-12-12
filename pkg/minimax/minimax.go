package minimax

import (
	"fmt"
)

const (
	maxPossibleVal      = 999_999_999.99
	minPossibleVal      = -999_999_999.99
	maxDepth       uint = 6 // bot's moves num is (maxDepth / 2)
)

type State interface {
	Eval() float64
	EvalNotTerminal() float64
	GetChildren(isMaximizer bool) []State
}

func FindBestUsingMinimax(currState State, isMaximizer bool) State {
	val, bestState := minimax(currState, minPossibleVal, maxPossibleVal, isMaximizer, 1)
	if val > 0.0 {
		fmt.Println("found a good move, val: ", val)
	} else {
		fmt.Println("did not find a good move, val: ", val)

		fmt.Println("searching for heuristic evaluation")
		ch := currState.GetChildren(isMaximizer)
		var bestVal float64
		for i := range ch {
			v := ch[i].EvalNotTerminal()
			if v > bestVal {
				bestState = ch[i]
				bestVal = v
			}
		}

	}
	return bestState
}

func minimax(state State, alpha, beta float64, isMaximizer bool, depth uint) (float64, State) {
	children := state.GetChildren(isMaximizer)
	if len(children) == 0 {
		return state.Eval() / float64(depth), state
	}

	if depth >= maxDepth {
		return 0, children[0]
	}

	var bestState State
	var bestVal float64
	if isMaximizer {
		maxVal := minPossibleVal
		for i, chState := range children {
			val, _ := minimax(chState, alpha, beta, false, depth+1)
			if i == 0 || val > maxVal {
				maxVal = val
				bestState = chState

				if val > alpha {
					alpha = val
					// if worst play for maximizer becomes better than max available
					// then no need to search more
					if alpha >= beta {
						break
					}
				}
			}
		}
		bestVal = maxVal
	} else {
		minVal := maxPossibleVal
		for i, chState := range children {
			val, _ := minimax(chState, alpha, beta, true, depth+1)
			if i == 0 || val < minVal {
				minVal = val
				bestState = chState

				if val < beta {
					beta = val
					if beta <= alpha {
						break
					}
				}
			}
		}
		bestVal = minVal
	}

	return bestVal, bestState
}
