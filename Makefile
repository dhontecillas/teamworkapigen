.PHONY: dockerupdate
builddockerupdater:
	wget -O ./bin/gh.tar.gz \
		https://github.com/cli/cli/releases/download/v1.8.1/gh_1.8.1_linux_amd64.tar.gz
	cd ./bin && tar -xzf gh.tar.gz && mv gh_1.8.1_linux_amd64/bin/gh . && \
		rm -rf gh.tar.gz && rm -rf gh_1.8.1_linux_amd64 && cd ..
	docker build . -f Dockerfile.create_update_pr -t dhontecillas/teamworkapigenpr:v0.1
	rm ./bin/gh

.PHONY: runupdater
runupdater:
	docker run --rm --name tdc \
		-v ${PWD}/priv_hnt/data:/data \
		-e GITHUB_TOKEN="$$GITHUB_TOKEN" \
		-e GITHUB_USER_EMAIL="$$GITHUB_USER_EMAIL" \
		-e GITHUB_USER_NAME="$$GITHUB_USER_NAME" \
		-e V1SWAGGER="$$V1SWAGGER" \
		-e V2SWAGGER="$$V2SWAGGER" \
		-e V3SWAGGER="$$V3SWAGGER" \
		-e PMSWAGGER="$$PMSWAGGER" \
		dhontecillas/teamworkapigenpr:v0.1
