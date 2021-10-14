SUBDIRS = $(shell find . -maxdepth 1 -type d|grep build)

docker:
	$(call make_exec, docker)

define make_exec
	@for dir in $(SUBDIRS) ; do \
		$(MAKE) $1 -C $$dir; \
		done;
endef
