package poset

import (
	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"math"
)

// TODO: make FrameInfo internal

type FrameInfo struct {
	TimeOffset int64 // may be negative
	TimeRatio  inter.Timestamp
}

type frameInfoMarshaling struct {
	TimeOffset uint64
	TimeRatio  inter.Timestamp
}

func (f *FrameInfo) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, frameInfoMarshaling{
		TimeOffset: uint64(f.TimeOffset + math.MaxInt64/2),
		TimeRatio:  f.TimeRatio,
	})
}

func (f *FrameInfo) DecodeRLP(st *rlp.Stream) error {
	m := frameInfoMarshaling{}
	if err := st.Decode(&m); err != nil {
		return err
	}
	f.TimeOffset = int64(m.TimeOffset) - math.MaxInt64/2
	f.TimeRatio = m.TimeRatio
	return nil
}

// GetConsensusTimestamp calc consensus timestamp for given event.
func (f *FrameInfo) GetConsensusTimestamp(e *Event) inter.Timestamp {
	return inter.Timestamp(int64(e.Lamport)*int64(f.TimeRatio) + f.TimeOffset)
}
