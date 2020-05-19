package main

import (
	"encoding/json"
	"fmt"
	"io"

	docker "github.com/fsouza/go-dockerclient"

	"github.com/chmouel/daemonless/pkg/builder"
	"github.com/containers/image/v5/types"

	istorage "github.com/containers/image/v5/storage"

	buildclientv1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"

	log "github.com/golang/glog"

	"os"

	"github.com/containers/storage"
)

// DockerClient is an interface to the Docker client that contains
// the methods used by the common builder
type DockerClient interface {
	BuildImage(opts docker.BuildImageOptions) error
	PushImage(opts docker.PushImageOptions, auth docker.AuthConfiguration) (string, error)
	RemoveImage(name string) error
	CreateContainer(opts docker.CreateContainerOptions) (*docker.Container, error)
	PullImage(opts docker.PullImageOptions, authSearchPaths []string) error
	RemoveContainer(opts docker.RemoveContainerOptions) error
	InspectImage(name string) (*docker.Image, error)
	TagImage(name string, opts docker.TagImageOptions) error
}

type builderConfig struct {
	out             io.Writer
	sourceSecretDir string
	dockerClient    DockerClient
	dockerEndpoint  string
	buildsClient    buildclientv1.BuildInterface
	cleanup         func()
	store           storage.Store
	blobCache       string
}

func main() {
	cfg := &builderConfig{}

	var systemContext types.SystemContext
	if registriesConfPath, ok := os.LookupEnv("BUILD_REGISTRIES_CONF_PATH"); ok && len(registriesConfPath) > 0 {
		if _, err := os.Stat(registriesConfPath); err == nil {
			systemContext.SystemRegistriesConfPath = registriesConfPath
		}
	}
	if registriesDirPath, ok := os.LookupEnv("BUILD_REGISTRIES_DIR_PATH"); ok && len(registriesDirPath) > 0 {
		if _, err := os.Stat(registriesDirPath); err == nil {
			systemContext.RegistriesDirPath = registriesDirPath
		}
	}
	if signaturePolicyPath, ok := os.LookupEnv("BUILD_SIGNATURE_POLICY_PATH"); ok && len(signaturePolicyPath) > 0 {
		if _, err := os.Stat(signaturePolicyPath); err == nil {
			systemContext.SignaturePolicyPath = signaturePolicyPath
		}
	}

	storeOptions, err := storage.DefaultStoreOptions(false, 0)
	if err != nil {
		return
	}
	if driver, ok := os.LookupEnv("BUILD_STORAGE_DRIVER"); ok {
		storeOptions.GraphDriverName = driver
	}
	if storageOptions, ok := os.LookupEnv("BUILD_STORAGE_OPTIONS"); ok {
		if err := json.Unmarshal([]byte(storageOptions), &storeOptions.GraphDriverOptions); err != nil {
			log.V(0).Infof("Error parsing BUILD_STORAGE_OPTIONS (%q): %v", storageOptions, err)
			return
		}
	}
	if storageConfPath, ok := os.LookupEnv("BUILD_STORAGE_CONF_PATH"); ok && len(storageConfPath) > 0 {
		if _, err := os.Stat(storageConfPath); err == nil {
			storage.ReloadConfigurationFile(storageConfPath, &storeOptions)
		}
	}

	store, err := storage.GetStore(storeOptions)
	cfg.store = store
	if err != nil {
		return
	}
	cfg.cleanup = func() {
		if _, err := store.Shutdown(false); err != nil {
			log.V(0).Infof("Error shutting down storage: %v", err)
		}
	}
	istorage.Transport.SetStore(store)

	// Default to using /var/cache/blobs as a blob cache, but allow its location
	// to be changed by setting $BUILD_BLOBCACHE_DIR.  Setting the location to an
	// empty value disables the cache.
	cfg.blobCache = "/var/cache/blobs"
	if blobCacheDir, isSet := os.LookupEnv("BUILD_BLOBCACHE_DIR"); isSet {
		cfg.blobCache = blobCacheDir
	}

	dockerClient, err := builder.GetDaemonlessClient(systemContext, store, os.Getenv("BUILD_ISOLATION"), "/var/cache/blobs")
	fmt.Println(dockerClient)
}
