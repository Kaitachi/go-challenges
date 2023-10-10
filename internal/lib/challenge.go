package lib

import (
	"embed"
	"io/fs"
	"path/filepath"
	"strings"
)

type Challenge struct {
	Assets		*embed.FS
	AssetPath	string
	Solution	string
	Algorithms []*string
}

// Kudos to @mrsoftware for the function below!
// Code adapted for use case
// https://gist.github.com/clarkmcc/1fdab4472283bb68464d066d6b4169bc?permalink_comment_id=4405804#gistcomment-4405804
func (c Challenge) GetFiles() (files []string, err error) {
	if err := fs.WalkDir(c.Assets, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
 
		if !strings.HasPrefix(path, c.AssetPath) {
			return nil
		}

		if !strings.HasPrefix(filepath.Base(path), c.Solution) {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}

