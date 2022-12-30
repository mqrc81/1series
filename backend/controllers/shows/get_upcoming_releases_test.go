package shows

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculateRange(t *testing.T) {
	type args struct {
		page         int
		pastReleases int
	}
	type want struct {
		amount          int
		offset          int
		possiblyHasMore bool
	}
	tests := []struct {
		name string
		args
		want
	}{
		{
			name: "#1",
			args: args{page: 1, pastReleases: 55},
			want: want{amount: 20, offset: 55, possiblyHasMore: true},
		},
		{
			name: "#2",
			args: args{page: 2, pastReleases: 55},
			want: want{amount: 20, offset: 75, possiblyHasMore: true},
		},
		{
			name: "#3",
			args: args{page: -1, pastReleases: 55},
			want: want{amount: 20, offset: 35, possiblyHasMore: true},
		},
		{
			name: "#4",
			args: args{page: -3, pastReleases: 55},
			want: want{amount: 15, offset: 0, possiblyHasMore: false},
		},
		{
			name: "#5",
			args: args{page: 4, pastReleases: 55},
			want: want{amount: 20, offset: 115, possiblyHasMore: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAmount, gotOffset, gotPossiblyHasMore := calculateRange(tt.args.page, tt.args.pastReleases)

			assert.Equal(t, tt.amount, gotAmount)
			assert.Equal(t, tt.offset, gotOffset)
			assert.Equal(t, tt.possiblyHasMore, gotPossiblyHasMore)

		})
	}
}
