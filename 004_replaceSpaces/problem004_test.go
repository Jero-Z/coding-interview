package replace

import (
	"reflect"
	"testing"
)

func Test_replaceStr(t *testing.T) {
	type args struct {
		subject    []byte
		charLength int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "empty string", args: args{subject: []byte{'a', ' ', 'b', ' ', 'c', 'x', 'x', 'x', 'x'}, charLength: 5}, want: []byte{'a', '%', '2', '0', 'b', '%', '2', '0', 'c'}},
		{name: "string with only spaces", args: args{subject: []byte{' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x'}, charLength: 5}, want: []byte{'%', '2', '0', '%', '2', '0', '%', '2', '0', '%', '2', '0', 'x', 'x', 'x', 'x'}},
		{name: "string with only x's", args: args{subject: []byte{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'}, charLength: 5}, want: []byte{'x', 'x', 'x', 'x', 'x', '%', '2', '0', '%', '2', '0', '%', '2', '0', '%', '2', '0'}},
		// 新的测试用例
		{name: "string with spaces at the beginning and end", args: args{subject: []byte{' ', ' ', 'a', ' ', 'b', ' ', ' ', ' '}, charLength: 5}, want: []byte{'%', '2', '0', '%', '2', '0', 'a', '%', '2', '0', 'b', '%', '2', '0', '%', '2', '0'}},
		{name: "string with spaces in the middle", args: args{subject: []byte{'a', ' ', ' ', 'b', ' ', 'c', 'x', 'x', 'x'}, charLength: 5}, want: []byte{'a', '%', '2', '0', ' ', '%', '2', '0', 'b', '%', '2', '0', 'c', 'x', 'x', 'x'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceStr(tt.args.subject, tt.args.charLength)
			if reflect.DeepEqual(tt.args.subject, tt.want) {
				t.Errorf("replace() = %v, want %v", tt.args.subject, tt.want)
			}
		})
	}
}
