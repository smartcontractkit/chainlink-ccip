package utils

import (
	"reflect"
	"testing"
)

func TestParseLogMessages(t *testing.T) {
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
					Name:         "Execute",
					ProgramID:    "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:         []string{},
					ComputeUnits: 35400,
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "Empty",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `empty` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps }",
							},
							ComputeUnits: 13620,
							InnerCalls:   []*AnchorInstruction{},
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
					Name:         "Execute",
					ProgramID:    "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:         []string{},
					ComputeUnits: 35463,
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "U8InstructionData",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `u8_instruction_data` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data 123",
							},
							ComputeUnits: 13648,
							InnerCalls:   []*AnchorInstruction{},
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
					Name:         "Execute",
					ProgramID:    "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:         []string{},
					ComputeUnits: 35152,
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "StructInstructionData",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `struct_instruction_data` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: Empty, remaining_accounts: [], bumps: EmptyBumps } and data Value { value: 234 }",
							},
							ComputeUnits: 13920,
							InnerCalls:   []*AnchorInstruction{},
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
					Name:         "Execute",
					ProgramID:    "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:         []string{},
					ComputeUnits: 69682,
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "AccountRead",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `account_read` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: AccountRead { u8_value: Account { account: Value { value: 1 }, info: AccountInfo { key: 8WGXBpVJrBATopzT8iXvRuvp5f3U63uB13tfQjGoi6rM, owner: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, is_signer: false, is_writable: false, executable: false, rent_epoch: 18446744073709551615, lamports: 953520, data.len: 9, data: 879ef47548cb18c201, .. } } }, remaining_accounts: [], bumps: AccountReadBumps { u8_value: 255 } }",
							},
							ComputeUnits: 45559,
							InnerCalls:   []*AnchorInstruction{},
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
					Name:         "Execute",
					ProgramID:    "6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX",
					Logs:         []string{},
					ComputeUnits: 139571,
					InnerCalls: []*AnchorInstruction{
						{
							Name:      "AccountMut",
							ProgramID: "4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ",
							Logs: []string{
								"Called `account_mut` Context { program_id: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, accounts: AccountMut { u8_value: Account { account: Value { value: 1 }, info: AccountInfo { key: 8WGXBpVJrBATopzT8iXvRuvp5f3U63uB13tfQjGoi6rM, owner: 4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ, is_signer: false, is_writable: true, executable: false, rent_epoch: 18446744073709551615, lamports: 953520, data.len: 9, data: 879ef47548cb18c201, .. } }, stub_caller: Signer { info: AccountInfo { key: BUx7YZMoVXCnT2BewMZc2hr8yxoiihtHdDuoa19D9R5q, owner: 6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX, is_signer: true, is_writable: true, executable: false, rent_epoch: 18446744073709551615, lamports: 2874480, data.len: 285, data: 2c3eace1f603b2211266a21317e30848020102010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000, .. } }, system_program: Program { info: AccountInfo { key: 11111111111111111111111111111111, owner: NativeLoader1111111111111111111111111111111, is_signer: false, is_writable: false, executable: true, rent_epoch: 0, lamports: 1, data.len: 14, data: 73797374656d5f70726f6772616d, .. } } }, remaining_accounts: [], bumps: AccountMutBumps { u8_value: 255 } }",
							},
							ComputeUnits: 111015,
							InnerCalls:   []*AnchorInstruction{},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseLogMessages(tt.logs)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test %s failed.\nExpected:\n%#v\nGot:\n%#v", tt.name, tt.expected, result)
			}
		})
	}
}
