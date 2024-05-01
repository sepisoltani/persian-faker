package bill

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Bill struct{}

// billTypes contains various types of Persian bills.
var billTypes = map[int]string{
	1:  "آب",
	2:  "برق",
	3:  "گاز",
	4:  "تلفن ثابت",
	5:  "تلفن همراه",
	6:  "عوارض شهرداری",
	7:  "مالیات بر ارزش افزوده",
	8:  "سازمان مالیات",
	9:  "جرایم راهنمایی و رانندگی",
	10: "بیمه",
	11: "تلویزیون",
}

// BillType returns a random bill type from the map.
func (Bill) BillType() string {

	keys := make([]int, 0, len(billTypes))
	for k := range billTypes {
		keys = append(keys, k)
	}
	return billTypes[keys[rng.Intn(len(keys))]]
}
