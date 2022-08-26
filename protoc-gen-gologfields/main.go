// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/nvx/go-proto-logfields/internal/genlogfields"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if f.Generate {
				genlogfields.GenerateFile(gen, f)
			}
		}
		return nil
	})
}
