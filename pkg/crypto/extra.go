package crypto

// XOR returns the XOR of two byte slices of equal length.
func XOR(a, b []byte) []byte {
	n := len(a)
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = a[i] ^ b[i]
	}
	return out
}

// Reverse returns a new slice with bytes in reverse order.
func Reverse(in []byte) []byte {
	out := make([]byte, len(in))
	for i := range in {
		out[len(in)-1-i] = in[i]
	}
	return out
}

// Checksum32 computes a simple checksum from the first 4 bytes of SHA256.
func Checksum32(data []byte) uint32 {
	sum := ComputeSHA3(data)
	return uint32(sum[0])<<24 | uint32(sum[1])<<16 | uint32(sum[2])<<8 | uint32(sum[3])
}

// PadZero pads data with zeros until its length is a multiple of n.
func PadZero(data []byte, n int) []byte {
	if n <= 0 {
		return data
	}
	rem := len(data) % n
	if rem == 0 {
		return append([]byte(nil), data...)
	}
	out := make([]byte, len(data)+(n-rem))
	copy(out, data)
	return out
}
