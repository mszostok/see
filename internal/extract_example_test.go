package internal_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.szostok/see/internal"
)

func TestSimpleExample(t *testing.T) {
	// given
	expExamples := []internal.Example{
		{
			Description: "",
			Command:     "kubectl alpha events -n foo",
		},
		{
			Description: " Update pod 'foo' with the annotation 'description' and the value 'my frontend' If the same annotation is set multiple times, only the last value will be applied",
			Command:     "kubectl annotate pods foo description='my frontend'",
		},
		{
			Description: " Update a pod identified by type and name in \"pod.json\"",
			Command:     "kubectl annotate -f pod.json description='my frontend'",
		},
		{
			Description: " Update pod 'foo' with the annotation 'description' and the value 'my frontend running nginx', overwriting any existing value",
			Command:     "kubectl annotate --overwrite pods foo description='my frontend running nginx'",
		},
		{
			Description: " Update all pods in the namespace",
			Command:     "kubectl annotate pods --all description='my frontend running nginx'",
		},
		{
			Description: " Update pod 'foo' only if the resource is unchanged from version 1",
			Command:     "kubectl annotate pods foo description='my frontend running nginx' --resource-version=1",
		},
		{
			Description: " Update pod 'foo' by removing an annotation named 'description' if it exists Does not require the --overwrite flag",
			Command:     "kubectl annotate pods foo description-## SEE ALSO",
		},
	}

	rawExample, err := os.ReadFile(filepath.Join("testdata", "example.md"))
	require.NoError(t, err)

	// when
	gotExamples := internal.ProcessSimpleExamples(string(rawExample))

	// then
	litter.Dump(gotExamples)
	assert.Equal(t, expExamples, gotExamples)
}
