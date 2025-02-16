import vscode from "vscode";
export function provideTestDataInsertAction(ctx: vscode.ExtensionContext) {
  const codeAction = vscode.languages.registerCodeActionsProvider(["go"], {
    provideCodeActions(_document, _range, context, _token) {
      const actions: vscode.CodeAction[] = [];
      if (context.diagnostics.length === 0) {
        const action1 = new vscode.CodeAction("Insert an Unique ID", vscode.CodeActionKind.QuickFix);
        action1.command = {
          title: "Insert an Unique ID",
          command: "go-templ:insert-uniqueID"
        };

        const action2 = new vscode.CodeAction("Insert Current Epoch Seconds", vscode.CodeActionKind.QuickFix);
        action2.command = {
          title: "Insert Current Epoch timestamp Seconds",
          command: "go-templ:insert-epoch-secs"
        };

        const action3 = new vscode.CodeAction("Insert Current Epoch Milliseconds", vscode.CodeActionKind.QuickFix);
        action3.command = {
          title: "Insert Current Epoch Timestamp Milliseconds",
          command: "go-templ:insert-epoch-mills"
        };

        const action4 = new vscode.CodeAction("Insert Current Time Formatted String", vscode.CodeActionKind.QuickFix);
        action4.command = {
          title: "insert current time formatted string by config",
          command: "go-templ:insert-fmt-time"
        };

        actions.push(action1, action2, action3, action4);
      }

      return actions;
    }
  });

  ctx.subscriptions.push(codeAction);
}
