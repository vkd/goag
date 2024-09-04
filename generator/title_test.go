package generator

import "testing"

func TestTitle(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"ShopName", "ShopName"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := Title(tt.s); got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}
