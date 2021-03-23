/*
 * emptyrepo.go
 *
 * Copyright 2021 Bill Zissimopoulos
 */
/*
 * This file is part of Hubfs.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

package providers

import (
	"io"
)

// When using:
//
//     var emptyRepository Repository = &emptyRepositoryT{}
//
// The emptyRepository variable is not initialized (at least during testing).
// May be related to https://github.com/golang/go/issues/44956
var emptyRepository Repository

type emptyRepositoryT struct {
}

func (*emptyRepositoryT) Close() (err error) {
	return nil
}

func (*emptyRepositoryT) SetDirectory(path string) error {
	return nil
}

func (*emptyRepositoryT) RemoveDirectory() error {
	return nil
}

func (*emptyRepositoryT) Name() string {
	return ""
}

func (*emptyRepositoryT) GetRefs() ([]Ref, error) {
	return []Ref{}, nil
}

func (*emptyRepositoryT) GetRef(name string) (Ref, error) {
	return nil, ErrNotFound
}

func (*emptyRepositoryT) GetTree(ref Ref, entry TreeEntry) ([]TreeEntry, error) {
	return []TreeEntry{}, nil
}

func (*emptyRepositoryT) GetTreeEntry(ref Ref, entry TreeEntry, name string) (TreeEntry, error) {
	return nil, ErrNotFound
}

func (*emptyRepositoryT) GetBlobReader(entry TreeEntry) (io.ReaderAt, error) {
	return nil, ErrNotFound
}

func init() {
	emptyRepository = &emptyRepositoryT{}
}