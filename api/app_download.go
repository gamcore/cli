package api

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
)

func (a App) doDownloadAndExtract(link, actualPath string) error {
	mf, err := a.Manifest()
	if err != nil {
		return err
	}
	file := tempFile(link[strings.LastIndex(link, "/")+1:])
	if mf.Updates.OneFile {
		actualPath = path.Join(actualPath, mf.Executable[0])
		if runtime.GOOS == "windows" {
			actualPath = actualPath + ".exe"
		}
	}
	res, err := httpClient.Get(link)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	bar := progressbar.DefaultBytes(res.ContentLength, fmt.Sprintf(`downloading "%s"`, a.Name))
	if err != nil {
		return err
	}
	outFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeType)
	if err != nil {
		return err
	}
	_, err = io.Copy(io.MultiWriter(bar, outFile), res.Body)
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	return extract(file, actualPath, mf.Updates.OneFile)
}

func tempFile(filename string) string {
	return path.Join(tempPath(), filename)
}

func extract(archive, out string, oneFile bool) error {
	z, err := os.Open(archive)
	if err != nil {
		return err
	}
	defer z.Close()
	switch {
	case strings.HasSuffix(path.Base(archive), ".zip"):
		return zipExtract(archive, out)
	case strings.HasSuffix(path.Base(archive), ".gz"):
		return gzipExtract(z, out)
	case strings.HasSuffix(path.Base(archive), ".bz"):
		return bzipExtract(z, out)
	case strings.HasSuffix(path.Base(archive), ".tar"):
		return tarExtract(z, out)
	case oneFile:
		return clone(archive, out)
	default:
		return ErrUnsupportedArchive
	}
}

func clone(s, t string) error {
	src, err := os.Open(s)
	if err != nil {
		return err
	}
	target, err := os.Open(t)
	if err != nil {
		return err
	}
	_, err = io.Copy(src, target)
	return err
}

func zipExtract(archive, out string) error {
	z, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer z.Close()

	for _, f := range z.File {
		fp := path.Join(out, f.Name)

		if !strings.HasPrefix(fp, path.Clean(out)+string(os.PathSeparator)) {
			return ErrArchiveInvalidPath
		}
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fp, f.Mode())
			continue
		}
		if err = os.MkdirAll(path.Dir(fp), os.ModeDir); err != nil {
			return err
		}
		dstFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}
		if _, err = io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}
		dstFile.Close()
		fileInArchive.Close()
	}

	return nil
}

func gzipExtract(archive *os.File, out string) error {
	z, err := gzip.NewReader(archive)
	if err != nil {
		return err
	}
	defer z.Close()

	if strings.HasSuffix(archive.Name(), ".tar.gz") {
		return tarExtract(z, out)
	} else {
		wr, err := os.Create(out)
		if err != nil {
			return err
		}
		defer wr.Close()
		_, err = io.Copy(wr, z)
		return err
	}
}

func bzipExtract(archive *os.File, out string) error {
	z := bzip2.NewReader(archive)
	if strings.HasSuffix(archive.Name(), ".tar.bz") {
		return tarExtract(z, out)
	} else {
		wr, err := os.Create(out)
		if err != nil {
			return err
		}
		defer wr.Close()
		_, err = io.Copy(wr, z)
		return err
	}
}

func tarExtract(archive io.Reader, out string) error {
	z := tar.NewReader(archive)

	for header, err := z.Next(); err != io.EOF; {
		if err != nil {
			return err
		}

		fp := path.Join(out, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(fp, info.Mode()); err != nil {
				return err
			}
			continue
		}

		cf, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
		if err != nil {
			return err
		}
		_, err = io.Copy(cf, z)
		cf.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
