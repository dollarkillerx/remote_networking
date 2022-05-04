package pkg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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
