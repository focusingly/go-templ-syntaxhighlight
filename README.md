# go-templ-syntaxhighlight

## Features

1. **html, js/ts, lua, shell, sql, yaml, xml, python, css** grammar highlight in golang raw string with `/* [lang] */` comment prefix;
   ![alt text](.assets/embed-highlight.png)
2. Test data quick insert by `Ctrl+.` for golang file, support insert timestamp, unique id, current time formatted-string(use dayjs)
   ![alt text](.assets/data-insert.png)
3. Support inline basic keyword completion in `/* sql */` commented raw string(it not yet finish)
   ![alt text](.assets/sql-completion.png)
> Tip: you can change unique id generate method and change time format style in setting.
![alt text](.assets/setting.png)

## Project Desciption
```plain
extension-host                extension entry
extension-server-works        provide inline-sql completions(its use lsp protocol)
```

## Note
- this extension not work finished, it maybe cause some problems
