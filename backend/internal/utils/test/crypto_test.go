package test

import (
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gookit/goutil/testutil/assert"
	"testing"
)

const (
	Text   = "1145141919810"
	AesKey = "32150285b345c48aa3492f9212f61ca2"
)

func TestAesCipher(t *testing.T) {
	cipher, err := utils.NewAesCipher(AesKey)
	if err != nil {
		t.Error(err)
	}
	encrypt, err := cipher.Encrypt(Text)
	if err != nil {
		t.Error(err)
	}
	decrypt, err := cipher.Decrypt(encrypt)
	if err != nil {
		t.Error(err)
	}
	assert.Eq(t, Text, decrypt)
}

func TestMD5(t *testing.T) {
	md5 := utils.MD5(Text)
	assert.Eq(t, AesKey, md5)
}

func TestSHA256(t *testing.T) {
	sha256 := utils.SHA256(Text)
	assert.Eq(t, "b7ab30a912521ac36e433a5cfc8b5c1037884487af45ae5311ced235ee77faef", sha256)
}
