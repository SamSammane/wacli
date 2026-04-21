package main

import "testing"

func TestParseLockOwnerPID(t *testing.T) {
	tests := []struct {
		name string
		info string
		want int
	}{
		{name: "pid line", info: "pid=50394\nacquired_at=2026-04-05T12:30:11Z", want: 50394},
		{name: "trimmed pid", info: " pid= 42 ", want: 42},
		{name: "missing pid", info: "acquired_at=2026-04-05T12:30:11Z"},
		{name: "invalid pid", info: "pid=abc"},
		{name: "zero pid", info: "pid=0"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := parseLockOwnerPID(tc.info); got != tc.want {
				t.Fatalf("parseLockOwnerPID() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestDoctorConnectionState(t *testing.T) {
	tests := []struct {
		name      string
		authed    bool
		connected bool
		lockHeld  bool
		connect   bool
		want      string
	}{
		{name: "connected wins", authed: true, connected: true, lockHeld: true, want: "connected"},
		{name: "locked paired session", authed: true, lockHeld: true, want: "locked_by_other_process"},
		{name: "connect requested stays disconnected", authed: true, lockHeld: true, connect: true, want: "disconnected"},
		{name: "plain disconnected", authed: true, want: "disconnected"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := doctorConnectionState(tc.authed, tc.connected, tc.lockHeld, tc.connect)
			if got != tc.want {
				t.Fatalf("doctorConnectionState() = %q, want %q", got, tc.want)
			}
		})
	}
}
