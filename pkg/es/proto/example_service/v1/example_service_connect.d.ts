// @generated by protoc-gen-connect-es v1.4.0
// @generated from file example_service/v1/example_service.proto (package example_service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { TypeRequest, TypeResponse } from "./example_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service example_service.v1.ExampleService
 */
export declare const ExampleService: {
  readonly typeName: "example_service.v1.ExampleService",
  readonly methods: {
    /**
     * @generated from rpc example_service.v1.ExampleService.Type
     */
    readonly type: {
      readonly name: "Type",
      readonly I: typeof TypeRequest,
      readonly O: typeof TypeResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

