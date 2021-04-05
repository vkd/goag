package generator

import "testing"

func TestPublicFieldName(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{"petId", "PetID"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.arg, func(t *testing.T) {
			if got := PublicFieldName(tt.arg); got != tt.want {
				t.Errorf("PublicFieldName() = %v, want %v", got, tt.want)
			}
		})
	}
}
