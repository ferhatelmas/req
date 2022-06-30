package wire

import (
	"bytes"
	"io"

	"github.com/imroc/req/v3/internal/protocol"
	"github.com/imroc/req/v3/internal/quicvarint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DATA_BLOCKED frame", func() {
	Context("when parsing", func() {
		It("accepts sample frame", func() {
			data := []byte{0x14}
			data = append(data, encodeVarInt(0x12345678)...)
			b := bytes.NewReader(data)
			frame, err := parseDataBlockedFrame(b, protocol.Version1)
			Expect(err).ToNot(HaveOccurred())
			Expect(frame.MaximumData).To(Equal(protocol.ByteCount(0x12345678)))
			Expect(b.Len()).To(BeZero())
		})

		It("errors on EOFs", func() {
			data := []byte{0x14}
			data = append(data, encodeVarInt(0x12345678)...)
			_, err := parseDataBlockedFrame(bytes.NewReader(data), protocol.Version1)
			Expect(err).ToNot(HaveOccurred())
			for i := range data {
				_, err := parseDataBlockedFrame(bytes.NewReader(data[:i]), protocol.Version1)
				Expect(err).To(MatchError(io.EOF))
			}
		})
	})

	Context("when writing", func() {
		It("writes a sample frame", func() {
			b := &bytes.Buffer{}
			frame := DataBlockedFrame{MaximumData: 0xdeadbeef}
			err := frame.Write(b, protocol.VersionWhatever)
			Expect(err).ToNot(HaveOccurred())
			expected := []byte{0x14}
			expected = append(expected, encodeVarInt(0xdeadbeef)...)
			Expect(b.Bytes()).To(Equal(expected))
		})

		It("has the correct min length", func() {
			frame := DataBlockedFrame{MaximumData: 0x12345}
			Expect(frame.Length(protocol.Version1)).To(Equal(1 + quicvarint.Len(0x12345)))
		})
	})
})