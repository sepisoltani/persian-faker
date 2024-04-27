package bill

import (
	"math/rand"
	"time"
)

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

// RandomBillType returns a random bill type from the map.
func RandomBillType() string {

	keys := make([]int, 0, len(billTypes))
	for k := range billTypes {
		keys = append(keys, k)
	}
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return billTypes[keys[rng.Intn(len(keys))]]
}
