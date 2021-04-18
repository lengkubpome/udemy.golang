package coveragedemo

import (
	"math"
	"testing"
)

func TestBasicOperation(t *testing.T) {
	x := operation(1, 2, '+')
	if x != 3 {
		t.Errorf("operation(%f,%f) = %f. Wanted: %f", 1.0, 2.0, x, 3.0)
	}
	x = operation(1, 2, '-')
	if x != -1 {
		t.Errorf("operation(%f,%f) = %f. Wanted: %f", 1.0, 2.0, x, -1.0)
	}
	x = operation(1, 2, '*')
	if x != 2 {
		t.Errorf("operation(%f,%f) = %f. Wanted: %f", 1.0, 2.0, x, 2.0)
	}
	x = operation(1.0, 3.0, '/')
	epsilon := math.Nextafter(1.0, 2.0) - 1.0 // epsilon คือค่าที่เล็กมากๆ
	if (x - (1.0 / 3.0)) > epsilon {          // 0.003333 > 0.003333
		t.Errorf("operation(%f,%f) = %f. Wanted: %f", 1.0, 2.0, x, 0.33)
	}
}

// สร้างไฟล์ coverate มาก่อน
// go test -cover -coverprofile cover.out
// ทำการตรวจสอบ function ที่ทดสอบแล้วผ่าน browser
// go tool cover -html cover.out
