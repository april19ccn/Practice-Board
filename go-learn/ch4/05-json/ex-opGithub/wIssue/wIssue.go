package wIssue

import (
	"os"
	"os/exec"
)

// 获取系统中的编辑器
func getEditor() string {
	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor
	}
	if editor := os.Getenv("VISUAL"); editor != "" {
		return editor
	}
	return "vim" // 默认使用 vim
}

// 创建临时文件
func createTempFile(content string) (string, error) {
	tmpfile, err := os.CreateTemp("", "issue_*.txt")
	if err != nil {
		return "", err
	}
	defer tmpfile.Close()

	if _, err := tmpfile.WriteString(content); err != nil {
		os.Remove(tmpfile.Name())
		return "", err
	}
	return tmpfile.Name(), nil
}

// 用编辑器打开临时文件
func openEditor(filename string) error {
	editor := getEditor()
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// 获取编辑后的内容，并清除临时文件
func getEditedContent(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	os.Remove(filename) // 清理临时文件
	return string(content), err
}

// 获取描述
func GetIssueDescription() (string, error) {
	tmpfile, _ := createTempFile("# 在此输入 Issue 描述\n")
	openEditor(tmpfile)
	description, err := getEditedContent(tmpfile)
	if err != nil {
		return "", err
	}

	return description, nil
}

// func main() {
// 	tmpfile, _ := createTempFile("# 在此输入 Issue 描述\n")
// 	openEditor(tmpfile)
// 	description, _ := getEditedContent(tmpfile)

// 	fmt.Println(description)
// }
