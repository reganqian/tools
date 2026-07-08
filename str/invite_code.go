package str

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
)

const salt = "WDZWKJ"
const strs = "ABCDEFGHJKMNPQRSTUVWXYZ123456789"

const (
	codeLength = 6
	codeBase   = uint64(32)
	codeSpace  = uint64(1073741824) // 32^6
)

// UserIDToInviteCode 生成固定 6 位、无碰撞、较不易混淆的邀请码。
// salt 应该是全局固定的服务端密钥，不要每个用户单独设置。
func UserIDToInviteCode(userID uint64) string {
	if userID == 0 {
		panic("userID must be positive")
	}
	if salt == "" {
		panic("salt must not be empty")
	}
	if userID >= codeSpace {
		panic("userID is too large for 6-character invite code")
	}

	n := permute(userID, salt)
	return encodeInviteCode(n)
}

func encodeInviteCode(n uint64) string {
	var result [codeLength]byte

	for i := codeLength - 1; i >= 0; i-- {
		result[i] = strs[n%codeBase]
		n /= codeBase
	}

	return string(result[:])
}

func permute(n uint64, salt string) uint64 {
	for {
		n = uint64(feistel32(uint32(n), salt))
		if n < codeSpace {
			return n
		}
	}
}

func feistel32(x uint32, salt string) uint32 {
	left := uint16(x >> 16)
	right := uint16(x)

	for round := byte(0); round < 6; round++ {
		f := feistelRound(right, salt, round)
		left, right = right, left^f
	}

	return uint32(left)<<16 | uint32(right)
}

func feistelRound(right uint16, salt string, round byte) uint16 {
	var buf [3]byte
	buf[0] = round
	binary.BigEndian.PutUint16(buf[1:], right)

	mac := hmac.New(sha256.New, []byte(salt))
	_, _ = mac.Write(buf[:])
	sum := mac.Sum(nil)

	return binary.BigEndian.Uint16(sum[:2])
}
