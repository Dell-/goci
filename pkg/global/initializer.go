package global

import (
	"github.com/Dell-/goci/models"
	"github.com/Dell-/goci/pkg/settings"
	log "github.com/go-clog/clog"
	"path"
	"io/ioutil"
	"github.com/mcuadros/go-version"
)

func Initialize() {
	settings.NewContext()
	checkVersion()
	err := models.NewEngine();
	if err != nil {
		log.Fatal(2, "DB fail: %v", err)
	}
}

// checkVersion checks if binary matches the version of templates files.
func checkVersion() {
	// Client
	fileVer := path.Join(settings.SERVER.StaticRootPath, "/.VERSION")
	data, err := ioutil.ReadFile(fileVer)
	if err != nil {
		log.Fatal(2, "Fail to read '%s': %v", fileVer, err)
	}
	tplVer := string(data)
	if tplVer != settings.APP.AppVer {
		if version.Compare(tplVer, settings.APP.AppVer, ">") {
			log.Fatal(
				2,
				"[%s] Binary version is lower than client file version, did you forget to recompile Gogs?",
				fileVer,
			)
		} else {
			log.Fatal(
				2,
				"[%s] Binary version is higher than client file version, did you forget to update client files?",
				fileVer,
			)
		}
	}
}
