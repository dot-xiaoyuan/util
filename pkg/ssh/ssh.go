package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
)

type Terminal struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Session  *ssh.Session
}

func (t *Terminal) New() error {
	// SSH
	sshConfig := &ssh.ClientConfig{
		User: t.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(t.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to Server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", t.Host, t.Port), sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
		return err
	}
	defer client.Close()

	// Create New Session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
		return err
	}
	defer session.Close()

	// Get current fd 获取当前终端的文件描述符
	fd := int(os.Stdout.Fd())

	// Get old term 获取终端的原始模式
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		log.Fatalf("failed to make raw ssh: %v", err)
	}
	defer terminal.Restore(fd, oldState) // 恢复终端模式

	// 设置会话标准输入、输出和错误
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// 获取当前终端尺寸
	width, height, err := terminal.GetSize(fd)
	if err != nil {
		log.Fatalf("failed to get ssh size: %v", err)
		return err
	}

	// 请求伪终端分配
	err = session.RequestPty("xterm-256color", height, width, ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_OSPEED: 14400,
		ssh.TTY_OP_ISPEED: 14400,
	})
	if err != nil {
		log.Fatalf("failed to request ssh: %v", err)
		return err
	}
	if err = session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %v", err)
		return err
	}
	if err = session.Wait(); err != nil {
		log.Fatalf("failed to wait: %v", err)
		return err
	}
	return nil
}

func (t *Terminal) Run(command string) {
	err := t.Session.Run(command)
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
