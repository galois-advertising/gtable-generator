// solopointer1202@gmail.com
// 20200109
package generator 

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func Cp(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	log.Printf("cp [%s] to [%s]", src, dst)
	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func DirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func archive_file_from_localfile_cmd(url, base_dir string) (string, string, error) {
	path := strings.Replace(url, "local://", fmt.Sprintf("%s/", base_dir), 1)
	target_file := fmt.Sprintf("%s/.temp/%s", base_dir, path)
	export_cmd := fmt.Sprintf("cp %s %s", url, target_file)
	return target_file, export_cmd, nil
}

func archive_file_from_gitlab_cmd(url, base_dir string) (string, string, error) {
	matched, err := regexp.Match("https://github.com/(.+)", []byte(url))
	if err != nil {
		log.Println(matched, err)
	}
	log.Println(matched)
	return "asdf", "sadf", nil
	//if m is not None:
	//    spl = m.group(1).split('/')
	//    account = spl[0]
	//    project = spl[1]
	//    branch = spl[3]
	//    path = '/'.join(spl[4:])
	//    clone_url = "ssh://git@github.com/%s/%s.git" % (account, project)
	//    clone_path = "%s/.temp" % (base_dir)
	//    target_file = "%s/.temp/%s" % (base_dir, path)
	//    export_cmd = "git clone --depth=1 --branch %s --single-branch %s %s" % (
	//        branch, clone_url, clone_path)
	//    return export_cmd, target_file
}

func ArchiveFile(url, base_dir string) (target_path string, e error) {
	var cmd, tareget_path string
	var err error = nil
	if strings.HasPrefix(url, "local://") {
		log.Println("processing local file.")
		cmd, target_path, err = archive_file_from_localfile_cmd(url, base_dir)
	} else if strings.HasPrefix(url, "https://") {
		log.Println("processing remote file.")
		cmd, target_path, err = archive_file_from_gitlab_cmd(url, base_dir)
	}
	log.Println(cmd, tareget_path, err)
	return
}
