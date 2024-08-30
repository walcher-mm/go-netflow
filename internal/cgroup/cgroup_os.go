//go:build !linux
// +build !linux

package cgroup

import (
	"errors"
)

type CgroupsLimiter struct {
}

func (r *CgroupsLimiter) Free() error {
	return nil
}

func (r *CgroupsLimiter) Configure(pid int, core float64, mb int) error {
	return errors.New("don't support cgroup")
}
