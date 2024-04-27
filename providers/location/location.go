package location

import (
	"math/rand"
	"time"
)

// Lists of provinces and cities
var provinces = []string{
	"آذربایجان شرقی", "آذربایجان غربی", "اردبیل", "اصفهان", "البرز",
	"ایلام", "بوشهر", "تهران", "چهارمحال و بختیاری", "خراسان جنوبی",
	"خراسان رضوی", "خراسان شمالی", "خوزستان", "زنجان", "سمنان",
	"سیستان و بلوچستان", "فارس", "قزوین", "قم", "کردستان",
	"کرمان", "کرمانشاه", "کهگیلویه و بویراحمد", "گلستان", "گیلان",
	"لرستان", "مازندران", "مرکزی", "هرمزگان", "همدان", "یزد",
}

var cities = []string{
	"تبریز", "ارومیه", "اردبیل", "اصفهان", "کرج",
	"ایلام", "بوشهر", "تهران", "شهرکرد", "بیرجند",
	"مشهد", "بجنورد", "اهواز", "زنجان", "سمنان",
	"زاهدان", "شیراز", "قزوین", "قم", "سنندج",
	"کرمان", "کرمانشاه", "یاسوج", "گرگان", "رشت",
	"خرم آباد", "ساری", "اراک", "بندرعباس", "همدان", "یزد",
}

// RandomProvince returns a random province.
func RandomProvince() string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return provinces[rng.Intn(len(provinces))]
}

// RandomCity returns a random city.
func RandomCity() string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return cities[rng.Intn(len(cities))]
}
