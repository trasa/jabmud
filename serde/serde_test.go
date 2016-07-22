package serde

import (
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type SerdeObj struct {
	Foo string
	Num int
}

type SerdeSuite struct {
	suite.Suite
}

func TestSerdeSuite(t *testing.T) {
	suite.Run(t, new(SerdeSuite))
}

func (suite *SerdeSuite) TestSerializeDeserialize() {
	original := SerdeObj{Foo: "foo", Num: 42}
	s := Serialize(original)

	deserialized := SerdeObj{}
	Deserialize(s, &deserialized)

	log.Printf("after: %v", deserialized)
	suite.Equal("foo", deserialized.Foo)
	suite.Equal(42, deserialized.Num)

}
