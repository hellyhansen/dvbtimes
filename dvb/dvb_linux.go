package dvb

import (
  "bytes"
  "encoding/binary"
  "github.com/hellyhansen/dvbtimes/ioctl"
  "os"
)

type FrontendInfo struct {
  Name [128]byte
  DeprecatedType uint32
  FreqMin uint32
  FreqMax uint32
  FreqStepSize uint32
  FreqTolerance uint32
  SymbolRateMin uint32
  SymbolRateMax uint32
  SymbolRateTolerance uint32
  DeprecatedNotifierDelay uint32
  Caps uint32
}

func (fi *FrontendInfo) GetName() string {
  return string(fi.Name[:])
}

var (
  getFrontendInfo = ioctl.NewIoctl(ioctl.READ, 'o', 61, binary.Size(&FrontendInfo{}))
)

func GetFrontendInfo(f *os.File) (*FrontendInfo, error) {
  fi := &FrontendInfo{}
  arg := getFrontendInfo.Slice()
  if err := getFrontendInfo.Call(f, arg); err != nil {
    return nil, err
  }
  if err := binary.Read(bytes.NewReader(arg), binary.LittleEndian, fi); err != nil {
    panic(err)
  }
  return fi, nil
}