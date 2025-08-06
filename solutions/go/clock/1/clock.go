package clock

import "strconv"

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	all_min := h*60 + m
	for all_min < 0 || all_min >= 1440 {
		if all_min < 0 {
			all_min = all_min + 1440
		}
		if all_min >= 1440 {
			all_min = all_min - 1440
		}
	}

	var cc Clock
	cc.hour = all_min / 60
	cc.minute = (all_min - (all_min/60)*60)
	return cc
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	h := c.hour
	m := c.minute
	var hs string
	var ms string
	if h < 10 {
		hs = "0" + strconv.Itoa(h)
	} else {
		hs = strconv.Itoa(h)
	}
	if m < 10 {
		ms = "0" + strconv.Itoa(m)
	} else {
		ms = strconv.Itoa(m)
	}
	return hs + ":" + ms
}
