package image

import (
	"testing"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_implicitReference_TagName(t *testing.T) {
	tests := []struct {
		name  string
		image string
		want  string
	}{
		{
			name:  "explicit tag",
			image: "khulnasoft/tunnel:0.15.0",
			want:  "0.15.0",
		},
		{
			name:  "implicit tag",
			image: "khulnasoft/tunnel",
			want:  "latest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := name.ParseReference(tt.image)
			require.NoError(t, err)

			ref := implicitReference{ref: r}

			got := ref.TagName()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_implicitReference_RepositoryName(t *testing.T) {
	tests := []struct {
		name  string
		image string
		want  string
	}{
		{
			name:  "explicit default registry",
			image: "index.docker.io/khulnasoft/tunnel:0.15.0",
			want:  "khulnasoft/tunnel",
		},
		{
			name:  "explicit default namespace",
			image: "library/alpine:3.12",
			want:  "alpine",
		},
		{
			name:  "implicit registry",
			image: "khulnasoft/tunnel:latest",
			want:  "khulnasoft/tunnel",
		},
		{
			name:  "implicit namespace",
			image: "alpine:latest",
			want:  "alpine",
		},
		{
			name:  "non-default registry",
			image: "gcr.io/library/alpine:3.12",
			want:  "gcr.io/library/alpine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := name.ParseReference(tt.image)
			require.NoError(t, err)

			ref := implicitReference{ref: r}

			got := ref.RepositoryName()
			assert.Equal(t, tt.want, got)
		})
	}
}
