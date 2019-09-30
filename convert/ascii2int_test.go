package convert

import "testing"

func TestAsciiToInt(t *testing.T) {
	i, err := AsciiToInt([]byte{'1', '2'})
	if err != nil {
		t.Fatal()
	}
	if i != 12 {
		t.Fatal()
	}

	i, err = AsciiToInt([]byte{'1', '2', '3', '4'})
	if err != nil {
		t.Fatal()
	}
	if i != 1234 {
		t.Fatal()
	}
}
