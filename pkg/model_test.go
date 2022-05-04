package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestModel(t *testing.T) {
	buf := new(bytes.Buffer)

	for i := 0; i < 10; i++ {
		p := NewPackage(PData, []byte(fmt.Sprintf("hello world: %d", i)))
		err := p.Pack(buf)
		if err != nil {
			panic(err)
		}
	}

	scanner := bufio.NewScanner(buf)
	scanner.Split(PackageScannerSplit)
	for scanner.Scan() {
		scannedPack := new(Package)
		scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		log.Println(scannedPack)
	}

	if err := scanner.Err(); err != nil {
		log.Println("EOF")
	}
}
