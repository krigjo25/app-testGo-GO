package main

import (
	"encoding/base64"
	"fmt"
)

type folio struct {
	buf  []byte
	seed int
}

func (f *folio) bar(b1, b2 *byte, i, j int) {
	*b1 = byte(j>>8) ^ f.buf[(i+f.seed)%len(f.buf)]
	*b2 = byte(j) ^ f.buf[(i+f.seed+1)%len(f.buf)]
}

func main() {
	secret, _ := base64.StdEncoding.
		DecodeString("eW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQ==")

	baz := make([]byte, 20)
	defer func() {
		fmt.Println(string(baz))
	}()

	f := folio{buf: secret, seed: 1337}
	defer f.bar(&baz[8], &baz[9], 8, 21)
	defer f.bar(&baz[0], &baz[1], 0, 4634)
	defer f.bar(&baz[12], &baz[13], 12, 2139)
	defer f.bar(&baz[10], &baz[11], 10, 21783)
	defer f.bar(&baz[14], &baz[15], 14, 20754)
	defer f.bar(&baz[16], &baz[17], 16, 3924)
	defer f.bar(&baz[4], &baz[5], 4, 16717)
	defer f.bar(&baz[2], &baz[3], 2, 3596)
	defer f.bar(&baz[6], &baz[7], 6, 258)

	f.seed = 65537
	f.bar(&baz[18], &baz[19], 18, 25410)
}
