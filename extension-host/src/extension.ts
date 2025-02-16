import vscode from "vscode";
import { LanguageClient } from "vscode-languageclient/node";
import { provideTestDataInsertAction } from "./actions";
import { registerBaseCommands } from "./commands";
import { loadClientOptions, loadLspServerLoadOptions } from "./conf";

let client: LanguageClient | void = void 0;

export function activate(context: vscode.ExtensionContext) {
  global.extensionCtx = context;

  registerBaseCommands(context);
  provideTestDataInsertAction(context);

  client = new LanguageClient("go-templ", "Go-Templ-Server", loadLspServerLoadOptions(context), loadClientOptions());
  client.start();
  context.subscriptions.push(client);
}

export function deactivate() {
  if (typeof client === "undefined") {
    return void 0;
  }

  return client.stop();
}
