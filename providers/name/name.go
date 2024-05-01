package name

import (
	"math/rand"
	"time"
)

var firstNames = []string{
	"علی",
	"مریم",
	"حسن",
	"زهرا",
	"رضا",
	"فاطمه",
	"محمد",
	"سارا",
	"جواد",
	"لیلا",
	"امیر",
	"نازنین",
	"مجید",
	"سمانه",
	"کیان",
	"نگین",
	"شهاب",
	"پریسا",
	"مهدی",
	"الهام",
	"عرفان",
	"تینا",
	"بهنام",
	"زینب",
	"پویا",
	"سپیده",
	"کامران",
	"ملیکا",
	"یاسر",
	"مینا",
	"بهزاد",
	"آناهیتا",
	"فرشاد",
	"مهسا",
	"رامین",
	"سحر",
	"عباس",
	"فریبا",
	"هومن",
	"آزاده",
	"بهروز",
	"سمیرا",
	"سعید",
	"شیرین",
	"آرش",
	"کتایون",
	"فرید",
	"هدیه",
	"مانی",
	"رویا",
}

var lastNames = []string{
	"محمدی",
	"هاشمی",
	"سلطانی",
	"احمدی",
	"رضایی",
	"کریمی",
	"جعفری",
	"حسینی",
	"نجفی",
	"شیرازی",
	"فرزانه",
	"علوی",
	"رستمی",
	"موسوی",
	"کاشانی",
	"تهرانی",
	"صادقی",
	"قاسمی",
	"میرزایی",
	"نوری",
	"رحیمی",
	"کوهی",
	"شاهین",
	"زارع",
	"بهاری",
	"عبدی",
	"چراغی",
	"امینی",
	"پهلوان",
	"تاج",
	"داودی",
	"صفری",
	"باقری",
	"زمانی",
	"کوچکی",
	"آقایی",
	"مهدوی",
	"جمشیدی",
	"گلزار",
	"پارسا",
	"صمدی",
	"شفیعی",
	"مرادی",
	"بابایی",
	"حیدری",
	"خانی",
	"یزدانی",
	"سلطانپور",
	"فخری",
	"ملکی",
}
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Name struct {
}

// FirstName returns a random Persian first name.
func (Name) FirstName() string {
	return firstNames[rng.Intn(len(firstNames))]
}

// LastName returns a random Persian last name.
func (Name) LastName() string {
	return lastNames[rng.Intn(len(lastNames))]
}

// FullName returns a random Persian full name.
func (name Name) FullName() string {
	return name.FirstName() + " " + name.LastName()
}
