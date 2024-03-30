package tag

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

// 创建自增tag
func IncrementTagVersion() {
	latestVersion := getLatestTag()
	newVersion := incrementVersion(latestVersion)

	fmt.Println("🚀 line 23 newVersion ➡️", newVersion)
}

// Basic example of how to list tags.
func getLatestTag() string {
	// 打开本地仓库
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Println("Failed to open repository:", err)
		os.Exit(1)
	}

	// 获取标签列表
	tagRefs, err := repo.Tags()
	if err != nil {
		fmt.Println("Failed to get tags:", err)
		os.Exit(1)
	}

	// 遍历标签列表，获取标签名称
	var tags []string
	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		tags = append(tags, tagRef.Name().Short())
		return nil
	})
	if err != nil {
		fmt.Println("Failed to iterate over tags:", err)
		os.Exit(1)
	}

	// 将标签名称按照语义化版本进行排序
	sort.Strings(tags)

	// 获取最新的标签版本号
	latestTag := tags[len(tags)-1]

	fmt.Println("Latest tag version:", latestTag)

	return latestTag
}

func incrementVersion(currentVersion string) string {
	// 解析当前版本号
	re := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)
	matches := re.FindStringSubmatch(currentVersion)
	if len(matches) != 4 {
		fmt.Println("Invalid version format:", currentVersion)
		os.Exit(1)
	}

	// 将版本号的每个部分转换为整数
	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])

	// 版本号自增逻辑
	patch++
	if patch > 9 {
		patch = 0
		minor++
		if minor > 9 {
			minor = 0
			major++
			// if major > 99 {
			// 	fmt.Println("Version number out of range")
			// 	os.Exit(1)
			// }
		}
	}

	// 重新构建版本号字符串
	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	return newVersion
}

func tagExists(tag string, r *git.Repository) bool {
	tagFoundErr := "tag was found"
	tags, err := r.TagObjects()
	if err != nil {
		log.Printf("get tags error: %s", err)
		return false
	}
	res := false
	err = tags.ForEach(func(t *object.Tag) error {
		if t.Name == tag {
			res = true
			return fmt.Errorf(tagFoundErr)
		}
		return nil
	})
	if err != nil && err.Error() != tagFoundErr {
		log.Printf("iterate tags error: %s", err)
		return false
	}
	return res
}

func setTag(r *git.Repository, tag string) (bool, error) {
	if tagExists(tag, r) {
		log.Printf("tag %s already exists", tag)
		return false, nil
	}
	log.Printf("Set tag %s", tag)
	h, err := r.Head()
	if err != nil {
		log.Printf("get HEAD error: %s", err)
		return false, err
	}
	_, err = r.CreateTag(tag, h.Hash(), &git.CreateTagOptions{
		Message: tag,
	})

	if err != nil {
		log.Printf("create tag error: %s", err)
		return false, err
	}

	return true, nil
}

func publicKey(filePath string) (*ssh.PublicKeys, error) {
	var publicKey *ssh.PublicKeys
	sshKey, _ := os.ReadFile(filePath)
	publicKey, err := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		return nil, err
	}
	return publicKey, err
}

func pushTags(r *git.Repository, publicKeyPath string) error {

	auth, _ := publicKey(publicKeyPath)

	po := &git.PushOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
		RefSpecs:   []config.RefSpec{config.RefSpec("refs/tags/*:refs/tags/*")},
		Auth:       auth,
	}
	err := r.Push(po)

	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			log.Print("origin remote was up to date, no push done")
			return nil
		}
		log.Printf("push to remote origin error: %s", err)
		return err
	}

	return nil
}
