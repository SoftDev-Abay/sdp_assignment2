package main

type Map struct {
	objects []imapObject
}

func (m *Map) AddObject(obj imapObject) {
	m.objects = append(m.objects, obj)
}

func (m *Map) GetObjects() []imapObject {
	return m.objects
}
