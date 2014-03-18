// Command docker-rebase is used to "rebase" a Docker image against an base
// image. The resulting image contains only those layers which are unique to the
// downstream image.
//
// Example
//
//     docker save base > base.tar
//     docker save app | docker-rebase base.tar > app.tar
package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no base image specified")
	}

	excluded, err := ls(os.Args[1])
	if err != nil {
		log.Panic(err)
	}

	if err := rebase(os.Stdin, os.Stdout, excluded); err != nil {
		log.Panic(err)
	}
}

func rebase(in io.Reader, out io.Writer, excluded map[string]bool) error {
	r := tar.NewReader(in)

	w := tar.NewWriter(out)
	defer w.Close()

	for {
		h, err := r.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if !excluded[h.Name] {
			if err := w.WriteHeader(h); err != nil {
				return err
			}

			if _, err := io.Copy(w, r); err != nil {
				return err
			}
		}
	}

	return nil
}

func ls(path string) (map[string]bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	files := make(map[string]bool)
	r := tar.NewReader(f)
	for {
		h, err := r.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		files[h.Name] = true
	}

	// don't exclude the metadata
	delete(files, "./")
	delete(files, "repositories")

	return files, nil
}
