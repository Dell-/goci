package settings

import (
	"gopkg.in/ini.v1"
	"os"
	"fmt"
	"os/exec"
	"path/filepath"
	log "github.com/go-clog/clog"
	"strings"
	"github.com/Unknwon/com"
	"path"
)

var (
	// App settings
	APP struct {
		Path     string
		LogPath  string `ini:"LOG_PATH"`
		DataPath string `ini:"DATA_PATH"`
		AppVer string `ini:"APP_VER"`
	}

	// Server settings
	SERVER struct {
		Scheme         string `ini:"SCHEME"`
		Domain         string `ini:"DOMAIN"`
		HTTPAddr       string `ini:"HTTP_ADDR"`
		HTTPPort       int  `ini:"HTTP_PORT"`
		Url            string `ini:"URL"`
		ApiUrl            string `ini:"API_URL"`
		StaticRootPath string `ini:"STATIC_ROOT_PATH"`
	}

	// Db settings
	DB struct {
		Username string `ini:"USERNAME"`
		Password string `ini:"PASSWORD"`
		Host     string `ini:"HOST"`
		Port     string `ini:"PORT"`
		Name     string `ini:"NAME"`
		Charset  string `ini:"CHARSET"`
	}

	// Global setting objects
	Cfg *ini.File
)

func init() {
	var err error
	err = log.New(log.CONSOLE, log.ConsoleConfig{})
	if err != nil {
		fmt.Printf("Fail to create new logger: %v\n", err)
		os.Exit(1)
	}
}

// execPath returns the executable path
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

// WorkDir returns absolute path of work directory.
func WorkDir() (string, error) {
	wd := os.Getenv("GOCI_WORK_DIR")
	if len(wd) > 0 {
		return wd, nil
	}

	execPath, err := execPath()

	if err != nil {
		log.Fatal(2, "Fail to get app path: %v\n", err)
	}

	execPath = strings.Replace(execPath, "\\", "/", -1)

	i := strings.LastIndex(execPath, "/")
	if i == -1 {
		return execPath, nil
	}
	return execPath[:i], nil
}

func NewContext() {
	workDir, err := WorkDir()

	if err != nil {
		log.Fatal(2, "Fail to get work directory: %v\n", err)
	}

	cfgFile := path.Join(workDir, "conf", "app.ini")

	if com.IsFile(cfgFile) {
		Cfg, err = ini.Load(cfgFile)
		if err != nil {
			log.Fatal(2, "Fail to load conf '%s': %v", cfgFile, err)
		}
	} else {
		log.Fatal(2, "Goci config file (%s) does not exist: %v\n", cfgFile, err)
	}

	APP.AppVer = Cfg.Section("APP").Key("APP_VER").String()
	APP.Path = workDir
	APP.LogPath = path.Join(workDir, Cfg.Section("APP").Key("LOG_PATH").MustString("log"))

	if !com.IsExist(APP.LogPath) {
		log.Fatal(2, "Log directory (%s) does not exist: %v\n", APP.LogPath, err)
	}

	APP.DataPath = path.Join(workDir, Cfg.Section("APP").Key("DATA_PATH").MustString("data"))

	if !com.IsExist(APP.DataPath) {
		log.Fatal(2, "Data directory (%s) does not exist: %v\n", APP.DataPath, err)
	}

	err = Cfg.Section("DB").MapTo(&DB)

	if err != nil {
		log.Fatal(2, "DB section not found: %v\n", err)
	}

	err = Cfg.Section("SERVER").MapTo(&SERVER)

	if err != nil {
		log.Fatal(2, "SERVER section not found: %v\n", err)
	}

	SERVER.StaticRootPath = path.Join(workDir, SERVER.StaticRootPath)
}
