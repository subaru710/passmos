package passport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

const (
	IPFS_STORE_TYPE = "ipfs"
)

type IpfsStore struct {
	Hashes []string
	sh     *shell.Shell
}

func NewIpfsStore(url string) *IpfsStore {
	sh := shell.NewShell(url)
	return &IpfsStore{Hashes: []string{}, sh: sh}
}

func (store IpfsStore) Type() string {
	return IPFS_STORE_TYPE
}

func (store IpfsStore) GetPersonalData() *PersonalData {
	if len(store.Hashes) == 0 {
		return nil
	}
	hash := store.Hashes[len(store.Hashes)-1]
	rc, err := store.sh.Cat(hash)
	if err != nil {
		fmt.Println("GetPersonalData error:", err)
		return nil
	}
	pdBytes, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Println("GetPersonalData error:", err)
		return nil
	}
	pd := &PersonalData{}
	json.Unmarshal(pdBytes, pd)
	return pd
}

func (store IpfsStore) SetPersonalData(pd PersonalData) (string, error) {
	pdBytes, _ := json.Marshal(pd)
	return store.sh.Add(bytes.NewReader(pdBytes))
}

func (store IpfsStore) HasPersonalData() bool {
	return len(store.Hashes) > 0
}
