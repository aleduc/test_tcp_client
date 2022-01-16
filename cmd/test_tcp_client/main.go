package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
	"net"
	"time"
)

const (
	serverAddress = "tcp_server:2701"
	zeroes        = 20
)

var (
	clientAddress = []byte("172.28.5.5")
)

func main() {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(Work(zeroes))
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf[:]))
}

func PrepareHashString(unixTime []byte, rndBytes []byte, nonce uint32) (hash, nonceByte []byte) {
	bNonce := make([]byte, 12)
	binary.LittleEndian.PutUint32(bNonce, nonce)
	return bytes.Join(
		[][]byte{
			clientAddress,
			unixTime,
			rndBytes,
			bNonce,
		},
		[]byte{},
	), bNonce
}

func Work(zeroes int) []byte {
	var hashInt big.Int
	target := big.NewInt(1)
	target.Lsh(target, uint(256-zeroes))
	nonce := uint32(0)
	bRandom := make([]byte, 12)
	_, _ = rand.Read(bRandom)
	bUnixTime := make([]byte, 8)
	binary.LittleEndian.PutUint64(bUnixTime, uint64(time.Now().Unix()))
	var (
		digest       [32]byte
		bNonce       []byte
		bytesForHash []byte
	)
	for {
		nonce++
		bytesForHash, bNonce = PrepareHashString(bUnixTime, bRandom, nonce)
		digest = sha256.Sum256(bytesForHash)
		hashInt.SetBytes(digest[:])

		if hashInt.Cmp(target) == -1 {
			break
		} else {
			nonce++
		}

	}

	return bytes.Join(
		[][]byte{
			bUnixTime,
			bRandom,
			bNonce,
			digest[:],
		},
		[]byte{},
	)
}
