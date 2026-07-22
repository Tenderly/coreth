// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

import "testing"

func TestMinimumGasConsumption(t *testing.T) {
	tests := []struct {
		name  string
		limit uint64
		want  uint64
	}{
		{name: "zero", limit: 0, want: 0},
		{name: "one rounds up", limit: 1, want: 1},
		{name: "even", limit: 21000, want: 10500},
		{name: "odd rounds up", limit: 21001, want: 10501},
		{name: "two", limit: 2, want: 1},
		{name: "three rounds up", limit: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumGasConsumption(tt.limit); got != tt.want {
				t.Fatalf("minimumGasConsumption(%d) = %d, want %d", tt.limit, got, tt.want)
			}
		})
	}
}
