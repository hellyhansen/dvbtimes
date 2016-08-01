package ioctl

import (
  "syscall"
  "os"
  "unsafe"
)

type Direction uint8

const (
  NONE Direction = 0
  WRITE Direction = 1
  READ Direction = 2
  READ_WRITE Direction = 3
)

const sizeMask = 0x3fff

type Ioctl struct {
  request uintptr
}

func (ioctl *Ioctl) Size() int {
  return int(ioctl.request >> 16) & sizeMask
}

func (ioctl *Ioctl) Slice() []byte {
  return make([]byte, ioctl.Size())
}

func (ioctl *Ioctl) Call(f *os.File, arg []byte) error {
  length := len(arg)
  if length < ioctl.Size() {
    panic("ioctl arg too small")
  }
  var argp uintptr
  if length > 0 {
    argp = uintptr(unsafe.Pointer(&arg[0]))
  }
  if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), ioctl.request, argp); errno != 0 {
    return os.NewSyscallError("ioctl", errno)
  }
  return nil
}

func NewIoctl(dir Direction, t, nr uint8, size int) *Ioctl {
  if size < 0 || size > sizeMask {
    panic("ioctl size too large")
  }
  return &Ioctl{
    request: uintptr(nr) | uintptr(t) << 8 | uintptr(size) << 16 | uintptr(dir) << 30,
  }
}
