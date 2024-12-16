package utils

import (
	"fmt"
	"os"

	"github.com/jfrog/jfrog-client-go/utils/log"
	"github.com/pelletier/go-toml"
)

type Package struct {
	Name     string `toml:"name"`
	Version  string `toml:"version"`
	Source   string `toml:"source"`
	Checksum string `toml:"checksum"`
}

type Lock struct {
	Packages []Package `toml:"package"`
}

func GetPackages() ([]Package, error) {
	LockFile := os.Getenv("CARGO_LOCKFILE")
	if LockFile == "" {
		LockFile = DefaultLockFile
	}

	if os.Getenv("CARGO_SKIP") != "true" {
		err := ExecuteCmd("cargo", "--version")
		if err != nil {
			log.Debug(fmt.Sprintf("Failed to run cargo: %s", err))
			return nil, err
		}

		err = ExecuteCmd("cargo", "generate-lockfile", "--frozen")
		if err != nil {
			log.Debug(fmt.Sprintf("Failed to run cargo: %s", err))
			return nil, err
		}
	}

	content, err := os.ReadFile(LockFile)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to read file: %s", err))
		return nil, err
	}

	var cargoLock Lock
	err = toml.Unmarshal(content, &cargoLock)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to parse TOML: %s", err))
		return nil, err
	}
	log.Debug(fmt.Sprintf("Packages: %d", len(cargoLock.Packages)))
	return cargoLock.Packages, nil
}
