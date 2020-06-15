package test

type mymap struct {
	cap     int64
	buckets []*bucket
}

type bucket struct {
	key interface{}
	val interface{}
}

func NewStructure(cap int64) *mymap {
	m := new(mymap)
	if cap == 0 {
		m.cap = 1 << 4
	} else {
		m.cap = cap
	}
	return m
}

func (m *mymap) isExist(k interface{}) (i int, isExist bool) {
	for i, b := range m.buckets {
		if k == b.key {
			return i, true
		}
	}
	return 0, false
}

func (m *mymap) put(k interface{}, v interface{}) {
	if i, exist := m.isExist(k); exist {
		m.buckets[i].val = v
	} else {
		b := bucket{
			key: k,
			val: v,
		}
		m.buckets = append(m.buckets, &b)
	}
}
func (m *mymap) get(k interface{}) interface{} {
	if i, exist := m.isExist(k); exist {
		return m.buckets[i].val
	}
	return nil
}
func (m *mymap) del(k interface{}) {
	if i, exist := m.isExist(k); exist {
		m.buckets = append(m.buckets[:i], m.buckets[i+1:]...)
	}
}
