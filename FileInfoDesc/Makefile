export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)


install.rootfs = /var/lib/ethos/ethos-default-$(TARGET_ARCH)/rootfs
install.minimaltd.rootfs = /var/lib/ethos/minimaltd/rootfs


.PHONY: all install
all: DirInfo

install: DirInfo
	ethosTypeInstall $(install.rootfs) $(install.minimaltd.rootfs) DirR
	install DirInfo $(install.rootfs)/programs
	install DirInfo $(install.minimaltd.rootfs)/programs
	echo -n /programs/DirInfo | ethosStringEncode > $(install.rootfs)/etc/init/console
	mkdir -p $(install.rootfs)/user/nobody/myDir
	setfattr -n user.ethos.typeHash -v $(shell egPrint DirR hash DirR) $(install.rootfs)/user/nobody/myDir
	

DirR.go: DirR.t
	$(ETN2GO) . DirR main $^

DirInfo: DirInfo.go DirR.go
	ethosGo $^ 

clean:
	rm -rf DirR/ DirRIndex/
	rm -f DirR.go
	rm -f DirInfo
	rm -f DirInfo.goo.ethos
