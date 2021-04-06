.PHONY: dockerupdate
dockerupdate:
	wget -O ./bin/gh.tar.gz \
		https://github.com/cli/cli/releases/download/v1.8.1/gh_1.8.1_linux_amd64.tar.gz
	cd ./bin && tar -xzf gh.tar.gz && mv gh_1.8.1_linux_amd64/bin/gh . && \
		rm -rf gh.tar.gz && rm -rf gh_1.8.1_linux_amd64 && cd ..
	docker build . -t dhontecillas/teamworkapigenpr:v0.1
	rm ./bin/gh

testdockerclone:
	docker run --rm --name tdc \
		-v ${PWD}/priv_hnt/data:/data \
		-e GITHUB_TOKEN="$$GTKN" \
		-e GITHUB_USER_EMAIL="$$GITHUB_USER_EMAIL" \
		-e GITHUB_USER_NAME="$$GITHUB_USER_NAME" \
		-e GITHUB_ID_RSA="$$GITHUB_ID_RSA" \
		-e V3SWAGGER="/data/swagger.v3.yaml" \
		dhontecillas/teamworkapigenpr:v0.1
		# gh repo clone dhontecillas/teamworkapigen /data/test_clone

interactive:
	docker run --rm --name tdci -ti \
		-v ${PWD}/priv_hnt/data:/data \
		-e GITHUB_TOKEN=${GTKN} \
		dhontecillas/teamworkapigenpr:v0.1 \
		/bin/bash
