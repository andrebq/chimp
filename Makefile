include $(GOROOT)/src/Make.inc

TARG=gochimp
GOFILES=\
	chimp/compile.go\
	chimp/file.go\
	chimp/log.go\
	chimp/task.go\
	chimp/filesystem.go\

include $(GOROOT)/src/Make.pkg
