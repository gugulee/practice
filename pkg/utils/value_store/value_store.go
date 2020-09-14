package valuestore

import "fmt"

// Value ...
type Value interface {
	Compare(value Value) int
	CompareValue(value Value) int
	String() string
}

var _ Value = &dummy{}

type dummy struct{}

func (d *dummy) Compare(value Value) int {
	return 0
}

func (d *dummy) CompareValue(value Value) int {
	return 0
}

func (d *dummy) String() string {
	return ""
}

var _ Value = &KV{}

// KV store key and value
type KV struct {
	Key   int
	Value int
}

// Compare compare key and value
func (kv *KV) Compare(value Value) int {
	kv1, ok := value.(*KV)

	if !ok {
		return -100
	}

	if kv.Value > kv1.Value {
		return 1
	} else if kv.Value < kv1.Value {
		return -1
	}

	// compare key when value are equal
	if kv.Key > kv1.Key {
		return 1
	} else if kv.Key < kv1.Key {
		return -1
	}

	// key and value are equal
	return 0
}

// CompareValue only compare value,
// return -100 if *KV does not implement ValueStore
func (kv *KV) CompareValue(value Value) int {
	kv1, ok := value.(*KV)

	if !ok {
		return -100
	}

	if kv.Value > kv1.Value {
		return 1
	} else if kv.Value < kv1.Value {
		return -1
	}
	return 0
}

func (kv *KV) String() string {
	return fmt.Sprintf("key: %d, value: %d", kv.Key, kv.Value)
}

var _ Value = &Ivalue{}

// Ivalue ...
type Ivalue struct {
	Value int
}

// Compare call CompareValue
func (iv *Ivalue) Compare(value Value) int {
	return iv.CompareValue(value)
}

// CompareValue return 1 if v.Value > vs.value,
// return -100 if *KV does not implement ValueStore
func (iv *Ivalue) CompareValue(value Value) int {
	iv1, ok := value.(*Ivalue)

	if !ok {
		return -100
	}

	if iv.Value > iv1.Value {
		return 1
	} else if iv.Value < iv1.Value {
		return -1
	}
	return 0
}

func (iv *Ivalue) String() string {
	return fmt.Sprintf("%d", iv.Value)
}
