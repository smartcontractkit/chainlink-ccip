package common

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestParseLogMessages(t *testing.T) {
	// basic log parsing tests(w/o events)
	tests := []struct {
		name     string
		logs     []string
		expected []*AnchorInstruction
	}{
		{
			name: "Test Case 1 - Empty Instruction",
			logs: []string{
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX invoke [1]",
				"Program log: Instruction: Execute",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ invoke [2]",
				"Program log: Instruction: Empty",
				"Program log: Called `empty` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps }",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ consumed 13620 of 180083 compute units",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ success",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX consumed 35400 of 200000 compute units",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX success",
			},
			expected: []*AnchorInstruction{
				{
					Name:      "Execute",
					ProgramID: "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:      []string{},
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "Empty",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `empty` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps }",
							},
							InnerCalls: []*AnchorInstruction{},
						},
					},
				},
			},
		},
		{
			name: "Test Case 2 - U8InstructionData",
			logs: []string{
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX invoke [1]",
				"Program log: Instruction: Execute",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ invoke [2]",
				"Program log: Instruction: U8InstructionData",
				"Program log: Called `u8_instruction_data` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data 123",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ consumed 13648 of 180048 compute units",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ success",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX consumed 35463 of 200000 compute units",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX success",
			},
			expected: []*AnchorInstruction{
				{
					Name:      "Execute",
					ProgramID: "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:      []string{},
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "U8InstructionData",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `u8_instruction_data` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data 123",
							},
							InnerCalls: []*AnchorInstruction{},
						},
					},
				},
			},
		},
		{
			name: "Test Case 3 - StructInstructionData",
			logs: []string{
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX invoke [1]",
				"Program log: Instruction: Execute",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ invoke [2]",
				"Program log: Instruction: StructInstructionData",
				"Program log: Called `struct_instruction_data` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data Value { value: 234 }",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ consumed 13920 of 180631 compute units",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ success",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX consumed 35152 of 200000 compute units",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX success",
			},
			expected: []*AnchorInstruction{
				{
					Name:      "Execute",
					ProgramID: "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:      []string{},
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "StructInstructionData",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `struct_instruction_data` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data Value { value: 234 }",
							},
							InnerCalls: []*AnchorInstruction{},
						},
					},
				},
			},
		},
		{
			name: "Test Case 4 - AccountRead",
			logs: []string{
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX invoke [1]",
				"Program log: Instruction: Execute",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ invoke [2]",
				"Program log: Instruction: AccountRead",
				"Program log: Called `account_read` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: AccountRead { u8_value: Account { account: Value { value: 1 }, info: AccountInfo { key: 8WGXBpVJrBATopzT8iXvRuvp5f3U63uB13tfQjGoi6rM, owner: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, is_signer: false, is_writable: false, executable: false, rent_epoch: 18446744073709551615, lamports: 953520, data.len: 9, data: 879ef47548cb18c201, .. } } }, remaining_accounts: [], bumps: AccountReadBumps { u8_value: 255 } }",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ consumed 45559 of 177765 compute units",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ success",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX consumed 69682 of 200000 compute units",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX success",
			},
			expected: []*AnchorInstruction{
				{
					Name:      "Execute",
					ProgramID: "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:      []string{},
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "AccountRead",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `account_read` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: AccountRead { u8_value: Account { account: Value { value: 1 }, info: AccountInfo { key: 8WGXBpVJrBATopzT8iXvRuvp5f3U63uB13tfQjGoi6rM, owner: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, is_signer: false, is_writable: false, executable: false, rent_epoch: 18446744073709551615, lamports: 953520, data.len: 9, data: 879ef47548cb18c201, .. } } }, remaining_accounts: [], bumps: AccountReadBumps { u8_value: 255 } }",
							},
							InnerCalls: []*AnchorInstruction{},
						},
					},
				},
			},
		},
		{
			name: "Test Case 5 - AccountMut",
			logs: []string{
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX invoke [1]",
				"Program log: Instruction: Execute",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ invoke [2]",
				"Program log: Instruction: AccountMut",
				"Program log: Called `account_mut` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: AccountMut { u8_value: Account { account: Value { value: 1 }, info: AccountInfo { key: 8WGXBpVJrBATopzT8iXvRuvp5f3U63uB13tfQjGoi6rM, owner: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, is_signer: false, is_writable: true, executable: false, rent_epoch: 18446744073709551615, lamports: 953520, data.len: 9, data: 879ef47548cb18c201, .. } }, stub_caller: Signer { info: AccountInfo { key: BUx7YZMoVXCnT2BewMZc2hr8yxoiihtHdDuoa19D9R5q, owner: 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX, is_signer: true, is_writable: true, executable: false, rent_epoch: 18446744073709551615, lamports: 2874480, data.len: 285, data: 2c3eace1f603b2211266a21317e30848020102010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000, .. } }, system_program: Program { info: AccountInfo { key: 11111111111111111111111111111111, owner: NativeLoader1111111111111111111111111111111, is_signer: false, is_writable: false, executable: true, rent_epoch: 0, lamports: 1, data.len: 14, data: 73797374656d5f70726f6772616d, .. } } }, remaining_accounts: [], bumps: AccountMutBumps { u8_value: 255 } }",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ consumed 111015 of 173365 compute units",
				"Program 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ success",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX consumed 139571 of 200000 compute units",
				"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX success",
			},
			expected: []*AnchorInstruction{
				{
					Name:      "Execute",
					ProgramID: "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:      []string{},
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "AccountMut",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `account_mut` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: AccountMut { u8_value: Account { account: Value { value: 1 }, info: AccountInfo { key: 8WGXBpVJrBATopzT8iXvRuvp5f3U63uB13tfQjGoi6rM, owner: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, is_signer: false, is_writable: true, executable: false, rent_epoch: 18446744073709551615, lamports: 953520, data.len: 9, data: 879ef47548cb18c201, .. } }, stub_caller: Signer { info: AccountInfo { key: BUx7YZMoVXCnT2BewMZc2hr8yxoiihtHdDuoa19D9R5q, owner: 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX, is_signer: true, is_writable: true, executable: false, rent_epoch: 18446744073709551615, lamports: 2874480, data.len: 285, data: 2c3eace1f603b2211266a21317e30848020102010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000, .. } }, system_program: Program { info: AccountInfo { key: 11111111111111111111111111111111, owner: NativeLoader1111111111111111111111111111111, is_signer: false, is_writable: false, executable: true, rent_epoch: 0, lamports: 1, data.len: 14, data: 73797374656d5f70726f6772616d, .. } } }, remaining_accounts: [], bumps: AccountMutBumps { u8_value: 255 } }",
							},
							InnerCalls: []*AnchorInstruction{},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseLogMessages(tt.logs, []EventMapping{})
			require.Equal(t, len(tt.expected), len(result), "Instruction count mismatch - expected: %d, got: %d", len(tt.expected), len(result))
			for i := range tt.expected {
				assertInstructionEqual(t, tt.expected[i], result[i])
			}
		})
	}

	// NOTE: events for test, unable to import due to circular dependency
	// contracts/mcm_events
	type OpExecuted struct {
		Nonce uint64           // nonce
		To    solana.PublicKey // to
		Data  []byte           // data: Vec<u8>
	}
	// contracts/timelock_events
	type CallScheduled struct {
		ID          [32]byte         // id
		Index       uint64           // index
		Target      solana.PublicKey // target
		Predecessor [32]byte         // predecessor
		Salt        [32]byte         // salt
		Delay       uint64           // delay
		Data        []byte           // data: Vec<u8>
	}

	t.Run("should parse nested cpi events correctly", func(t *testing.T) {
		// 18 mint events
		logs := []string{
			"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX invoke [1]",
			"Program log: Instruction: Execute",
			"Program LoCoNsJFuhTkSQjfdDfn3yuwqhSYoPujmviRHVCzsqn invoke [2]",
			"Program log: Instruction: ScheduleBatch",
			"Program data: v1Vap4TfuDn+kvsVqBr/AteFVu/weaOHT6IuYqTyhWqu+Oo0N6H7vgAAAAAAAAAABt324e51j94YQl285GzN2rYa/E2DuQ0n/r35KNihi/w0fAqd5gQfkYIohsaMAW0Tz4R1O58t9Hes5QbmGAti3wAAAZPU+0A5PuT/0+5So7QAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAKAAAADADodkgXAAAACQ==",
			"Program data: v1Vap4TfuDn+kvsVqBr/AteFVu/weaOHT6IuYqTyhWqu+Oo0N6H7vgEAAAAAAAAABt324e51j94YQl285GzN2rYa/E2DuQ0n/r35KNihi/w0fAqd5gQfkYIohsaMAW0Tz4R1O58t9Hes5QbmGAti3wAAAZPU+0A5PuT/0+5So7QAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAKAAAADADQ7ZAuAAAACQ==",
			"Program data: v1Vap4TfuDn+kvsVqBr/AteFVu/weaOHT6IuYqTyhWqu+Oo0N6H7vgIAAAAAAAAABt324e51j94YQl285GzN2rYa/E2DuQ0n/r35KNihi/w0fAqd5gQfkYIohsaMAW0Tz4R1O58t9Hes5QbmGAti3wAAAZPU+0A5PuT/0+5So7QAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAKAAAADAC4ZNlFAAAACQ==",
			"Program LoCoNsJFuhTkSQjfdDfn3yuwqhSYoPujmviRHVCzsqn consumed 23402 of 165315 compute units",
			"Program LoCoNsJFuhTkSQjfdDfn3yuwqhSYoPujmviRHVCzsqn success",
			"Program data: 3Q/UHSP8/04EAAAAAAAAAAUSRxvy9oST12hyi0H01X8FgtbLw1CWTtStGNnHCsctMAAAAPKMV2pH4lYg/pL7Faga/wLXhVbv8Hmjh0+iLmKk8oVqrvjqNDeh+74BAAAAAAAAAA==",
			"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX consumed 60678 of 200000 compute units",
			"Program 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX success",
		}

		timelockPubkey := solana.PublicKey{}
		err := timelockPubkey.Set("LoCoNsJFuhTkSQjfdDfn3yuwqhSYoPujmviRHVCzsqn")
		require.NoError(t, err)

		tokenPubkey := solana.PublicKey{}
		terr := tokenPubkey.Set("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
		require.NoError(t, terr)

		expected := []*AnchorInstruction{
			{
				Name:      "Execute",
				ProgramID: "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
				EventData: []*EventData{
					{
						EventName:   "OpExecuted",
						Base64Data:  "3Q/UHSP8/04EAAAAAAAAAAUSRxvy9oST12hyi0H01X8FgtbLw1CWTtStGNnHCsctMAAAAPKMV2pH4lYg/pL7Faga/wLXhVbv8Hmjh0+iLmKk8oVqrvjqNDeh+74BAAAAAAAAAA==",
						DecodedData: []byte{221, 15, 212, 29, 35, 252, 255, 78, 4, 0, 0, 0, 0, 0, 0, 0, 5, 18, 71, 27, 242, 246, 132, 147, 215, 104, 114, 139, 65, 244, 213, 127, 5, 130, 214, 203, 195, 80, 150, 78, 212, 173, 24, 217, 199, 10, 199, 45, 48, 0, 0, 0, 242, 140, 87, 106, 71, 226, 86, 32, 254, 146, 251, 21, 168, 26, 255, 2, 215, 133, 86, 239, 240, 121, 163, 135, 79, 162, 46, 98, 164, 242, 133, 106, 174, 248, 234, 52, 55, 161, 251, 190, 1, 0, 0, 0, 0, 0, 0, 0},
						Data: &OpExecuted{
							Nonce: 4,
							To:    timelockPubkey,
							Data:  []byte{242, 140, 87, 106, 71, 226, 86, 32, 254, 146, 251, 21, 168, 26, 255, 2, 215, 133, 86, 239, 240, 121, 163, 135, 79, 162, 46, 98, 164, 242, 133, 106, 174, 248, 234, 52, 55, 161, 251, 190, 1, 0, 0, 0, 0, 0, 0, 0},
						},
					},
				},
				InnerCalls: []*AnchorInstruction{
					{
						Name:      "ScheduleBatch",
						ProgramID: "LoCoNsJFuhTkSQjfdDfn3yuwqhSYoPujmviRHVCzsqn",
						EventData: []*EventData{{
							EventName:   "CallScheduled",
							Base64Data:  "v1Vap4TfuDn+kvsVqBr/AteFVu/weaOHT6IuYqTyhWqu+Oo0N6H7vgAAAAAAAAAABt324e51j94YQl285GzN2rYa/E2DuQ0n/r35KNihi/w0fAqd5gQfkYIohsaMAW0Tz4R1O58t9Hes5QbmGAti3wAAAZPU+0A5PuT/0+5So7QAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAKAAAADADodkgXAAAACQ==",
							DecodedData: []byte{191, 85, 90, 167, 132, 223, 184, 57, 254, 146, 251, 21, 168, 26, 255, 2, 215, 133, 86, 239, 240, 121, 163, 135, 79, 162, 46, 98, 164, 242, 133, 106, 174, 248, 234, 52, 55, 161, 251, 190, 0, 0, 0, 0, 0, 0, 0, 0, 6, 221, 246, 225, 238, 117, 143, 222, 24, 66, 93, 188, 228, 108, 205, 218, 182, 26, 252, 77, 131, 185, 13, 39, 254, 189, 249, 40, 216, 161, 139, 252, 52, 124, 10, 157, 230, 4, 31, 145, 130, 40, 134, 198, 140, 1, 109, 19, 207, 132, 117, 59, 159, 45, 244, 119, 172, 229, 6, 230, 24, 11, 98, 223, 0, 0, 1, 147, 212, 251, 64, 57, 62, 228, 255, 211, 238, 82, 163, 180, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 12, 0, 232, 118, 72, 23, 0, 0, 0, 9},
							Data: &CallScheduled{
								ID:          [32]byte{254, 146, 251, 21, 168, 26, 255, 2, 215, 133, 86, 239, 240, 121, 163, 135, 79, 162, 46, 98, 164, 242, 133, 106, 174, 248, 234, 52, 55, 161, 251, 190},
								Index:       0,
								Target:      tokenPubkey,
								Predecessor: [32]byte{52, 124, 10, 157, 230, 4, 31, 145, 130, 40, 134, 198, 140, 1, 109, 19, 207, 132, 117, 59, 159, 45, 244, 119, 172, 229, 6, 230, 24, 11, 98, 223},
								Salt:        [32]byte{0, 0, 1, 147, 212, 251, 64, 57, 62, 228, 255, 211, 238, 82, 163, 180, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
								Delay:       1,
								Data:        []byte{12, 0, 232, 118, 72, 23, 0, 0, 0, 9},
							},
						},
							{
								EventName:   "CallScheduled",
								Base64Data:  "v1Vap4TfuDn+kvsVqBr/AteFVu/weaOHT6IuYqTyhWqu+Oo0N6H7vgEAAAAAAAAABt324e51j94YQl285GzN2rYa/E2DuQ0n/r35KNihi/w0fAqd5gQfkYIohsaMAW0Tz4R1O58t9Hes5QbmGAti3wAAAZPU+0A5PuT/0+5So7QAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAKAAAADADQ7ZAuAAAACQ==",
								DecodedData: []byte{191, 85, 90, 167, 132, 223, 184, 57, 254, 146, 251, 21, 168, 26, 255, 2, 215, 133, 86, 239, 240, 121, 163, 135, 79, 162, 46, 98, 164, 242, 133, 106, 174, 248, 234, 52, 55, 161, 251, 190, 1, 0, 0, 0, 0, 0, 0, 0, 6, 221, 246, 225, 238, 117, 143, 222, 24, 66, 93, 188, 228, 108, 205, 218, 182, 26, 252, 77, 131, 185, 13, 39, 254, 189, 249, 40, 216, 161, 139, 252, 52, 124, 10, 157, 230, 4, 31, 145, 130, 40, 134, 198, 140, 1, 109, 19, 207, 132, 117, 59, 159, 45, 244, 119, 172, 229, 6, 230, 24, 11, 98, 223, 0, 0, 1, 147, 212, 251, 64, 57, 62, 228, 255, 211, 238, 82, 163, 180, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 12, 0, 208, 237, 144, 46, 0, 0, 0, 9},
								Data: &CallScheduled{
									ID:          [32]byte{254, 146, 251, 21, 168, 26, 255, 2, 215, 133, 86, 239, 240, 121, 163, 135, 79, 162, 46, 98, 164, 242, 133, 106, 174, 248, 234, 52, 55, 161, 251, 190},
									Index:       1,
									Target:      tokenPubkey,
									Predecessor: [32]byte{52, 124, 10, 157, 230, 4, 31, 145, 130, 40, 134, 198, 140, 1, 109, 19, 207, 132, 117, 59, 159, 45, 244, 119, 172, 229, 6, 230, 24, 11, 98, 223},
									Salt:        [32]byte{0, 0, 1, 147, 212, 251, 64, 57, 62, 228, 255, 211, 238, 82, 163, 180, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
									Delay:       1,
									Data:        []byte{12, 0, 208, 237, 144, 46, 0, 0, 0, 9},
								},
							},
							{
								EventName:   "CallScheduled",
								Base64Data:  "v1Vap4TfuDn+kvsVqBr/AteFVu/weaOHT6IuYqTyhWqu+Oo0N6H7vgIAAAAAAAAABt324e51j94YQl285GzN2rYa/E2DuQ0n/r35KNihi/w0fAqd5gQfkYIohsaMAW0Tz4R1O58t9Hes5QbmGAti3wAAAZPU+0A5PuT/0+5So7QAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAKAAAADAC4ZNlFAAAACQ==",
								DecodedData: []byte{191, 85, 90, 167, 132, 223, 184, 57, 254, 146, 251, 21, 168, 26, 255, 2, 215, 133, 86, 239, 240, 121, 163, 135, 79, 162, 46, 98, 164, 242, 133, 106, 174, 248, 234, 52, 55, 161, 251, 190, 2, 0, 0, 0, 0, 0, 0, 0, 6, 221, 246, 225, 238, 117, 143, 222, 24, 66, 93, 188, 228, 108, 205, 218, 182, 26, 252, 77, 131, 185, 13, 39, 254, 189, 249, 40, 216, 161, 139, 252, 52, 124, 10, 157, 230, 4, 31, 145, 130, 40, 134, 198, 140, 1, 109, 19, 207, 132, 117, 59, 159, 45, 244, 119, 172, 229, 6, 230, 24, 11, 98, 223, 0, 0, 1, 147, 212, 251, 64, 57, 62, 228, 255, 211, 238, 82, 163, 180, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 12, 0, 184, 100, 217, 69, 0, 0, 0, 9},
								Data: &CallScheduled{
									ID:          [32]byte{254, 146, 251, 21, 168, 26, 255, 2, 215, 133, 86, 239, 240, 121, 163, 135, 79, 162, 46, 98, 164, 242, 133, 106, 174, 248, 234, 52, 55, 161, 251, 190},
									Index:       2,
									Target:      tokenPubkey,
									Predecessor: [32]byte{52, 124, 10, 157, 230, 4, 31, 145, 130, 40, 134, 198, 140, 1, 109, 19, 207, 132, 117, 59, 159, 45, 244, 119, 172, 229, 6, 230, 24, 11, 98, 223},
									Salt:        [32]byte{0, 0, 1, 147, 212, 251, 64, 57, 62, 228, 255, 211, 238, 82, 163, 180, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
									Delay:       1,
									Data:        []byte{12, 0, 184, 100, 217, 69, 0, 0, 0, 9},
								},
							},
						},
					},
				},
			},
		}

		result := ParseLogMessages(logs, []EventMapping{
			EventMappingFor[OpExecuted]("OpExecuted"),
			EventMappingFor[CallScheduled]("CallScheduled"),
		})
		require.Equal(t, len(expected), len(result), "Instruction count mismatch - expected: %d, got: %d", len(expected), len(result))

		// verify Execute instruction result
		for i, instruction := range result {
			require.Equal(t, expected[i].Name, instruction.Name)
			require.Equal(t, expected[i].ProgramID, instruction.ProgramID)

			// verify OpExecuted event
			require.Equal(t, len(expected[i].EventData), len(instruction.EventData))
			if len(expected[i].EventData) > 0 {
				require.Equal(t, expected[i].EventData[0].EventName, instruction.EventData[0].EventName)
				require.Equal(t, expected[i].EventData[0].Base64Data, instruction.EventData[0].Base64Data)
				require.Equal(t, expected[i].EventData[0].DecodedData, instruction.EventData[0].DecodedData)
				require.Equal(t, expected[i].EventData[0].Data.(*OpExecuted).To, instruction.EventData[0].Data.(*OpExecuted).To)
			}

			// verify inner calls (ScheduleBatch)
			require.Equal(t, len(expected[i].InnerCalls), len(instruction.InnerCalls))
			if len(instruction.InnerCalls) > 0 {
				innerCall := instruction.InnerCalls[0]
				expectedInner := expected[i].InnerCalls[0]

				require.Equal(t, expectedInner.Name, innerCall.Name)
				require.Equal(t, expectedInner.ProgramID, innerCall.ProgramID)

				// verify CallScheduled events and their indices
				require.Equal(t, len(expectedInner.EventData), len(innerCall.EventData))
				for j := range innerCall.EventData {
					event := innerCall.EventData[j]
					expectedEvent := expectedInner.EventData[j]

					require.Equal(t, expectedEvent.EventName, event.EventName)
					require.Equal(t, expectedEvent.Base64Data, event.Base64Data)
					require.Equal(t, expectedEvent.DecodedData, event.DecodedData)

					scheduledEvent, ok := event.Data.(*CallScheduled)
					require.True(t, ok)

					expectedEventData, ok := expectedEvent.Data.(*CallScheduled)
					require.True(t, ok)

					require.Equal(t, uint64(j), scheduledEvent.Index, "Event index mismatch at position %d", j)
					require.Equal(t, expectedEventData.ID, scheduledEvent.ID, "Event ID mismatch at position %d", j)
					require.Equal(t, expectedEventData.Target, scheduledEvent.Target, "Event target mismatch at position %d", j)
					require.Equal(t, expectedEventData.Predecessor, scheduledEvent.Predecessor, "Event predecessor mismatch at position %d", j)
					require.Equal(t, expectedEventData.Salt, scheduledEvent.Salt, "Event salt mismatch at position %d", j)
					require.Equal(t, expectedEventData.Delay, scheduledEvent.Delay, "Event delay mismatch at position %d", j)
					require.Equal(t, expectedEventData.Data, scheduledEvent.Data, "Event data mismatch at position %d", j)
				}
			}
		}
	})
}

func assertInstructionEqual(t *testing.T, expected, actual *AnchorInstruction) {
	t.Helper()

	require.Equal(t, expected.Name, actual.Name, "Instruction name mismatch - expected: %s, got: %s", expected.Name, actual.Name)
	require.Equal(t, expected.ProgramID, actual.ProgramID, "Program ID mismatch - expected: %s, got: %s", expected.ProgramID, actual.ProgramID)
	require.Equal(t, len(expected.Logs), len(actual.Logs), "Log count mismatch - expected: %d, got: %d", len(expected.Logs), len(actual.Logs))

	for i := range expected.Logs {
		require.Equal(t, expected.Logs[i], actual.Logs[i], "Log mismatch at index %d - expected: %s, got: %s", i, expected.Logs[i], actual.Logs[i])
	}

	require.Equal(t, len(expected.InnerCalls), len(actual.InnerCalls), "Inner calls count mismatch - expected: %d, got: %d", len(expected.InnerCalls), len(actual.InnerCalls))

	for i := range expected.InnerCalls {
		assertInstructionEqual(t, expected.InnerCalls[i], actual.InnerCalls[i])
	}
}
