// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package blowfish implements Bruce Schneier's Blowfish encryption algorithm.
//
// Blowfish is a legacy cipher and its short block size makes it vulnerable to
// birthday bound attacks (see https://sweet32.info). It should only be used
// where compatibility with legacy systems, not security, is the goal.
//
// Deprecated: any new system should use AES (from crypto/aes, if necessary in
// an AEAD mode like crypto/cipher.NewGCM) or XChaCha20-Poly1305 (from
// golang.org/x/crypto/chacha20poly1305).
package blowfish // import "golang.org/x/crypto/blowfish"

// The code is a port of Bruce Schneier's C implementation.
// See https://www.schneier.com/blowfish.html.

import "strconv"

// A Cipher is an instance of Blowfish encryption using a particular key.
type Cipher struct {
	p              [18]uint32
	s0, s1, s2, s3 [256]uint32
}

type KeySizeError int

func (k KeySizeError) Error() string {
	return "crypto/blowfish: invalid key size " + strconv.Itoa(int(k))
}

// NewCipher creates and returns a Cipher.
// The key argument should be the Blowfish key, from 1 to 56 bytes.
func NewCipher(key []byte) (*Cipher, error) {
	var result Cipher
	if k := len(key); k < 1 || k > 56 {
		return nil, KeySizeError(k)
	}
	initCipher(&result)
	ExpandKey(key, &result)
	return &result, nil
}

// Encrypt encrypts the 8-byte buffer src using the key k
// and stores the result in dst.
// Note that for amounts of data larger than a block,
// it is not safe to just call Encrypt on successive blocks;
// instead, use an encryption mode like CBC (see crypto/cipher/cbc.go).
func (c *Cipher) Encrypt(dst, src []byte, sIndex, dIndex int) {
	l := uint32(src[sIndex+3])<<24 | uint32(src[sIndex+2])<<16 | uint32(src[sIndex+1])<<8 | uint32(src[sIndex])
	r := uint32(src[sIndex+7])<<24 | uint32(src[sIndex+6])<<16 | uint32(src[sIndex+5])<<8 | uint32(src[sIndex+4])
	l, r = encryptBlock(l, r, c)
	dst[dIndex+3], dst[dIndex+2], dst[dIndex+1], dst[dIndex] = byte(l>>24), byte(l>>16), byte(l>>8), byte(l)
	dst[dIndex+7], dst[dIndex+6], dst[dIndex+5], dst[dIndex+4] = byte(r>>24), byte(r>>16), byte(r>>8), byte(r)
}

// Decrypt decrypts the 8-byte buffer src using the key k
// and stores the result in dst.
func (c *Cipher) Decrypt(dst, src []byte, sIndex, dIndex int) {
	l := uint32(src[sIndex+3])<<24 | uint32(src[sIndex+2])<<16 | uint32(src[sIndex+1])<<8 | uint32(src[sIndex])
	r := uint32(src[sIndex+7])<<24 | uint32(src[sIndex+6])<<16 | uint32(src[sIndex+5])<<8 | uint32(src[sIndex+4])
	r, l = decryptBlock(l, r, c)
	dst[dIndex+3], dst[dIndex+2], dst[dIndex+1], dst[dIndex] = byte(r>>24), byte(r>>16), byte(r>>8), byte(r)
	dst[dIndex+7], dst[dIndex+6], dst[dIndex+5], dst[dIndex+4] = byte(l>>24), byte(l>>16), byte(l>>8), byte(l)
}

func initCipher(c *Cipher) {
	copy(c.p[0:], p[0:])
	copy(c.s0[0:], s0[0:])
	copy(c.s1[0:], s1[0:])
	copy(c.s2[0:], s2[0:])
	copy(c.s3[0:], s3[0:])
}
