package blocksources

import (
	"bytes"
	"github.com/Redundancy/go-sync/patcher"
	"testing"
)

const STRING_DATA = "abcdefghijklmnopqrst"

var BYTE_BLOCK_SOURCE = NewReadSeekerBlockSource(
	bytes.NewReader([]byte(STRING_DATA)),
)

func TestReadFirstBlock(t *testing.T) {
	BYTE_BLOCK_SOURCE.RequestBlock(
		patcher.MissingBlockSpan{
			BlockSize:  4,
			StartBlock: 0,
			EndBlock:   0,
		},
	)

	result := <-BYTE_BLOCK_SOURCE.GetResultChannel()

	if result.StartBlock != 0 {
		t.Errorf("Wrong start block: %v", result.StartBlock)
	}

	EXPECTED := STRING_DATA[:4]

	if bytes.Compare(result.Data, []byte(EXPECTED)) != 0 {
		t.Errorf(
			"Unexpected result data: \"%v\" expected: \"%v\"",
			string(result.Data),
			EXPECTED,
		)
	}
}
