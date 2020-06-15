package test

import "testing"

func TestMap(t *testing.T) {
	m := NewStructure(0)
	m.put("a", "value")
	m.put("a", "value2")
	t.Log(m)

	t.Log(m.get("a"))
	t.Log("del key a")
	m.del("a")
	t.Log(m.get("a"))

}
