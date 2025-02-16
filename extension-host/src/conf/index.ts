import net from "net";
import os from "os";
import path from "path";
import vscode, { workspace } from "vscode";
import {
  RevealOutputChannelOn,
  TransportKind,
  type LanguageClientOptions,
  type ServerOptions,
  type StreamInfo
} from "vscode-languageclient/node";

export enum LspServerMode {
  NetDebug = "net-dbg",
  UnixSoc = "unix-sock",
  StdProd = "std-prod"
}

export const unixSockDomainDbgPath = path.join(os.tmpdir(), "lsp-server-dbg.sock");

export const getMode = (): LspServerMode => {
  const prefix = "-mode=";
  let mode = process.argv.find((arg) => arg.startsWith("-mode"));

  if (!mode) {
    const f = process.env["PROD_MODE"];
    if (f) {
      mode = `${prefix}${f}`;
    } else {
      mode = `${prefix}${LspServerMode.StdProd}`;
    }
  }

  const op = mode.substring(prefix.length) as LspServerMode | void;
  switch (op) {
    case LspServerMode.NetDebug:
    case LspServerMode.StdProd:
      return op;
    default:
      return LspServerMode.StdProd;
  }
};

export function loadLspServerLoadOptions(context: vscode.ExtensionContext): ServerOptions {
  switch (getMode()) {
    case LspServerMode.NetDebug:
      const netConn = net.createConnection(8443, "localhost", () => {
        console.info("Connect Local Lsp Debug Server By Remote Net Successfully");
      });
      return async () =>
        ({
          reader: netConn,
          writer: netConn,
          detached: true
        }) as StreamInfo;
    case LspServerMode.UnixSoc:
      const unixSockDomainConn = net.createConnection(unixSockDomainDbgPath, () => {
        console.info("Connect Local Lsp Debug Server By Unix Socket Domain Successfully");
      });
      return async () =>
        ({
          reader: unixSockDomainConn,
          writer: unixSockDomainConn,
          detached: true
        }) as StreamInfo;
    case LspServerMode.StdProd:
      const serverPath = context.asAbsolutePath(path.join("dist", "server", "lsp-server.exe"));
      return {
        run: {
          command: serverPath,
          transport: TransportKind.stdio
        },
        debug: {
          command: serverPath,
          transport: TransportKind.stdio
        }
      } as ServerOptions;
  }
}

export const loadClientOptions = (): LanguageClientOptions => {
  const t = {} as { workspaceFolder: vscode.WorkspaceFolder };
  if (workspace.workspaceFolders?.length) {
    t.workspaceFolder = workspace.workspaceFolders[0];
  }

  return {
    documentSelector: [{ scheme: "file", language: "go" }],
    markdown: {
      isTrusted: true,
      supportHtml: true
    },
    revealOutputChannelOn: RevealOutputChannelOn.Info,
    ...t
  };
};
