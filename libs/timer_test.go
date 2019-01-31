package libs

import "testing"

func TestTimeStamp(t *testing.T) {
	TimeStamp()
}
func BenchmarkTimeStamp(b *testing.B) {
	TimeStamp()
}

func TestNow(t *testing.T) {
	Now()
}

func BenchmarkNow(b *testing.B) {
	Now()
}

func TestGetTimeByDurationBegin(t *testing.T) {
	mark := "2019-01-26 00:00:00"
	if GetTimeByDurationBegin(-3) != mark {
		t.Fail()
	}
}

func BenchmarkGetTimeByDurationBegin(b *testing.B) {
	mark := "2019-01-26 00:00:00"
	if GetTimeByDurationBegin(-3) != mark {
		b.Fail()
	}
}

func TestStrToTime(t *testing.T) {
	x := "2019-01-26 00:00:00"
	StrToTime(x)
}

//func TestGetTimeByDurationEnd(t *testing.T) {
//	mark := "2019-01-26 23:59:59"
//	if GetTimeByDurationBegin(-3) != mark {
//		t.Fail()
//	}
//}
