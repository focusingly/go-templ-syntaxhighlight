import type { ExtensionContext } from "vscode";

export {};

declare global {
  var extensionCtx: ExtensionContext;
}
