# https://github.com/go-gitea/gitea
LIB_NAME=gitea
LIB=github.com/go-gitea/$(LIB_NAME)
LIB_BRANCH=master
LIB_TAG=v1.9.5
LIB_FSPATH=$(GOPATH)/src/$(LIB)

GO111MODULE=on

SAMPLE_NAME=
SAMPLE_FSPATH=$(LIB_FSPATH)/$(SAMPLE_NAME)

CLOUD_PROJECT_ID=winwisely-cloudrun-gitea
CLOUD_PROJECT_URL=????

