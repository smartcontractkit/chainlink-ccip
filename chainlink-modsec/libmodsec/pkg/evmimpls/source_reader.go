package evmimpls

import (
	"context"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

var _ modsectypes.SourceReader = (*EVMSourceReader)(nil)

type EVMSourceReader struct {
	logSubscriber      ethereum.LogFilterer
	messageChan        chan modsectypes.Message
	logsChan           chan types.Log
	onRampProxyAddress common.Address
	eventSig           common.Hash
	logger             *log.Logger
	evmEventCodec      EVMEventCodec
	wg                 sync.WaitGroup
	runCtxCancel       context.CancelFunc
}

func NewEVMSourceReader(
	logSubscriber ethereum.LogFilterer,
	onRampProxyAddress common.Address,
	eventSig common.Hash,
	logger *log.Logger,
	evmEventCodec EVMEventCodec,
) *EVMSourceReader {
	return &EVMSourceReader{
		logSubscriber:      logSubscriber,
		messageChan:        make(chan modsectypes.Message),
		logsChan:           make(chan types.Log),
		onRampProxyAddress: onRampProxyAddress,
		eventSig:           eventSig,
		logger:             logger,
		evmEventCodec:      evmEventCodec,
		wg:                 sync.WaitGroup{},
	}
}

// Close implements modsectypes.SourceReader.
func (r *EVMSourceReader) Close() error {
	r.runCtxCancel()
	r.wg.Wait()
	close(r.messageChan)
	return nil
}

// GetMessages implements modsectypes.SourceReader.
func (r *EVMSourceReader) GetMessages(ctx context.Context, query modsectypes.MessageQueryArgs) ([]modsectypes.Message, error) {
	panic("unimplemented - please use Messages() instead")
}

// Messages implements modsectypes.SourceReader.
func (r *EVMSourceReader) Messages() <-chan modsectypes.Message {
	return r.messageChan
}

func (r *EVMSourceReader) Start(ctx context.Context) error {
	r.wg.Add(1)
	runCtx, runCtxCancel := context.WithCancel(context.Background())
	go func() {
		defer r.wg.Done()
		sub, err := r.logSubscriber.SubscribeFilterLogs(runCtx, ethereum.FilterQuery{
			Addresses: []common.Address{r.onRampProxyAddress},
			Topics:    [][]common.Hash{{r.eventSig}},
		}, r.logsChan)
		if err != nil {
			r.logger.Printf("failed to subscribe to logs: %v\n", err)
			return
		}
		defer sub.Unsubscribe()

		for {
			select {
			case <-runCtx.Done():
				return
			case log := <-r.logsChan:
				r.logger.Printf("received log: %v\n", log)
				message, err := r.evmEventCodec.Decode(ctx, log.Data)
				if err != nil {
					r.logger.Printf("failed to decode message: %v\n", err)
					continue
				}
				r.messageChan <- message
			}
		}
	}()
	r.runCtxCancel = runCtxCancel
	return nil
}
