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
		{"-,A-_ c3", "AC"},
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
