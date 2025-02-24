package lbtc_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sel "github.com/smartcontractkit/chain-selectors"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/lbtc"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type lbtcMessage struct {
	// payloadHash is the array of hex string which represents hashes of LBTC payload to request attestation
	payloadHash []string
	// attestationResponse is the response from the Attestation API
	attestationResponse string
	// attestationResponseStatus is the status code of the response from the Attestation API
	attestationResponseStatus int
}

func (u *lbtcMessage) tokenData(idx int) []byte {
	var result map[string]interface{}

	err := json.Unmarshal([]byte(u.attestationResponse), &result)
	if err != nil {
		panic(err)
	}

	attestations, ok := result["attestations"].([]interface{})
	if !ok {
		panic("attestations not found")
	}
	attestation, ok := attestations[idx].(map[string]interface{})
	if !ok {
		panic("attestation not found")
	}
	bytes, err := cciptypes.NewBytesFromString(attestation["attestation"].(string))
	if err != nil {
		panic(err)
	}
	return bytes
}

//nolint:lll
var (
	// binance_smart_chain-mainnet -> ethereum-mainnet-base-1
	// https://bscscan.com/tx/0x5ccf7f9737673e6e044147d50848686ff279b629094a4ea711e3b7ceb617197e
	// https://mainnet.prod.lombard.finance/api/bridge/v1/deposits/getByHash -> {"messageHash": ["0x117f49bfccd85ce2d0ad3a2c9bc27af2abd43eed0cbaeb2ddf5098cbd6bb8bcf"]}
	m1 = lbtcMessage{
		payloadHash: []string{"0x117f49bfccd85ce2d0ad3a2c9bc27af2abd43eed0cbaeb2ddf5098cbd6bb8bcf"},
		attestationResponse: `{
			"attestations": [{
				"message_hash":"0x117f49bfccd85ce2d0ad3a2c9bc27af2abd43eed0cbaeb2ddf5098cbd6bb8bcf",
				"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000038000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b000000000000000000000000097bcc72a1d3d09c13c6fcb489ad1f8776d9bacc000000000000000000000000000000000000000000000000000000000002848800000000000000000000000000000000000000000000000000000000000003da0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000040d9933ebfc1d76440c89ef681353707ed74d83e01cd1742df5fbe16cef03f63f22ce03b18e5656cd068227907e1dfdb6ad661b3398e6b8c361f81f505fa3587200000000000000000000000000000000000000000000000000000000000000040a3b634e583ac7ae8d9efc45f2543286ef5255b52a34583ab3cf176c9566ddd267da002cb471638c9944acaef22e79d776b72947c3d856d179dd7b5c6cbbe94350000000000000000000000000000000000000000000000000000000000000040f21186c33eb7e81bd8823ba9456d1f433493da9a490b63011320f2af3f4369fd0c3e7d70cace00ced689ceda62d22e060d88fb9b2aa6a3df8937313b5f7cbfed00000000000000000000000000000000000000000000000000000000000000404df4ffedfeccf2939bed8b1155e3f396b366f02fec4f99c127df768be9cbb7cc09333ec3a0c174050f5902faa660de8132a867b2ea034c65cd35b4ffe5e92a450000000000000000000000000000000000000000000000000000000000000040ce766bb75675c315df57263790a6a7c3e65b518e837c9b6910909070757872d82eee9dcaaa23143d18e2266dc0910fb358b27a26fca545bb53a1f76ec02ce917",
				"status":"NOTARIZATION_STATUS_SESSION_APPROVED"
			}]
		}`,
		attestationResponseStatus: 200,
	}

	// binance_smart_chain-mainnet -> ethereum-mainnet-base-1
	// https://bscscan.com/tx/0x317fc5482bbd0975525bdc476d4203655286761b33bbd1f400a6674356da6f7c
	// https://mainnet.prod.lombard.finance/api/bridge/v1/deposits/getByHash -> {"messageHash": ["0x27bf6eb2920da82a6a1294ceff503733c5a46a36d6d6c56a006f8720c399574b"]}
	m2 = lbtcMessage{
		payloadHash: []string{"0xbca4f38f27d1aaec0d36ceda9990f3508f72aa44fa90371962bc23a6d7b6429d", "0x27bf6eb2920da82a6a1294ceff503733c5a46a36d6d6c56a006f8720c399574b", "0x5455ad825ac854ec2bfee200961d62ea57269bd248b782ed727ab33fd698e061"},
		attestationResponse: `{
			"attestations": [
				{
					"message_hash":"0xbca4f38f27d1aaec0d36ceda9990f3508f72aa44fa90371962bc23a6d7b6429d",
					"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000002105000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b000000000000000000000000f4dc338b1b1184f84a461f0bb2f974fc90a814560000000000000000000000000000000000000000000000000000000008f0d18000000000000000000000000000000000000000000000000000000000000000760000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000040cc1da209c50d5df4b03c9fc9d3053ccef515b648354598114c10416d3fd692462bd3ec04d11b2c09eb119c13b4e8236a3b22c60d7f2534c9cbda23e31402984c0000000000000000000000000000000000000000000000000000000000000040a92241c515ad60095a7b823266922f8daad601915b59d232f4a6d21ce0f08650671e73a4e447ae0dd5ccdd47b3b825a8f6dfb8ebf4c53c4a2192d21e56da2b1000000000000000000000000000000000000000000000000000000000000000409e598a93d1ac8672da38aa9eb072070b0edc100d97f53c8b4b642aed9173d94b07f72630b320ae7f8fb0bf7082f15f5ed92f0f0b23c8a1c0ccd9336f6debcb6e00000000000000000000000000000000000000000000000000000000000000404999635b1bd3cc9a6051bed7e5940b2345f3b88ac39d508b3fb5fc2c32961bed451b15675ea97a58c822097934981c228d234a48ed5f4a77c6e2a41196a69845000000000000000000000000000000000000000000000000000000000000004014bb3ed8be0f6bd73167cac19eb922674dc513defbbc3ead2df54d37bdf3924422a384ad0c1ef41aa8a9003fabadb5ba91c7e32de96d370061e9474836cd1532",
					"status":"NOTARIZATION_STATUS_SESSION_APPROVED"
				},
				{
					"message_hash":"0x27bf6eb2920da82a6a1294ceff503733c5a46a36d6d6c56a006f8720c399574b",
					"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000038000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000002105000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000002608f4b20ef1ad0a95dbedb1b1d3e59f1b8ca401000000000000000000000000000000000000000000000000000000000013491c00000000000000000000000000000000000000000000000000000000000003d40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000040bcc09d9e009a36730686ea34e496d648bac3b076bc05239d0e06afed14f7f71c73d29724519311a7c7c5a26407aaf6cd2297fc105b9b4b4b0cd697ea677c39ce0000000000000000000000000000000000000000000000000000000000000040e326031805bbd52604688e303d3a0b8f1a312cce7d4d6794f5b67f2e8d3890dc73e1a6c8c0ce62d39796f64e6f2f86741d35518a0b0bbfdd648cd20ace2ae4a90000000000000000000000000000000000000000000000000000000000000040adf606c947c91bdfb7e443a870a9afd49a41c7852e0ad41781f3d7b09bb9e4bf4b3916d5571c9622a90b34a2350a68d39eb5d0f5f4adef8fadf1e68c461106ad000000000000000000000000000000000000000000000000000000000000004070e824e6f8180e7aee99f39af2ff8a0f1c3a94f7902b73fedfa827518ad722ac0eb9fb934f8b2d0e76236ef0ccb64a7faa43e8ef3d940ee1918c77be8eb3263c00000000000000000000000000000000000000000000000000000000000000407ac5168c3302549e8fd98c9e576d66801b3c3c675989f76356aee511cd8017ca7c42210c01d1abf7c69acfd250a25f325c9c16597fddde000a3cf8fdd6a13182",
					"status":"NOTARIZATION_STATUS_SESSION_APPROVED"
				},
				{
					"message_hash":"0x5455ad825ac854ec2bfee200961d62ea57269bd248b782ed727ab33fd698e061",
					"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000038000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000006975a0ff0bd36fb0cb94f43c095936cfb07de9d500000000000000000000000000000000000000000000000000000000000304a800000000000000000000000000000000000000000000000000000000000003de0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000040e598e75fa7d0f4167ab451faa7813af965bcfa00dc9e3c98cbe87c2faef390b616a20d9507d7dc5c1861559265c9e2cbd331de94f6e43df2f8ff426223fbbfbd0000000000000000000000000000000000000000000000000000000000000040e8d4eef7d9e384e1e45bce1ab2c6c272940829e4e97c30d19f7eba89e59ecfdb418cf7e1aa7fd37c98f94f994c9e1bef25afc2c88b3bf04b2fc1bfbd1529cc780000000000000000000000000000000000000000000000000000000000000040dcdb9dd9dbf4313ece1eb4949e04883116225d3c5907f43d6a6ddb37ffbf3240203969641a3cbabace340315cbfa1d6954d8e63a2541dfafe58b5e17d53e7dd0000000000000000000000000000000000000000000000000000000000000004092330a9da5ddaffb34da37526172d3b7b8bac70c101a8b1dd2722541afb2504335430549a4a4f9647088b18b8700b68f03a2eced918141d30686fa0bf5ae045d000000000000000000000000000000000000000000000000000000000000004058b9f40e888281284be0b5f84e1549d91bba2e7de47b5065a0f4468e4c07f6f016493eae6239c81ae11161005b80883fa620a24405c5512710dd6da1aa40316c",
					"status":"NOTARIZATION_STATUS_SESSION_APPROVED"
				}
			]
		}`,
		attestationResponseStatus: 200,
	}

	// https://mainnet.prod.lombard.finance/api/bridge/v1/deposits/getByHash -> {"messageHash": ["0x48302ed57e4c10905ed6ae5acc432919723cc65ba0a8319b59061472515047ef", "0x82fe5d9633b805ab473e302a4c20b9a91bcfc52e41c86d12dc1416726b11dcf7", "0xfec6a8389b7a6aa0afc3d154ce7c7b60bbad11631bf4890fa29c044790bbac54", "0xfceb7130de362329c1e1402774a62da95bc2a4300004633bc45cdea6afcc7a85"]}
	m3 = lbtcMessage{
		payloadHash: []string{"0x48302ed57e4c10905ed6ae5acc432919723cc65ba0a8319b59061472515047ef", "0x82fe5d9633b805ab473e302a4c20b9a91bcfc52e41c86d12dc1416726b11dcf7", "0xfec6a8389b7a6aa0afc3d154ce7c7b60bbad11631bf4890fa29c044790bbac54", "0xfceb7130de362329c1e1402774a62da95bc2a4300004633bc45cdea6afcc7a85"},
		attestationResponse: `{
			"attestations": [
				{
					"message_hash":"0x48302ed57e4c10905ed6ae5acc432919723cc65ba0a8319b59061472515047ef",
					"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000002105000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000001b648ade1ef219c87987cd60eba069a7faf1621f0000000000000000000000000000000000000000000000000000000008583b0000000000000000000000000000000000000000000000000000000000000000950000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000040c26183b4e4752b062f7e5dbdccba71f37997c0af2ce6a28db97e6b55648ea5cb5deca115b16f7fd969135f0279abc639f8e26a4d93e4713532993f387f90ad6e00000000000000000000000000000000000000000000000000000000000000404ca4577b7936deff2aa26b9dd906eb352c35c18a280a59141e742969927cb78d1561b919675980b0eaa9fc41d1985688db89910b4e5c1576b1edf49cde88c101000000000000000000000000000000000000000000000000000000000000004028d437358f7af12b48879cfb4ff2df72bae54e3dd7d87c5722b185577e6c363068736c23ea67422670a2e8eeb19572adddfec9fd04460e7b43c7c40ecb0c533f0000000000000000000000000000000000000000000000000000000000000040cbd3d5231a50d7e98216607202c000f0665b339d045438133cb5fe56520f40f3529be0773afa3aceb0fdaa99bce1818ec0e5335ae3f421663c4fff07baea701a0000000000000000000000000000000000000000000000000000000000000040fdd02f81bd45951a1344d84a1f29f8269aead3f637f450384b4aa7960d065e015d1e4a288097b11cb7103605d5064f081ec0be5fda48074c1f280fae3db3a946",
					"status":"NOTARIZATION_STATUS_SESSION_APPROVED"
				},
				{
					"message_hash":"0x82fe5d9633b805ab473e302a4c20b9a91bcfc52e41c86d12dc1416726b11dcf7",
					"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000002105000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000002eae9ac17edc1966ff12003afecdd1cad1a7de30000000000000000000000000000000000000000000000000000000000098fbac00000000000000000000000000000000000000000000000000000000000000960000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000040c985a727df981d734b86007f33531620e93e4bdb30f260f7142c643e2f20da5264ecbfa5fb1fb3d086d1c1b2fd5fa93e082d2788e58fa5c8466f7b035f6a138a00000000000000000000000000000000000000000000000000000000000000403c07e222472606fc3400bca7581b08d464007b54519c8bc729a7b129f3da0eae4612670a96bc8d47ec43bfb0107a36e80cde024733b0ded58a3ce3c0ff27ecfb000000000000000000000000000000000000000000000000000000000000004022b155feb1a193cfbfe77f7df605ea890bbd3eea8a72869c97ef0f156d3b99676d37557fed35eca2b8c52590f1df4b37f58dfc0c75dbddf18df09b2e6565503500000000000000000000000000000000000000000000000000000000000000405c6f3dbfcb1e21647f037410e00871068e56730bd27bf429830bf2cb065a8c7c1b5e3e9bcc33fb80aae7eeead1922dfb86ed48259503c6ac7b2d078c2427a2f200000000000000000000000000000000000000000000000000000000000000409f5f36668b844b7f9e4ef404c3def915ceeaf16a93079fe2b35592deee2d120f45c423ffc0ab07fd8be3390e1f7440c5cb60f37968e7d4742eec12a627c36c56",
					"status":"NOTARIZATION_STATUS_SESSION_APPROVED"
				},
				{
					"message_hash":"0xfec6a8389b7a6aa0afc3d154ce7c7b60bbad11631bf4890fa29c044790bbac54",
					"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000002105000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b00000000000000000000000086867e7552417f4e31ee72946116e196617524440000000000000000000000000000000000000000000000000000000004f5b55000000000000000000000000000000000000000000000000000000000000000970000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000000404ef41420152776c1385916532647ee1db973a9f3fb5a3af2fd4fec1a709e383170f8a940c70e5aacb22dc75522386963bbdc7d650dea11219c7a163f84322317000000000000000000000000000000000000000000000000000000000000004030bd23fae9ca9eda68d8b724c3a3c326c2aa86ca37cf84dab9b88b4e6a14acaa49ce5f9fc5b3c7222992a17d2afe9aba23863b4a583f2b998adeef0c618e47410000000000000000000000000000000000000000000000000000000000000040e208fdb10d80fbff04b93ff6896c049e9350e0bcc077a0cb627120c957e9632f5cd1bcabc50bef567439c7fcfe14119e4a1fea5b433b2a919f08da2d3247804100000000000000000000000000000000000000000000000000000000000000406eb0bea43ea23056f074152068cebca78430d9906664eb44745a234f11ceb74e3c33d6e4ae251af44b0444cfe08ad15a54627d600c941dc6dc0de61a2d18234d00000000000000000000000000000000000000000000000000000000000000409e9bf6bcff0a09a8d0c1bdbde9af77ca58513f3c70211c6f6571d64326054c180e9e067467dd1510bf2fbef0412debeeb1c6a3113c8b84ca8558ef810d0fd4e8",
					"status":"NOTARIZATION_STATUS_SESSION_APPROVED"
				},
				{
					"message_hash":"0xfceb7130de362329c1e1402774a62da95bc2a4300004633bc45cdea6afcc7a85",
					"status":"NOTARIZATION_STATUS_SESSION_APPROVED",
					"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000002105000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000001984e3b342f1bc8476557f1686e73b24b76799ed0000000000000000000000000000000000000000000000000000000000009c4000000000000000000000000000000000000000000000000000000000000000980000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000040b7d67457824915a99c624cdff679558fc0f958eacc59423b2204f4ec91222c5b3938ab6b8848150e923b071c77b3e71103387400d6f41faab5a6b553a11b71a80000000000000000000000000000000000000000000000000000000000000040e0b35545caf0787c5c4917f6eb47e37c78b1875d0eeaa7163488bc01e6e74df826eead971a9ee2c31eb75efe52de71edb8b98aea0da840ac4d81a719fc0c090b00000000000000000000000000000000000000000000000000000000000000403a2246d825a67baf616df4613d04ed388253886a7f22785c0ae5a385965678d30095b66320a96d76988de9b7934aec95b652c0583a77db3937fc4df2190eb1ea0000000000000000000000000000000000000000000000000000000000000040de87d0b8dc4c2fd0ddd66a49171eb6cad344025b67897042f8330fb7c5756c6922a59af84245e50728e0f17abc06ef86af784252d0ddf65d684723eda72ee858000000000000000000000000000000000000000000000000000000000000004059b9328a6efa95c7c969a67bb7fec3fe0e08e7522b1cf4443b230d9de7aa2495301fc48243aae0a66cf6e42cd0f681e3245a56056529b577ccfd381ee7699da7"
				}
			]
		}`,
		attestationResponseStatus: 200,
	}

	// These are all approved, but testing all possible statuses
	// https://mainnet.prod.lombard.finance/api/bridge/v1/deposits/getByHash -> {"messageHash": ["0x74fd1a76d356b1e331e7fc586c27f52cb22f5207962f562f566eeeecfb2e6454", "0xfd9574739cbc2f206226ce545b5ee9c9022ac9828add7667d3a0506e3cabac3b", "0x6567a3bec9881c1101e90224cb5f6b0f13d1f2923d3c310c4acde91accd3f474", "0x33dca299c74ca51a70936eab2668a549cacf5c99be556021577b3f3b77b7b850", "0x1531c52169ad7aa7edeb520aac02c91d272aff9ed2099a3a7638a805db62b3bb"]}
	m4 = lbtcMessage{
		payloadHash: []string{"0x74fd1a76d356b1e331e7fc586c27f52cb22f5207962f562f566eeeecfb2e6454", "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "0xfd9574739cbc2f206226ce545b5ee9c9022ac9828add7667d3a0506e3cabac3b", "0x6567a3bec9881c1101e90224cb5f6b0f13d1f2923d3c310c4acde91accd3f474", "0x33dca299c74ca51a70936eab2668a549cacf5c99be556021577b3f3b77b7b850", "0x1531c52169ad7aa7edeb520aac02c91d272aff9ed2099a3a7638a805db62b3bb"},
		attestationResponse: `{
			"attestations": [
				{
					"message_hash":"0x74fd1a76d356b1e331e7fc586c27f52cb22f5207962f562f566eeeecfb2e6454",
					"attestation":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000000000000000000000000000000000000002105000000000000000000000000a869817b48b25eee986bdf4be04062e6fd2c418b0000000000000000000000000a05c538da9e321ad07f6f18424fae7fc4fa25db0000000000000000000000000000000000000000000000000000000014dc938000000000000000000000000000000000000000000000000000000000000000990000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002200000000000000000000000000000000000000000000000000000000000000040dfcb983a714595ac5ce52704a202da95e09dbef64a2fb53e12f4cb1377c521ed0e447278ac8cbfb21abaad60dcb829a27ee61b0c803770846b16894a0bce3daa000000000000000000000000000000000000000000000000000000000000004021585159861085ebf4a52f352872df047533cb286f44497d7adec228349c12556fa55cf8f127f190666d2884b8b3c1c1baff71cbdd6eccb3703146421325e45d00000000000000000000000000000000000000000000000000000000000000404628962fca1e75619230d056e5054050d39beb8ae407b0282391e92302257b9f524cec0b6acc7166efe58e3e693279d51e9c49ac1aeaa4ce29e760e3ed602b2a0000000000000000000000000000000000000000000000000000000000000040b2658a5da29fa33599bd555601b93a91900cdfd715b9d17f2ce123e0f544338771d772d6498e8db689d63239875b1b2f2f502d69d71a1700886a458c15e8dfd60000000000000000000000000000000000000000000000000000000000000040e8f915415fde22ad637c498d07fe33d2b577e9292cca9917fbad21ec4be0477b51a40571cbe3bdf6d8a375bee7e0368dc1dabec6062094cf0440e45b28e600f4",
					"status":"NOTARIZATION_STATUS_SESSION_APPROVED"
				},
				{
					"message_hash":"0xfd9574739cbc2f206226ce545b5ee9c9022ac9828add7667d3a0506e3cabac3b",
					"status":"NOTARIZATION_STATUS_PENDING"
				},
				{
					"message_hash":"0x6567a3bec9881c1101e90224cb5f6b0f13d1f2923d3c310c4acde91accd3f474",
					"status":"NOTARIZATION_STATUS_SUBMITTED"
				},
				{
					"message_hash":"0x33dca299c74ca51a70936eab2668a549cacf5c99be556021577b3f3b77b7b850",
					"status":"NOTARIZATION_STATUS_FAILED"
				},
				{
					"message_hash":"0x1531c52169ad7aa7edeb520aac02c91d272aff9ed2099a3a7638a805db62b3bb",
					"status":"NOTARIZATION_STATUS_UNSPECIFIED"
				}
			]
		}`,
		attestationResponseStatus: 200,
	}

	// https://mainnet.prod.lombard.finance/api/bridge/v1/deposits/getByHash -> {"messageHash": ["0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"]}
	m5 = lbtcMessage{
		payloadHash: []string{"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		attestationResponse: `{
		    "code": 13,
		    "message": "failed to get deposits by hash set: rpc error: code = InvalidArgument desc = invalid hash"
		}`,
		attestationResponseStatus: 500,
	}
)

// This test focuses on almost e2e flows for USDC message
// It aims to make it as real as possible by using real payloads from the chain and
// real responses from the Attestation API. As long as the Attestation API is supporting old events
// you should be able to just copy-paste block explorer and attestation requests to the browser
// and see those responses in action
func Test_LBTC_Flow(t *testing.T) {
	bscPool := internal.RandBytes().String()
	ethPool := internal.RandBytes().String()
	bscChain := cciptypes.ChainSelector(sel.BINANCE_SMART_CHAIN_MAINNET.Selector)
	ethChain := cciptypes.ChainSelector(sel.ETHEREUM_MAINNET.Selector)

	httpMessages := []lbtcMessage{m1, m2, m3, m4, m5}

	// Mock http server to return proper payloads
	server := mockHTTPServerResponse(t, httpMessages)
	defer server.Close()

	config := pluginconfig.LBTCObserverConfig{
		AttestationConfig: pluginconfig.AttestationConfig{
			AttestationAPI:         server.URL,
			AttestationAPIInterval: commonconfig.MustNewDuration(1 * time.Microsecond),
			AttestationAPITimeout:  commonconfig.MustNewDuration(1 * time.Second),
			AttestationAPICooldown: commonconfig.MustNewDuration(5 * time.Minute),
		},
		SourcePoolAddressByChain: map[cciptypes.ChainSelector]string{
			bscChain: bscPool,
			ethChain: ethPool,
		},
	}

	baseObserver, err := lbtc.NewLBTCTokenDataObserver(
		logger.Test(t),
		cciptypes.ChainSelector(sel.ETHEREUM_MAINNET_BASE_1.Selector),
		config,
	)
	require.NoError(t, err)

	tt := []struct {
		name     string
		messages exectypes.MessageObservations
		want     exectypes.TokenDataObservations
	}{
		{
			name: "single valid message from bsc to base",
			messages: exectypes.MessageObservations{
				bscChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(bscPool, m1.payloadHash[0]),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				bscChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m1.tokenData(0)),
						},
					},
				},
			},
		},
		{
			name: "multiple valid messages and tokens from bsc to base",
			messages: exectypes.MessageObservations{
				bscChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(bscPool, m2.payloadHash[0]),
							createToken(bscPool, m2.payloadHash[1]),
						},
					},
					2: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(bscPool, m2.payloadHash[2]),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				bscChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m2.tokenData(0)),
							exectypes.NewSuccessTokenData(m2.tokenData(1)),
						},
					},
					2: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m2.tokenData(2)),
						},
					},
				},
			},
		},
		{
			name: "multiple source chain tokens",
			messages: exectypes.MessageObservations{
				bscChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(bscPool, m3.payloadHash[0]),
							createToken(bscPool, m3.payloadHash[1]),
						},
					},
				},
				ethChain: {
					10: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(ethPool, m3.payloadHash[2]),
						},
					},
					11: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(ethPool, m3.payloadHash[3]),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				bscChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m3.tokenData(0)),
							exectypes.NewSuccessTokenData(m3.tokenData(1)),
						},
					},
				},
				ethChain: {
					10: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m3.tokenData(2)),
						},
					},
					11: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m3.tokenData(3)),
						},
					},
				},
			},
		},
		{
			name: "messages with tokens failing to get ready",
			messages: exectypes.MessageObservations{
				bscChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(bscPool, m4.payloadHash[0]),
						},
					},
					2: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							// attestation missing
							createToken(bscPool, m4.payloadHash[1]),
						},
					},
					3: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							// attestation pending
							createToken(bscPool, m4.payloadHash[2]),
						},
					},
				},
				ethChain: {
					10: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							// attestation submitted
							createToken(ethPool, m4.payloadHash[3]),
						},
					},
					11: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							// attestation failed
							createToken(ethPool, m4.payloadHash[4]),
						},
					},
					12: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							// attestation unspecified
							createToken(ethPool, m4.payloadHash[5]),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				bscChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m4.tokenData(0)),
						},
					},
					2: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(tokendata.ErrDataMissing),
						},
					},
					3: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(tokendata.ErrNotReady),
						},
					},
				},
				ethChain: {
					10: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(tokendata.ErrNotReady),
						},
					},
					11: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(tokendata.ErrUnknownResponse),
						},
					},
					12: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(tokendata.ErrUnknownResponse),
						},
					},
				},
			},
		},
		{
			name: "attestation api internal server error",
			messages: exectypes.MessageObservations{
				bscChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(bscPool, m5.payloadHash[0]),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				bscChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(tokendata.ErrUnknownResponse),
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err1 := baseObserver.Observe(context.Background(), tc.messages)
			require.NoError(t, err1)
			require.Equal(t, tc.want, got)
		})

	}
}

func createToken(pool string, payloadHash string) cciptypes.RampTokenAmount {
	sourcePoolAddress, err := cciptypes.NewUnknownAddressFromHex(pool)
	if err != nil {
		panic(err)
	}
	extraData := internal.MustDecode(payloadHash)
	return cciptypes.RampTokenAmount{
		SourcePoolAddress: sourcePoolAddress,
		ExtraData:         extraData,
		Amount:            cciptypes.NewBigIntFromInt64(100),
	}
}

func mockHTTPServerResponse(t *testing.T, messages []lbtcMessage) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyRaw, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		var body map[string][]string
		err = json.Unmarshal(bodyRaw, &body)
		require.NoError(t, err)
		payloadHashes, ok := body["messageHash"]
		require.True(t, ok)
		for _, message := range messages {
			matches := true
			for _, payloadHash := range message.payloadHash {
				matches = matches && slices.Contains(payloadHashes, payloadHash)
			}
			if matches {
				w.WriteHeader(message.attestationResponseStatus)
				_, err = w.Write([]byte(message.attestationResponse))
				require.NoError(t, err)
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
	}))
	return ts
}
