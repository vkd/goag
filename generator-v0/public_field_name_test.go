package generator

import "testing"

func TestPublicFieldName(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{"petId", "PetID"},
		{"petID", "PetID"},
		{"pe-tId", "PeTID"},
		{"pe-Id", "PeID"},
		{"apple/orange", "AppleOrange"},
		{"ABC", "ABC"},
		{"AaBbCc", "AaBbCc"},
		{"A-aB_bC c", "AABBCC"},
		{"A-_ c", "AC"},
		{"-,A-_ c3", "AC3"},
		{"c3", "C3"},
		{"3c", "C"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.arg, func(t *testing.T) {
			if got := PublicFieldName(tt.arg); got != tt.want {
				t.Errorf("PublicFieldName(%q) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}
