package cmd

const (
	defaultDomain          = "mydomain.adam"
	defaultIP              = "192.168.0.1"
	defaultEVEIP           = "192.168.1.2"
	defaultUUID            = "1"
	defaultAdamTag         = "0.0.26"
	defaultEveTag          = "5.1.11"
	defaultEvePrefixInTar  = "bits"
	defaultEveRepo         = "https://github.com/lf-edge/eve.git"
	defaultLinuxKitVersion = "v0.7"
	defaultImageTag        = "eden-alpine"
	defaultFileToSave      = "./test.tar"
	defaultImage           = "library/alpine"
	defaultRegistry        = "docker.io"
	defaultIsLocal         = false
	defaultQemuFileToSave  = "qemu.conf"
	defaultQemuCpus        = 4
	defaultQemuMemory      = 4096
	defaultEserverPort     = "8888"
	defaultFilename        = "rootfs.img"
	imageID                = "1ab8761b-5f89-4e0b-b757-4b87a9fa93ec"
	dataStoreID            = "eab8761b-5f89-4e0b-b757-4b87a9fa93ec"
	baseID                 = "22b8761b-5f89-4e0b-b757-4b87a9fa93ec"
	defaultTestProg        = "eden.integration.test"
	defaultTestScript      = "eden.integration.test"
	rootFSVersionPattern   = `^(\d+\.*){2,3}.*-(xen|kvm|acrn)-(amd64|arm64)$`
	controllerModePattern  = `^(?P<Type>(file|proto|adam|zedcloud)):\/\/(?P<URL>.+)$`
)

var (
	defaultQemuHostFwd = map[string]string{"2222": "22"}
	cobraToViper       = map[string]string{
		"adam.dist":   "adam-dist",
		"adam.tag":    "adam-tag",
		"adam.port":   "adam-port",
		"adam.domain": "domain",
		"adam.ip":     "ip",
		"adam.eve-ip": "eve-ip",
		"adam.force":  "adam-force",
		"adam.v1":     "api-v1",

		"eve.arch":         "eve-arch",
		"eve.os":           "eve-os",
		"eve.accel":        "eve-accel",
		"eve.hv":           "hv",
		"eve.serial":       "eve-serial",
		"eve.pid":          "eve-pid",
		"eve.log":          "eve-log",
		"eve.firmware":     "eve-firmware",
		"eve.repo":         "eve-repo",
		"eve.tag":          "eve-tag",
		"eve.base-tag":     "eve-base-tag",
		"eve.hostfwd":      "eve-hostfwd",
		"eve.dist":         "eve-dist",
		"eve.base-dist":    "eve-base-dist",
		"eve.qemu-config":  "qemu-config",
		"eve.uuid":         "uuid",
		"eve.image-file":   "image-file",
		"eve.dtb-part":     "dtb-part",
		"eve.config-part":  "config-part",
		"eve.base-version": "os-version",

		"eden.images.dist":   "image-dist",
		"eden.images.docker": "docker-yml",
		"eden.images.vm":     "vm-yml",
		"eden.download":      "download",
		"eden.eserver.ip":    "eserver-ip",
		"eden.eserver.port":  "eserver-port",
		"eden.eserver.pid":   "eserver-pid",
		"eden.eserver.log":   "eserver-log",
		"eden.certs-dist":    "certs-dist",
		"eden.bin-dist":      "bin-dist",
		"eden.ssh-key":       "ssh-key",
		"eden.test-bin":      "eden.integration.test",
		"test-script":        "eden.integration.tests.txt",
	}
)
