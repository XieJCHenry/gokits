package maps

import (
	"fmt"
	"github.com/XieJCHenry/gokits/collections/tuple"
	"reflect"
	"testing"
)

func TestGetKeys(t *testing.T) {
	type args[KEY comparable, VALUE any] struct {
		m map[KEY]VALUE
	}
	type testCase[KEY comparable, VALUE any] struct {
		name string
		args args[KEY, VALUE]
		want []KEY
	}
	tests := []testCase[string, int]{
		{
			name: "normal",
			args: args[string, int]{
				m: map[string]int{
					"key1": 1,
				},
			},
			want: []string{"key1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKeys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValues(t *testing.T) {
	type args[KEY comparable, VALUE any] struct {
		m map[KEY]VALUE
	}
	type testCase[KEY comparable, VALUE any] struct {
		name string
		args args[KEY, VALUE]
		want []VALUE
	}
	tests := []testCase[string, int]{
		{
			name: "normal",
			args: args[string, int]{
				m: map[string]int{
					"key1": 1,
					"key2": 1,
				},
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValues(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromTuples(t *testing.T) {
	type args struct {
		tuples []tuple.Tuple
	}
	type testCase[KEY comparable, VALUE any] struct {
		name    string
		args    args
		want    map[KEY]VALUE
		wantErr bool
	}
	tests := []testCase[string, int]{
		{
			name: "empty",
			args: args{
				tuples: nil,
			},
			want:    make(map[string]int),
			wantErr: false,
		},
		{
			name: "tuple size is invalid",
			args: args{
				tuples: []tuple.Tuple{
					tuple.NewFrom(1, 2, 3),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				tuples: []tuple.Tuple{
					tuple.NewFrom("key1", 1),
					tuple.NewFrom("key2", 2),
					tuple.NewFrom("key3", 3),
					tuple.NewFrom("key4", 4),
					tuple.NewFrom("key4", 4),
				},
			},
			want: map[string]int{
				"key1": 1,
				"key2": 2,
				"key3": 3,
				"key4": 4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromTuples[string, int](tt.args.tuples)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromTuples() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromTuples() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromStruct(t *testing.T) {
	type args[VALUE any, KEY comparable] struct {
		elements     []VALUE
		keyGenerator func(elem VALUE) (KEY, bool)
	}
	type testCase[VALUE any, KEY comparable] struct {
		name string
		args args[VALUE, KEY]
		want map[KEY]VALUE
	}
	tests := []testCase[struct {
		ID  string
		Num int
	}, string]{
		{
			name: "normal",
			args: args[struct {
				ID  string
				Num int
			}, string]{
				elements: []struct {
					ID  string
					Num int
				}{
					{
						ID:  "id1",
						Num: 1,
					},
					{
						ID:  "id2",
						Num: 2,
					},
				},
				keyGenerator: func(elem struct {
					ID  string
					Num int
				}) (string, bool) {
					return fmt.Sprintf("%s-%d", elem.ID, elem.Num), true
				},
			},
			want: map[string]struct {
				ID  string
				Num int
			}{
				"id1-1": {
					ID:  "id1",
					Num: 1,
				},
				"id2-2": {
					ID:  "id2",
					Num: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromStruct(tt.args.elements, tt.args.keyGenerator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromKeyValues(t *testing.T) {
	type args[KEY comparable, VALUE any] struct {
		keys   []KEY
		values []VALUE
	}
	type testCase[KEY comparable, VALUE any] struct {
		name string
		args args[KEY, VALUE]
		want map[KEY]VALUE
	}
	tests := []testCase[string, int]{
		{
			name: "empty",
			args: args[string, int]{
				keys:   []string{},
				values: []int{},
			},
			want: make(map[string]int),
		},
		{
			name: "normal",
			args: args[string, int]{
				keys:   []string{"key1", "key2"},
				values: []int{1, 2},
			},
			want: map[string]int{
				"key1": 1,
				"key2": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromKeyValues(tt.args.keys, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromKeyValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
