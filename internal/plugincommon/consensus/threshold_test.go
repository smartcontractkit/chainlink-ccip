package consensus

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GteFPlusOne(t *testing.T) {
	type args struct {
		f           int
		val         int
		expectedRes bool
	}
	testCase := []args{
		{
			f:           1,
			val:         2,
			expectedRes: true,
		},
		{
			f:           4,
			val:         5,
			expectedRes: true,
		},
		{
			f:           4,
			val:         4,
			expectedRes: false,
		},
	}

	for _, tc := range testCase {
		res := GteFPlusOne(tc.f, tc.val)
		require.Equal(t, tc.expectedRes, res)
	}
}

func Test_LtFPlusOne(t *testing.T) {
	type args struct {
		f           int
		val         int
		expectedRes bool
	}
	testCase := []args{
		{
			f:           2,
			val:         2,
			expectedRes: true,
		},
		{
			f:           5,
			val:         6,
			expectedRes: false,
		},
		{
			f:           4,
			val:         4,
			expectedRes: true,
		},
	}

	for _, tc := range testCase {
		res := LtFPlusOne(tc.f, tc.val)
		require.Equal(t, tc.expectedRes, res)
	}
}

func Test_LtTwoFPlusOne(t *testing.T) {
	type args struct {
		f           int
		val         int
		expectedRes bool
	}
	testCase := []args{
		{
			f:           2,
			val:         3,
			expectedRes: true,
		},
		{
			f:           1,
			val:         3,
			expectedRes: false,
		},
		{
			f:           4,
			val:         7,
			expectedRes: true,
		},
	}

	for _, tc := range testCase {
		res := LtTwoFPlusOne(tc.f, tc.val)
		require.Equal(t, tc.expectedRes, res)
	}
}
