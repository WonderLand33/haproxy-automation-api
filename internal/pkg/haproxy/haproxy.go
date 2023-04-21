package haproxy

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/scylladb/go-set"
	"github.com/scylladb/go-set/strset"
)

var (
	blockedIPFilePath = "/etc/haproxy/blocked.ips"
)

// 注意：本地若没有安装haproxy的话，请注释该方法。
func Reload() error {
	cmd := exec.Command("systemctl", "reload", "haproxy")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("reload haproxy 出错: %w", err)
	}
	return nil
}

func GetBlockedIPs() (*strset.Set, error) {
	fp, err := os.Open(blockedIPFilePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件出错: %w", err)
	}

	defer fp.Close()

	ips := set.NewStringSet()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			ips.Add(text)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("读取文件出错: %w", err)
	}

	return ips, nil
}

func Update(ips *strset.Set) error {
	fp, err := os.OpenFile(blockedIPFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开文件出错: %w", err)
	}

	defer fp.Close()

	content := strings.Join(ips.List(), "\n")
	_, err = fp.WriteString(content)
	if err != nil {
		return fmt.Errorf("写入文件出错: %w", err)
	}

	err = Reload()
	if err != nil {
		return err
	}

	return nil
}
