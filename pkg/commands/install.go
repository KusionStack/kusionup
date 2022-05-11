package commands

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func installCmd() *cobra.Command {
	installCmd := &cobra.Command{
		Use:     "install [VERSION]",
		Aliases: []string{"i"},
		Short:   `Install Kusion with a version, alias: "i"`,
		Long:    `Install Kusion by providing a version. If no version is provided, install the latest Kusion (internal@latest).`,
		Example: `
  kusionup install
  kusionup install internal@1.15.2
`,
		PersistentPreRunE: preRunInstall,
		RunE:              runInstall,
	}

	return installCmd
}

func preRunInstall(_ *cobra.Command, _ []string) error {
	http.DefaultTransport = &userAgentTransport{http.DefaultTransport}
	return nil
}

func runInstall(_ *cobra.Command, args []string) error {
	var (
		sourceVer string
		err       error
	)

	if len(args) == 0 {
		sourceVer = latestKusionSourceVersion()
	} else {
		sourceVer = args[0]
	}

	source, ver, err := ParseSourceVersion(sourceVer)
	if err != nil {
		return err
	}

	err = install(source, ver)
	if err != nil {
		return err
	}

	return switchVer(GetSourceVersionTitle(source, ver))
}

func latestKusionSourceVersion() string {
	return "internal@latest"
}

func switchVer(verSuffix string) error {
	err := symlink(verSuffix)

	if err == nil {
		logger.Printf("Default Kusion is set to '%s'", verSuffix)
	}

	return err
}

func symlink(verSuffix string) error {
	current := KusionupCurrentDir()
	versionDir := KusionupVersionDir(verSuffix)

	if _, err := os.Stat(versionDir); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("kusion version %s is not installed. Install it with `kusionup install`", verSuffix)
		}

		return err
	}

	// ignore error, similar to rm -f
	os.Remove(current)

	return os.Symlink(versionDir, current)
}

func install(source, version string) error {
	// Mkdir or return when existed
	targetDir := KusionupSourceVersionDir(source, version)
	if _, err := os.Stat(filepath.Join(targetDir, unpackedOkay)); err == nil {
		logger.Printf("%s: already downloaded in %v", version, targetDir)

		return nil
	}

	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}

	// Get download url
	rs := registedReleaseSources[source]

	downloadURL, err := rs.GetDownloadURL(version)
	if err != nil {
		return err
	}

	// Head download url
	res, err := http.Head(downloadURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return fmt.Errorf("no binary release of %v for %v/%v at %v", version, getOS(), runtime.GOARCH, downloadURL)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %v checking size of %v", http.StatusText(res.StatusCode), downloadURL)
	}

	// Download
	base := path.Base(downloadURL)
	archiveFile := filepath.Join(targetDir, base)

	if fi, err := os.Stat(archiveFile); err != nil || fi.Size() != res.ContentLength {
		if err != nil && !os.IsNotExist(err) {
			// Something weird. Don't try to download.
			return err
		}

		if err := copyFromURL(archiveFile, downloadURL); err != nil {
			return fmt.Errorf("error downloading %v: %v", downloadURL, err)
		}

		fi, err = os.Stat(archiveFile)
		if err != nil {
			return err
		}

		if fi.Size() != res.ContentLength {
			return fmt.Errorf("downloaded file %s size %v doesn't match server size %v", archiveFile, fi.Size(), res.ContentLength)
		}
	}

	// TODO: check checksum
	// wantSHA, err := slurpURLToString(downloadURL + ".sha256")
	// if err != nil {
	// 	return err
	// }
	// if err := verifySHA256(archiveFile, strings.TrimSpace(wantSHA)); err != nil {
	// 	return fmt.Errorf("error verifying SHA256 of %v: %v", archiveFile, err)
	// }

	logger.Printf("Unpacking %v ...", archiveFile)

	if err := unpackArchive(targetDir, archiveFile); err != nil {
		return fmt.Errorf("extracting archive %v: %v", archiveFile, err)
	}

	if err := ioutil.WriteFile(filepath.Join(targetDir, unpackedOkay), nil, 0o644); err != nil {
		return err
	}

	logger.Printf("Success: %s downloaded in %v", version, targetDir)

	return nil
}

// unpackArchive unpacks the provided archive zip or tar.gz file to targetDir,
// removing the "go/" prefix from file entries.
func unpackArchive(targetDir, archiveFile string) error {
	switch {
	case strings.HasSuffix(archiveFile, ".zip"):
		return unpackZip(targetDir, archiveFile)
	case strings.HasSuffix(archiveFile, ".tar.gz") || strings.HasSuffix(archiveFile, ".tgz"):
		return unpackTarGz(targetDir, archiveFile)
	default:
		return errors.New("unsupported archive file")
	}
}

