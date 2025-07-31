export type TestEventEmitter = {
  "version": "0.1.1-dev",
  "name": "test_event_emitter",
  "instructions": [
    {
      "name": "emitCcipCctpMsgSent",
      "accounts": [
        {
          "name": "clock",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "args",
          "type": {
            "defined": "CcipCctpMessageSentEventArgs"
          }
        }
      ]
    }
  ],
  "types": [
    {
      "name": "CcipCctpMessageSentEventArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "messageSentBytes",
            "type": "bytes"
          },
          {
            "name": "remoteChainSelector",
            "type": "u64"
          },
          {
            "name": "originalSender",
            "type": "publicKey"
          },
          {
            "name": "eventAddress",
            "type": "publicKey"
          },
          {
            "name": "msgTotalNonce",
            "type": "u64"
          },
          {
            "name": "sourceDomain",
            "type": "u32"
          },
          {
            "name": "cctpNonce",
            "type": "u64"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "CcipCctpMessageSentEvent",
      "fields": [
        {
          "name": "originalSender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "remoteChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "msgTotalNonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "eventAddress",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "sourceDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "cctpNonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "messageSentBytes",
          "type": "bytes",
          "index": false
        }
      ]
    }
  ]
};

export const IDL: TestEventEmitter = {
  "version": "0.1.1-dev",
  "name": "test_event_emitter",
  "instructions": [
    {
      "name": "emitCcipCctpMsgSent",
      "accounts": [
        {
          "name": "clock",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "args",
          "type": {
            "defined": "CcipCctpMessageSentEventArgs"
          }
        }
      ]
    }
  ],
  "types": [
    {
      "name": "CcipCctpMessageSentEventArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "messageSentBytes",
            "type": "bytes"
          },
          {
            "name": "remoteChainSelector",
            "type": "u64"
          },
          {
            "name": "originalSender",
            "type": "publicKey"
          },
          {
            "name": "eventAddress",
            "type": "publicKey"
          },
          {
            "name": "msgTotalNonce",
            "type": "u64"
          },
          {
            "name": "sourceDomain",
            "type": "u32"
          },
          {
            "name": "cctpNonce",
            "type": "u64"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "CcipCctpMessageSentEvent",
      "fields": [
        {
          "name": "originalSender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "remoteChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "msgTotalNonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "eventAddress",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "sourceDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "cctpNonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "messageSentBytes",
          "type": "bytes",
          "index": false
        }
      ]
    }
  ]
};
