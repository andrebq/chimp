include $(GOROOT)/src/Make.inc

TARG=chimp
GOFILES=\
	chimp/compile.go\
	chimp/file.go\
	chimp/log.go\
	chimp/task.go\
	chimp/filesystem.go\
	chimp/goenv.go\

include $(GOROOT)/src/Make.pkg
