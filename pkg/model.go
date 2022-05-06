package pkg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync"
)

type Package struct {
	Version [2]byte // 版本号
	Length  int16   // 数据长度
	Data    []byte  // 数据
}

type PackageType byte

const (
	PData      PackageType = '0'
	PHeartbeat PackageType = '1'
	PNewConn   PackageType = '2'
)

func NewPackage(packageType PackageType, data []byte) *Package {
	pk := &Package{
		Version: [2]byte{'V', byte(packageType)},
		Data:    data,
	}
	pk.Length = int16(len(data))
	return pk
}

func (p *Package) Pack(write io.Writer) error {
	var err error
	err = binary.Write(write, binary.BigEndian, &p.Version)
	if err != nil {
		return err
	}

	err = binary.Write(write, binary.BigEndian, &p.Length)
	if err != nil {
		return err
	}
	err = binary.Write(write, binary.BigEndian, &p.Data)
	return err
}

func (p *Package) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.Version)
	if err != nil {
		return err
	}

	if p.Version[0] != 'V' {
		return errors.New("not data")
	}

	err = binary.Read(reader, binary.BigEndian, &p.Length)
	if err != nil {
		return err
	}

	p.Data = make([]byte, p.Length)
	err = binary.Read(reader, binary.BigEndian, &p.Data)
	return err
}

func PackageScannerSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && data[0] == 'V' {
		if len(data) > 4 {
			var length int16
			binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &length)
			if int(length)+4 <= len(data) {
				return int(length) + 4, data[:int(length)+4], nil
			}
		}
	}

	return
}

func (p *Package) String() string {
	return fmt.Sprintf("version: %s length: %d data: %s",
		p.Version, p.Length, p.Data)
}

var (
	SPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 576)
		},
	} // small buff pool
	LPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 64*1024+262)
		},
	} // large buff pool for udp
)

// Transport rw1 and rw2
func Transport(rw1, rw2 io.ReadWriter) error {
	errc := make(chan error, 1)
	go func() {
		b := LPool.Get().([]byte)
		defer LPool.Put(b)

		_, err := io.CopyBuffer(rw1, rw2, b)
		errc <- err
	}()

	go func() {
		b := LPool.Get().([]byte)
		defer LPool.Put(b)

		_, err := io.CopyBuffer(rw2, rw1, b)
		errc <- err
	}()

	if err := <-errc; err != nil && err != io.EOF {
		return err
	}
	return nil
}
