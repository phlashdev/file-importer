package diff

import "testing"

func TestNewDirectoryDiffer(t *testing.T) {
	data := []struct {
		name         string
		src          string
		dest         string
		expectedSrc  string
		expectedDest string
		err          error
	}{
		{"no_src", "", "/dest", "", "", ErrNoSrc},
		{"no_dest", "/src", "", "", "", ErrNoDest},
		{"created", "/src", "/dest", "/src", "/dest", nil},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			differ, err := NewDirectoryDiffer(d.src, d.dest)
			if err != d.err {
				t.Errorf("Expected error %v, got %v", d.err, err)
			}
			expected := directoryDiffer{src: d.expectedSrc, dest: d.expectedDest}
			if differ != expected {
				t.Errorf("Expected %v, got %v", expected, differ)
			}
		})
	}
}
