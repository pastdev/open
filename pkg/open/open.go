package open

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

func isWSL() bool {
	_, defined := os.LookupEnv("WSL_DISTRO_NAME")
	return defined
}

func wslPath(path string) (string, error) {
	var outBuff bytes.Buffer
	var errBuff bytes.Buffer
	cmd := exec.Command("wslpath", "-w", path)
	cmd.Stderr = &errBuff
	cmd.Stdout = &outBuff
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("wslpath run: %w\n--- stderr ---\n%s\n--- end stderr ---", err, &errBuff)
	}
	return outBuff.String(), nil
}

func wslUrl(u string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return u, fmt.Errorf("wslurl parse: %w", err)
	}
	if parsed.Scheme == "" || parsed.Scheme == "file" {
		return wslPath(parsed.Path)
	}
	return u, nil
}

func Open(url string) (*os.Process, error) {
	switch runtime.GOOS {
	case "darwin":
		return OpenDarwin(url)
	case "linux":
		return OpenLinux(url)
	case "windows":
		return OpenWindows(url)
	}
	return nil, fmt.Errorf("unsupported os: %s", runtime.GOOS)
}

func OpenDarwin(url string) (*os.Process, error) {
	cmd := exec.Command("open", url)
	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("opendarwin exec open: %w", err)
	}
	return cmd.Process, nil
}

func OpenLinux(url string) (*os.Process, error) {
	if isWSL() {
		url, err := wslUrl(url)
		if err != nil {
			return nil, fmt.Errorf("openlinux wslurl: %w", err)
		}
		return OpenWindows(url)
	}

	cmd := exec.Command("xdg-open", url)
	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("openlinux exec open: %w", err)
	}
	return cmd.Process, nil
}

func OpenWindows(url string) (*os.Process, error) {
	cmd := exec.Command("powershell.exe", "-c", fmt.Sprintf("Start-Process -FilePath %s", url))
	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("openonwindows exec powershell start-process: %w", err)
	}
	return cmd.Process, nil
}
