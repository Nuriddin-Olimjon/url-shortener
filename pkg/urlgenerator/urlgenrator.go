package urlgenerator

import (
	"errors"

	"github.com/speps/go-hashids"
)

var hashID *hashids.HashID

func init() {
	hashConf := hashids.NewData()
	hashConf.Salt = "salt"
	hashConf.MinLength = 8 // Hash ID length, 8 means a hash ID such as xxxxxxxx
	hashId, err := hashids.NewWithData(hashConf)
	if err != nil {
		panic(err)
	}
	hashID = hashId
}

func Decode(hash string) (int32, error) {
	id, err := hashID.DecodeInt64WithError(hash)
	if err != nil {
		return 0, err
	}
	if len(id) != 1 {
		return 0, errors.New("invalid hash string for id")
	}
	return int32(id[0]), nil
}
func Encode(id int32) (string, error) {
	return hashID.EncodeInt64([]int64{int64(id)})
}
