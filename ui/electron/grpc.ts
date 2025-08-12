import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import path from "node:path";

const pkgDef = protoLoader.loadSync(path.join(process.cwd(), "../core/api/control.proto"));
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const proto = (grpc.loadPackageDefinition(pkgDef) as any).vadim.control.v1;

export function makeClient() {
  // На Windows временно используем TCP 127.0.0.1:50055 (см. server.go)
  const address = process.platform === "win32" ? "127.0.0.1:50055" : "unix:/tmp/vadim-decktop.sock";
  return new proto.VadimControl(address, grpc.credentials.createInsecure());
}
