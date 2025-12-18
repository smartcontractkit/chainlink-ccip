package changesets

import "github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"

func CreateNodeKeysBundle(
	nodes []*clclient.ChainlinkClient,
	chainName string,
	chainID string,
) ([]clclient.NodeKeysBundle, error) {
	nkb := make([]clclient.NodeKeysBundle, 0)
	for _, n := range nodes {
		p2pkeys, err := n.MustReadP2PKeys()
		if err != nil {
			return nil, err
		}

		peerID := p2pkeys.Data[0].Attributes.PeerID
		txKey, _, err := n.CreateTxKey(chainName, chainID)
		if err != nil {
			return nil, err
		}
		ocrKey, _, err := n.CreateOCR2Key(chainName)
		if err != nil {
			return nil, err
		}
		nkb = append(nkb, clclient.NodeKeysBundle{
			PeerID:  peerID,
			OCR2Key: *ocrKey,
			TXKey:   *txKey,
		})
	}
	return nkb, nil
}
