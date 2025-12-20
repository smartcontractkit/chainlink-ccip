package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

const peerIDPrefix = "p2p_"

type PeerID libocrtypes.PeerID

func MakePeerID(s string) (PeerID, error) {
	var peerID PeerID
	return peerID, peerID.UnmarshalString(s)
}

func (p PeerID) String() string {
	// Handle a zero peerID more gracefully, i.e. print it as empty string rather
	// than `p2p_`
	if p == (PeerID{}) {
		return ""
	}
	return fmt.Sprintf("%s%s", peerIDPrefix, p.Raw())
}

func (p PeerID) Raw() string {
	return libocrtypes.PeerID(p).String()
}

func (p *PeerID) UnmarshalString(s string) error {
	return p.UnmarshalText([]byte(s))
}

func (p *PeerID) MarshalText() ([]byte, error) {
	if *p == (PeerID{}) {
		return nil, nil
	}
	return []byte(p.Raw()), nil
}

func (p *PeerID) UnmarshalText(bs []byte) error {
	input := string(bs)
	if strings.HasPrefix(input, peerIDPrefix) {
		input = string(bs[len(peerIDPrefix):])
	}

	if input == "" {
		return nil
	}

	var peerID libocrtypes.PeerID
	err := peerID.UnmarshalText([]byte(input))
	if err != nil {
		return errors.New(fmt.Sprintf(`PeerID#UnmarshalText("%v"): %v`, input, err))
	}
	*p = PeerID(peerID)
	return nil
}

func (p *PeerID) Scan(value any) error {
	*p = PeerID{}
	switch s := value.(type) {
	case string:
		if s != "" {
			return p.UnmarshalText([]byte(s))
		}
	case nil:
	default:
		return errors.New("incompatible type for PeerID scan")
	}
	return nil
}

func (p PeerID) Value() (driver.Value, error) {
	b, err := libocrtypes.PeerID(p).MarshalText()
	return string(b), err
}

func (p PeerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PeerID) UnmarshalJSON(input []byte) error {
	var result string
	if err := json.Unmarshal(input, &result); err != nil {
		return err
	}

	return p.UnmarshalText([]byte(result))
}

func MustPeerIDFromString(s string) PeerID {
	p := PeerID{}
	if err := p.UnmarshalString(s); err != nil {
		panic(err)
	}
	return p
}

func CreateNodeKeysBundle(
	nodes []*clclient.ChainlinkClient,
	chainName string,
	chainID string,
) (map[string]clclient.NodeKeysBundle, error) {
	nkb := make(map[string]clclient.NodeKeysBundle)
	for _, n := range nodes {
		p2pkeys, err := n.MustReadP2PKeys()
		if err != nil {
			return nil, err
		}
		existingOCR2Keys, err := n.MustReadOCR2Keys()
		if err != nil {
			return nil, err
		}
		var txKey *clclient.TxKey
		var ocrKey *clclient.OCR2Key
		for _, key := range existingOCR2Keys.Data {
			if key.Attributes.ChainType == chainName {
				fmt.Println("OCR2 key already exists for node:", key.Attributes.ChainType)
				ocrKey = &clclient.OCR2Key{
					Data: key,
				}
				break
			}
		}
		existingTxnKeys, _, err := n.ReadTxKeys(chainName)
		if err != nil {
			return nil, err
		}
		if len(existingTxnKeys.Data) > 1 {
			return nil, fmt.Errorf("more than one txn key for chain %s on node %s", chainName, n.Config.URL)
		} else if len(existingTxnKeys.Data) == 1 {
			txKey = &clclient.TxKey{
				Data: existingTxnKeys.Data[0],
			}
		}

		peerID := p2pkeys.Data[0].Attributes.PeerID
		if txKey == nil {
			txKey, _, err = n.CreateTxKey(chainName, chainID)
			if err != nil {
				return nil, err
			}
		}
		if ocrKey == nil {
			ocrKey, _, err = n.CreateOCR2Key(chainName)
			if err != nil {
				return nil, err
			}
		}
		nkb[peerID] = clclient.NodeKeysBundle{
			PeerID:  peerID,
			OCR2Key: *ocrKey,
			TXKey:   *txKey,
		}
	}
	return nkb, nil
}
