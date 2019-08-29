package dep

import (
	"fmt"
	"strings"

	"github.com/jasinner/depcheck/pkg/managers/version"
)

func ParseManifest(manifest map[string][]byte) ([]version.Dependency, error) {
	lock, err := readLock(manifest["Gopkg.lock"])
	if err != nil {
		fmt.Println("Skipping dependencies in unparsable Gopkg.lock")
		return nil, nil
	}
	list := []version.Dependency{}
	for _, p := range lock.Projects {
		digest := strings.TrimPrefix(p.Digest, "1:")
		if len(digest) > 8 {
			digest = digest[0:8]
		}
		list = append(list, version.Dependency{
			Name:       p.Name,
			Version:    p.Version,
			Digest:     digest,
			Repository: p.Source,
		})
	}
	return list, nil
}
