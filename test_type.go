package test_1

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

type TestT struct {
	Id1, Id2, Id3 int
}

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}
