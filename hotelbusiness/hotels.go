//go:build !solution

package hotelbusiness

import "sort"

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

type trafficState struct {
	ins  int
	outs int
}

func ComputeLoad(guests []Guest) []Load {
	states := make(map[int]*trafficState)

	for _, g := range guests {
		if state, ok := states[g.CheckInDate]; !ok {
			states[g.CheckInDate] = &trafficState{1, 0}
		} else {
			state.ins++
			if state.ins == state.outs {
				delete(states, g.CheckInDate)
			}
		}

		if state, ok := states[g.CheckOutDate]; !ok {
			states[g.CheckOutDate] = &trafficState{0, 1}
		} else {
			state.outs++

			if state.ins == state.outs {
				delete(states, g.CheckOutDate)
			}
		}
	}

	keys := make([]int, 0, len(states))
	for k := range states {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	load := 0
	loads := make([]Load, 0)

	for _, day := range keys {
		if state, ok := states[day]; !ok || state.ins == state.outs {
			continue
		} else {
			load += state.ins - state.outs
			loads = append(loads, Load{StartDate: day, GuestCount: load})
		}
	}

	return loads
}
