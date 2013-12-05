package crucible

import (
  "net/http"
  "archive/zip"
  "strings"
  "path/filepath"
  "io/ioutil"
  "io"
  "os"
  "fmt"
)

const (
  githubUrl string   = "https://api.github.com/repos"
  githubPackage string = "zipball/master"
)

func createDir(home string, playbook string, path string) {
  err := os.MkdirAll(filepath.Join(home, crucibleDir, playbook, path), 0755)
  HandleErr(err);
}

func createFile(home string, playbook string, path string, file *zip.File) {
  content, err := file.Open(); HandleErr(err);
  defer content.Close()

  dst, err := os.OpenFile(filepath.Join(home, crucibleDir, playbook, path), os.O_CREATE|os.O_RDWR, 0644); HandleErr(err);
  defer dst.Close()

  _, err = io.Copy(dst, content); HandleErr(err);
}

func DownloadPlayBook(playbook string) {
  fmt.Printf("Downloading %s...",playbook)

  url := strings.Join([]string{githubUrl, playbook, githubPackage}, "/")

  resp, err := http.Get(url); HandleErr(err);
  defer resp.Body.Close()

  file, err := ioutil.TempFile(os.TempDir(), "foo"); HandleErr(err);

  io.Copy(file, resp.Body)

  r, err := zip.OpenReader(file.Name()); HandleErr(err);
  defer r.Close()

  home := HomeDir()
  createPlayBookTree(home, playbook)

  for _, f := range r.File[1:] {
    path := strings.Split(f.Name,string(os.PathSeparator))[1:]

    if filepath.Dir(f.Name) == filepath.Clean(f.Name) {
      createDir(home, playbook, filepath.Join(path...))
    } else {
      createFile(home, playbook, filepath.Join(path...), f)
    }
  }

  fmt.Println("complete!")
}

func createPlayBookTree(home string, playBook string) {
  params := strings.Split(playBook, "/")
  playBookDir := filepath.Join(home, crucibleDir, params[0], params[1])
  err := os.MkdirAll(playBookDir, 0755); HandleErr(err);
}
