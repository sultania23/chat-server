package hash_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/sultania23/chat-server/pkg/hash"
	"testing"
)

const (
	salt                 = "asd51cg"
	templatePasswordHash = "08baf1a5ec61af402f2f022ce3ba14c814d802c1"
)

func TestHash(t *testing.T) {
	testPassword := "vb1hj7er"

	testHash := hash.NewSHA1Hasher(salt)

	testPasswordHash := testHash.Hash(testPassword)

	assert.Equal(t, templatePasswordHash, testPasswordHash,
		fmt.Sprintf("Incorrect result. Expect %s, got %s", templatePasswordHash, testPasswordHash))
}