// unpackTarGz is the tar.gz implementation of unpackArchive.
func unpackTarGz(targetDir, archiveFile string) error {
	r, err := os.Open(archiveFile)
	if err != nil {
		return err
	}
	defer r.Close()

	madeDir := map[string]bool{}

	zr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	tr := tar.NewReader(zr)

	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if !validRelPath(f.Name) {
			return fmt.Errorf("tar file contained invalid name %q", f.Name)
		}

		rel := filepath.FromSlash(strings.TrimPrefix(f.Name, "go/"))
		abs := filepath.Join(targetDir, rel)

		fi := f.FileInfo()
		mode := fi.Mode()

		switch {
		case mode.IsRegular():
			// Make the directory. This is redundant because it should
			// already be made by a directory entry in the tar
			// beforehand. Thus, don't check for errors; the next
			// write will fail with the same error.
			dir := filepath.Dir(abs)
			if !madeDir[dir] {
				if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
					return err
				}

				madeDir[dir] = true
			}

			wf, err := os.OpenFile(abs, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode.Perm())
			if err != nil {
				return err
			}

			n, err := io.Copy(wf, tr)
			if closeErr := wf.Close(); closeErr != nil && err == nil {
				err = closeErr
			}

			if err != nil {
				return fmt.Errorf("error writing to %s: %v", abs, err)
			}

			if n != f.Size {
				return fmt.Errorf("only wrote %d bytes to %s; expected %d", n, abs, f.Size)
			}

			if !f.ModTime.IsZero() {
				if err := os.Chtimes(abs, f.ModTime, f.ModTime); err != nil {
					// benign error. Gerrit doesn't even set the
					// modtime in these, and we don't end up relying
					// on it anywhere (the gomote push command relies
					// on digests only), so this is a little pointless
					// for now.
					logger.Printf("error changing modtime: %v", err)
				}
			}
		case mode.IsDir():
			if err := os.MkdirAll(abs, 0o755); err != nil {
				return err
			}

			madeDir[abs] = true
		default:
			return fmt.Errorf("tar file entry %s contained unsupported file type %v", f.Name, mode)
		}
	}

	return nil
}

// unpackZip is the zip implementation of unpackArchive.
func unpackZip(targetDir, archiveFile string) error {
	zr, err := zip.OpenReader(archiveFile)
	if err != nil {
		return err
	}
	defer zr.Close()

	for _, f := range zr.File {
		name := strings.TrimPrefix(f.Name, "go/")

		outpath := filepath.Join(targetDir, name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(outpath, 0o755); err != nil {
				return err
			}

			continue
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		// File
		if err := os.MkdirAll(filepath.Dir(outpath), 0o755); err != nil {
			return err
		}

		out, err := os.OpenFile(outpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(out, rc)
		rc.Close()

		if err != nil {
			out.Close()
			return err
		}

		if err := out.Close(); err != nil {
			return err
		}
	}

	return nil
}

// copyFromURL downloads srcURL to dstFile.
func copyFromURL(dstFile, srcURL string) (err error) {
	f, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			f.Close()
			os.Remove(dstFile)
		}
	}()

	c := &http.Client{
		Transport: &userAgentTransport{&http.Transport{
			// It's already compressed. Prefer accurate ContentLength.
			// (Not that GCS would try to compress it, though)
			DisableCompression: true,
			DisableKeepAlives:  true,
			Proxy:              http.ProxyFromEnvironment,
		}},
	}

	res, err := c.Get(srcURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	pw := &progressWriter{w: f, total: res.ContentLength}

	n, err := io.Copy(pw, res.Body)
	if err != nil {
		return err
	}

	if res.ContentLength != -1 && res.ContentLength != n {
		return fmt.Errorf("copied %v bytes; expected %v", n, res.ContentLength)
	}

	pw.update() // 100%

	return f.Close()
}

type progressWriter struct {
	w     io.Writer
	n     int64
	total int64
	last  time.Time
}

func (p *progressWriter) update() {
	end := " ..."
	if p.n == p.total {
		end = ""
	}

	fmt.Fprintf(os.Stderr, "Downloaded %5.1f%% (%*d / %d bytes)%s\n",
		(100.0*float64(p.n))/float64(p.total),
		ndigits(p.total), p.n, p.total, end)
}

func ndigits(i int64) int {
	var n int
	for ; i != 0; i /= 10 {
		n++
	}

	return n
}

func (p *progressWriter) Write(buf []byte) (n int, err error) {
	n, err = p.w.Write(buf)
	p.n += int64(n)

	if now := time.Now(); now.Unix() != p.last.Unix() {
		p.update()
		p.last = now
	}

	return
}

// getOS returns runtime.GOOS. It exists as a function just for lazy
// testing of the Windows zip path when running on Linux/Darwin.
func getOS() string {
	return runtime.GOOS
}

// unpackedOkay is a sentinel zero-byte file to indicate that the Kusion
// version was downloaded and unpacked successfully.
const unpackedOkay = ".unpacked-success"

func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}

	return true
}

type userAgentTransport struct {
	rt http.RoundTripper
}

func (uat userAgentTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	version := runtime.Version()
	if strings.Contains(version, "devel") {
		// Strip the SHA hash and date. We don't want spaces or other tokens (see RFC2616 14.43)
		version = "devel"
	}

	r.Header.Set("User-Agent", "kusionup/"+version)

	return uat.rt.RoundTrip(r)
}
