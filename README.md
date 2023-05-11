# ingiosapis

This is a simple Go project that uses build.buf to work with Google Protocol Buffer files.  
The project produces a Go package containing the generated code, which includes 
marshaling code produced by two different packages (comparing features):

* github.com/mitchellh/protoc-gen-go-json
* github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto

and validation code produced by github.com/envoyproxy/protoc-gen-validate, which requires that messages exist in the
proto files, otherwise it has nothing to generate.

If all required plug-ins support remote packages, then local generation is not necessary.
You can delete the buf.gen.yaml file, 'buf push' the proto files to BSR, and 'go get' the package from BSR
for use in a Go application.  Simply add the package to the go.mod file and import the package where it is used.

