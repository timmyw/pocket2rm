package pocket2rm

import "testing"

func TestFixForFileName(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ "all_legal", args{ "sample.output.json"}, "sample.output.json" },
		{ "some_legal", args{ "sample.!utpu*.json"}, "sample._utpu_.json" },
		{ "none_legal", args{ "*\"^$'"}, "_____" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixForFileName(tt.args.inp); got != tt.want {
				t.Errorf("FixForFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
