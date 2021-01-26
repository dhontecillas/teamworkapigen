

Go Config options:


- **packageName**: Go package name (convention: lowercase). (Default: openapi)
- **packageVersion**: Go package version. (Default: 1.0.0)
- **hideGenerationTimestamp**: Hides the generation timestamp when files are generated. (Default: true)
- **isGoSubmodule**: whether the generated Go module is a submodule (Default: false)
- **withGoCodegenComment**: whether to include Go codegen comment to disable Go Lint and collapse by default GitHub in PRs and diffs (Default: false)
- **withXml**: whether to include support for application/xml content type and include XML annotations in the model (works with libraries that provide support for JSON and XML) (Default: false)
- **enumClassPrefix**: Prefix enum with class name (Default: false)
- **structPrefix**: whether to prefix struct with the class name. e.g. DeletePetOpts => PetApiDeletePetOpts (Default: false)
- **withAWSV4Signature**: whether to include AWS v4 signature support (Default: false)
- **prependFormOrBodyParameters**: Add form or body parameters to the beginning of the parameter list. (Default: false)



Python Config options:
- **packageName**: python package name (convention: snake_case). (Default: openapi_client)
- **projectName**: python project name in setup.py (e.g. petstore-api).
- **packageVersion**: python package version. (Default: 1.0.0)
- **packageUrl**: python package URL.
- **sortParamsByRequiredFlag**: Sort method arguments to place required parameters before optional parameters. (Default: true)
- **hideGenerationTimestamp**: Hides the generation timestamp when files are generated. (Default: true)
- **generateSourceCodeOnly**: Specifies that only a library source code is to be generated. (Default: false)
- **useNose**: use the nose test framework (Default: false)
- **library**: library template (sub-template) to use: asyncio, tornado, urllib3 (Default: urllib3)
