package tpms

import (
	"strconv"
	"testing"
)

type sensoMock struct {
	value int
}

func (m sensoMock) popNextPressurePsiValue() int {
	return m.value
}

func Test_alarm_check_sensoroff(t *testing.T) {
	mySensor := sensoMock{}
	alarmToTest := NewAlarm(&mySensor)

	values := []int{17, 18, 19, 20, 21}

	for _, v := range values {
		t.Run(strconv.Itoa(v), func(t *testing.T) {
			mySensor.value = v
			alarmToTest.check()
			shouldBeOff(t, alarmToTest)
		})
	}

}

func Test_alarm_check_sensoron(t *testing.T) {
	mySensor := sensoMock{}
	alarmToTest := NewAlarm(&mySensor)

	values := []int{0, -1, 10000, 16, 22}

	for _, v := range values {
		t.Run(strconv.Itoa(v), func(t *testing.T) {
			mySensor.value = v
			alarmToTest.check()
			shouldBeOn(t, alarmToTest)
		})
	}
}

func shouldBeOff(t *testing.T, a *alarm) {
	if a.alarmOn {
		t.Errorf("alarm should be off")
	}
}

func shouldBeOn(t *testing.T, a *alarm) {
	if !a.alarmOn {
		t.Errorf("alarm should be on")
	}
}
