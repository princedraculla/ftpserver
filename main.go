package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/goftp/server"
)

func main() {
	cdrFile := &CDRFile{}
	factory := &DriverFactory{
		RootPath: "/home/amir/ftp_files",
	}

	opts := &server.ServerOpts{
		Factory:        factory,
		Hostname:       "127.0.0.1",
		Port:           2121,
		Auth:           &Auth{},
		TLS:            true,
		CertFile:       "/home/amir/server.crt",
		KeyFile:        "/home/amir/server.key",
		ExplicitFTPS:   true,
		PassivePorts:   "30000-30009",
		WelcomeMessage: "Welcome to FTPS server",
	}

	ftpServer := server.NewServer(opts)
	log.Println("Starting FTP server on :2121")
	err := ftpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	files, err := DatFiles()
	if err != nil {
		log.Fatal("found error : ", err)
	}
	for _, file := range files {
		cdr, err := cdrFile.DecodeCDRFile(file)
		if err != nil {
			log.Fatal("error while decoding cdr : ", err)
		}
		fmt.Println("the decoded CDR : ", cdr)
	}
}

// DriverFactory implements server.DriverFactory
type DriverFactory struct {
	RootPath string
}

func (f *DriverFactory) NewDriver() (server.Driver, error) {
	return &Driver{
		RootPath: f.RootPath,
	}, nil
}

// Driver implements server.Driver interface
type Driver struct {
	RootPath string
}

// Rename implements server.Driver.
func (d *Driver) Rename(string, string) error {
	panic("unimplemented")
}

func (d *Driver) realPath(path string) string {
	fmt.Printf("realPath function ==> %+v\n", filepath.Join(d.RootPath, path))
	return filepath.Join(d.RootPath, path)
}

func (d *Driver) Init(conn *server.Conn) {}

func (d *Driver) Stat(path string) (server.FileInfo, error) {
	absPath := d.realPath(path)
	f, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}
	return &FileInfo{f}, nil
}

func (d *Driver) ChangeDir(path string) error {
	absPath := d.realPath(path)
	_, err := os.Stat(absPath)
	return err
}

func (d *Driver) ListDir(path string, callback func(server.FileInfo) error) error {
	absPath := d.realPath(path)
	files, err := os.ReadDir(absPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}
		if err := callback(&FileInfo{info}); err != nil {
			return err
		}
	}
	return nil
}

func (d *Driver) DeleteDir(path string) error {
	return os.RemoveAll(d.realPath(path))
}

func (d *Driver) DeleteFile(path string) error {
	return os.Remove(d.realPath(path))
}

func (d *Driver) MakeDir(path string) error {
	return os.MkdirAll(d.realPath(path), 0755)
}

func (d *Driver) GetFile(path string, offset int64) (int64, io.ReadCloser, error) {
	absPath := d.realPath(path)
	f, err := os.Open(absPath)
	if err != nil {
		return 0, nil, err
	}

	info, err := f.Stat()
	if err != nil {
		return 0, nil, err
	}

	if offset > 0 {
		_, err = f.Seek(offset, io.SeekStart)
		if err != nil {
			return 0, nil, err
		}
	}

	return info.Size(), f, nil
}

func (d *Driver) PutFile(path string, data io.Reader, appendData bool) (int64, error) {
	absPath := d.realPath(path)
	cdrFile := &CDRFile{}

	fmt.Println("the absolute path for the file : ", absPath)

	var f *os.File
	var err error

	if appendData {
		f, err = os.OpenFile(absPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return 0, nil
		}
	} else {
		f, err = os.Create(absPath)
		if err != nil {
			return 0, nil
		}
	}

	defer f.Close()

	n, err := io.Copy(f, data)
	if err != nil {
		return 0, err
	}
	cdr, err := cdrFile.DecodeCDRFile(absPath)
	if err != nil {
		return 0, err
	}
	fmt.Println("[Decoded CDR File] : ", cdr)
	return n, nil
}

// FileInfo implements server.FileInfo
type FileInfo struct {
	os.FileInfo
}

func (f *FileInfo) Owner() string {
	return "0"
}

func (f *FileInfo) Group() string {
	return "0"
}

func (f *FileInfo) ModTime() time.Time {
	return f.FileInfo.ModTime()
}

// Auth implements server.Auth interface
type Auth struct{}

func (a *Auth) CheckPasswd(username, password string) (bool, error) {
	return username == "user" && password == "pass", nil
}
