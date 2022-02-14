# go-proto-logfields

A protoc plugin for generating log field extractors.

## Use-case

It is useful to augment log messages with additional context from the execution environment. For example, when handling an RPC, it may be useful to add some fields of the request message to any messages logged while handling the request.

## Install

* Install Go.
* Install `protoc` and `protoc-gen-go`, and ensure they are in your path.
* `go install github.com/nvx/go-proto-logfields/protoc-gen-gologfields`

## Usage

Given a an RPC:
```
syntax = "proto3";
package doer_of_something;
import "logfields.proto";

Service Doer {
  rpc DoSomething (SomethingRequest) returns (SomethingResponse) {}
}
```

The request might be annotated as such:
```
message SomethingRequest {
    string what = 1 [(improbable.logfield) = {name: "what_was_requested"}];
}
```

The logfields extractor can be generated with buf.build using `buf generate.
See the `buf.yaml` and `buf.gen.yaml` files in this repository for an example.
