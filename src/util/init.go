package util

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/fatih/color"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"unicode"
)

// GetOs 获取操作系统类型和CPU架构
func GetOs() (string, string) {
	// 获取当前操作系统类型
	osType := runtime.GOOS
	// 获取当前CPU架构
	archType := runtime.GOARCH
	return osType, archType
}

// ExecShell 执行shell命令
func ExecShell(command string) (string, int) {
	osType, _ := GetOs()
	var cmd *exec.Cmd
	switch osType {
	case "windows":
		cmd = exec.Command("cmd", "/c", command)
		break
	default:
		cmd = exec.Command("bash", "-c", command)
		break
	}
	// 创建读写器来实时打印命令的标准输出
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Sprintf("failed to create stdout pipe: %v", err), 106
	}
	stdoutReader := bufio.NewReader(stdoutPipe)

	// 创建读写器来实时打印命令的错误输出
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Sprintf("failed to create stderr pipe: %v", err), 107
	}
	stderrReader := bufio.NewReader(stderrPipe)

	// 启动命令
	if err := cmd.Start(); err != nil {
		return fmt.Sprintf("failed to start command: %v", err), 108
	}

	var outStr bytes.Buffer
	//var errStr bytes.Buffer
	// 在两个独立的goroutine中读取并打印标准输出和错误输出
	go func() {
		for {
			line, _, err := stdoutReader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				color.Red("Error reading from stdout: %v\n", err)
				break
			}
			color.White("[STDOUT] %s\n", line)
			outStr.Write(line)
		}
	}()

	go func() {
		for {
			line, _, err := stderrReader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				color.Red("Error reading from stderr: %v\n", err)
				break
			}
			fmt.Printf("[STDERR] %s\n", line)
			//errStr.Write(line)
		}
	}()

	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		return fmt.Sprintf("command finished with errors: %v", err), 109
	}

	outputStr := outStr.String()
	if osType == "windows" {
		// 从标准输出读取GBK编码的字节流，并转为UTF-8字符串
		decoder := mahonia.NewDecoder("gbk")
		return decoder.ConvertString(outputStr), 0
	}
	return outputStr, 0
}

// IsLinux 判断是否为Linux系统
func IsLinux() bool {
	osType, _ := GetOs()
	//if osType == "linux" && GetDistributionName() == "CentOS" {
	if osType == "linux" || osType == "darwin" {
		return true
	} else {
		return false
	}
}

// GetDistributionName 通过/etc/issue文件获取Linux发行版信息
func GetDistributionName() string {
	// Linux发行版信息存放路径
	filePath := "/etc/issue"
	contentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		color.Red("failed to read distribution name from file\n", err)
		os.Exit(201)
	}
	distributionName := strings.TrimSpace(string(contentBytes))
	return distributionName
}

// IsInstallMaven 判断是否安装Maven
func IsInstallMaven() bool {
	// 检查常见的Maven环境变量路径
	m2Home := os.Getenv("M2_HOME")
	if m2Home != "" {
		_, err := exec.LookPath(filepath.Join(m2Home, "bin", "mvn"))
		return err == nil
	}

	mavenHome := os.Getenv("MAVEN_HOME")
	if mavenHome != "" {
		_, err := exec.LookPath(filepath.Join(mavenHome, "bin", "mvn"))
		return err == nil
	}

	// 如果上述环境变量都未设置，则尝试直接查找系统的mvn命令
	_, err := exec.LookPath("mvn")
	return err == nil
}

// IsInstallJdk 判断是否安装JDK
func IsInstallJdk() bool {
	javaHome := os.Getenv("JAVA_HOME")
	if javaHome != "" {
		_, err := exec.LookPath(filepath.Join(javaHome, "bin", "java"))
		return err == nil
	}
	_, err := exec.LookPath("java")
	return err == nil
}

// IsInstallNode 判断是否安装Node
func IsInstallNode() bool {
	_, err := exec.LookPath("npm")
	if err != nil {
		_, err := exec.LookPath("node")
		return err == nil
	}
	return err == nil
}

// IsInstallGit 判断是否安装Git
func IsInstallGit() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

// IsInstallDocker 判断是否安装Docker
func IsInstallDocker() bool {
	_, err := exec.LookPath("docker")
	return err == nil
}

// IsInstallK8s 判断是否安装K8s
func IsInstallK8s() bool {
	_, err := exec.LookPath("kubectl")
	return err == nil
}

// GetPWD 获取当前工作目录
func GetPWD() string {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		color.Red("failed to get current directory, %v\n", err)
		os.Exit(202)
		return ""
	}
	return dir
}

// CheckFileExists 检查文件是否存在
func CheckFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// WriteContentToFile 写入文件
func WriteContentToFile(filepath string, content string) {
	// 创建或打开一个文件，如果不存在则创建
	file, err := os.Create(filepath)
	if err != nil {
		color.Red("failed to create file, %v\n", err)
	}
	defer file.Close()
	_, err = io.WriteString(file, content)
	if err != nil {
		color.Red("failed to write content, %v\n", err)
	}
}

// GetAppName 获取应用名称
func GetAppName(output string) string {
	splits := strings.Split(output, string(os.PathSeparator))
	split := splits[len(splits)-1]
	return strings.Split(split, ".")[0]
}

// GetYmsHmsStr 获取时间戳
func GetYmsHmsStr() string {
	// 获取当前时间
	now := time.Now()
	// 设置格式模板
	layout := "20060102150405"
	// 格式化时间为字符串
	return now.Format(layout)
}

// IsBlank 检查字符串s是否只包含空白字符或为空
func IsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
