import dayjs from "dayjs";
import { nanoid } from "nanoid";
import vscode from "vscode";

const insertTextInCursor = (text: string) => {
  const editor = vscode.window.activeTextEditor;
  if (editor) {
    const position = editor.selection.active;
    editor.edit((editBuilder) => {
      editBuilder.insert(position, text);
    });
  }
};

export enum UniqueIDMark {
  UUID32 = "uuid-32",
  UUID36 = "uuid-36",
  NANOID = "nano-id"
}

export function registerBaseCommands(ctx: vscode.ExtensionContext) {
  const uniqueIDCmd = vscode.commands.registerCommand("go-templ:insert-uniqueID", () => {
    const id = vscode.workspace.getConfiguration().get<UniqueIDMark>("conf.goTempl.uniqueID", UniqueIDMark.UUID36);
    switch (id) {
      case UniqueIDMark.UUID32:
        insertTextInCursor(crypto.randomUUID().replaceAll("-", ""));
        break;
      case UniqueIDMark.UUID36:
        insertTextInCursor(crypto.randomUUID());
        break;
      case UniqueIDMark.NANOID:
        insertTextInCursor(nanoid());
        break;
      default:
        vscode.window.showWarningMessage(`Unknown Unique ID: ${id}, Use Fallback: ${UniqueIDMark.UUID36}`);
        insertTextInCursor(crypto.randomUUID());
    }
  });

  const epochSecCmd = vscode.commands.registerCommand("go-templ:insert-epoch-secs", () => {
    insertTextInCursor(`${(Date.now() / 1000) | 0}`);
  });

  const epochMillCmd = vscode.commands.registerCommand("go-templ:insert-epoch-mills", () => {
    insertTextInCursor(`${Date.now()}`);
  });

  const fmtTimeCmd = vscode.commands.registerCommand("go-templ:insert-fmt-time", () => {
    const conf = vscode.workspace.getConfiguration().get<string>("conf.goTempl.time.format");

    try {
      const str = dayjs().format(conf || "YYYY-MM-DDTHH:mm:ss.sss");
      insertTextInCursor(str);
    } catch (e) {
      vscode.window.showErrorMessage(`An Error Time Config For Go-Templ: ${conf}`);
    }
  });

  ctx.subscriptions.push(uniqueIDCmd, epochSecCmd, epochMillCmd, fmtTimeCmd);
}
