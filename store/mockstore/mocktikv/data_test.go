package mocktikv

import (
	"encoding/hex"
	"fmt"
	"github.com/pingcap/goleveldb/leveldb"
	"testing"
)

func TestReadData(t *testing.T) {
	ldb, err := leveldb.OpenFile("/tmp/tinysql", nil)
	if err != nil {
		t.Fatal(err)
	}
	bs, _ := hex.DecodeString("7480000000000000155f72800000000000ea62")
	t.Log(bs)
	v, err := ldb.Get(bs, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

//read key t??_r????*?9?w??, value ?f???f???tom ||| [116 128 0 0 0 0 0 0 255 21 95 114 128 0 0 0 0 255 0 0 1 0 0 0 0 0 250 250 42 253 57 153 119 255 254], [0 0 0 0 0 0 0 0 0 0 132 102 198 2 213 5 1 0 136 102 198 2 213 5 20 128 0 3 0 0 0 1 2 3 1 0 4 0 5 0 1 116 111 109 0]
//read key t??_r??u1??*?0?K??, value ?>???>???tom ||| [116 128 0 0 0 0 0 0 255 21 95 114 128 0 0 0 0 255 0 117 49 0 0 0 0 0 250 250 42 253 48 193 75 255 255], [0 0 0 0 0 0 0 0 0 0 172 62 207 2 213 5 0 0 180 62 207 2 213 5 20 128 0 3 0 0 0 1 2 3 1 0 4 0 5 0 2 116 111 109 0]
//read key t??_r???a??*뚾w??, value ?Ae??Ae??tom ||| [116 128 0 0 0 0 0 0 255 21 95 114 128 0 0 0 0 255 0 234 97 0 0 0 0 0 250 250 42 235 154 190 119 255 255], [0 0 0 0 0 0 0 0 0 0 128 65 101 20 213 5 0 0 136 65 101 20 213 5 20 128 0 3 0 0 0 1 2 3 1 0 4 0 5 0 3 116 111 109 0]
//read key t??_r???b??*넛k??, value ?d{??d{??tom ||| [116 128 0 0 0 0 0 0 255 21 95 114 128 0 0 0 0 255 0 234 98 0 0 0 0 0 250 250 42 235 132 155 107 255 255], [0 0 0 0 0 0 0 0 0 0 144 100 123 20 213 5 0 0 148 100 123 20 213 5 20 128 0 3 0 0 0 1 2 3 1 0 4 0 5 0 3 116 111 109 0]
func TestRangeData(t *testing.T) {
	ldb, err := leveldb.OpenFile("/tmp/tinysql", nil)
	if err != nil {
		t.Fatal(err)
	}
	iter := ldb.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("read key %s, value %s ||| %v, %v\n", string(iter.Key()), string(iter.Value()), iter.Key(), iter.Value())
	}
}
