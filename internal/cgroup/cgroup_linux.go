//go:build linux
// +build linux

package cgroup

import (
	"fmt"

	"github.com/containerd/cgroups"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/spf13/cast"
)

type CgroupsLimiter struct {
	controls []cgroups.Cgroup
}

// configure
func (r *CgroupsLimiter) Configure(pid int, core float64, mbn int) error {
	const (
		cpuUnit = 10000
		memUnit = 1024 * 1024
	)

	if core <= 0 {
		core = 1
	}

	var (
		quota  int64  = int64(core * cpuUnit) // core * 1u
		period uint64 = 10000                 // 1u
		mem    int64  = int64(mbn * memUnit)
	)

	cfg := &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Period: &period,
			Quota:  &quota,
		},
	}

	if mem != 0 {
		cfg.Memory = &specs.LinuxMemory{
			Limit: &mem,
		}
	}

	// file as /sys/fs/cgroup/cpu/netflow/...
	cgroupPath := "/netflow"
	control, err := cgroups.New(cgroups.V1, cgroups.StaticPath(cgroupPath), cfg)
	if err != nil {
		return err
	}

	r.controls = append(r.controls, control)
	err = control.Add(cgroups.Process{Pid: cast.ToInt(pid)})
	return err
}

// free
func (r *CgroupsLimiter) Free() error {
	for _, ctrl := range r.controls {
		ctrl.Delete()
	}
	fmt.Println("exit")
	return nil
}
