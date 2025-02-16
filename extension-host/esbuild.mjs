import esbuild from "esbuild";
import alias from "esbuild-plugin-alias";

const production = process.argv.includes("--production");
const watch = process.argv.includes("--watch");

const BRIGHT_GREEN = "\x1b[92m";
const BRIGHT_RED = "\x1b[91m";
const RESET = "\x1b[0m";

/**
 * @type {esbuild.Plugin}
 */
const esbuildProblemMatcherPlugin = {
  name: "esbuild-problem-matcher",
  setup(build) {
    build.onStart(() => {
      console.log(`${BRIGHT_GREEN}[watch] build started${RESET}`);
    });
    build.onEnd((result) => {
      result.errors.forEach(({ text, location }) => {
        console.error(`${BRIGHT_RED}✘ [ERROR] ${text}${RESET}`);
        console.error(`${BRIGHT_RED}    ${location?.file}:${location?.line}:${location?.column}:${RESET}`);
      });
      console.log(`${BRIGHT_GREEN}[watch] build finished${RESET}`);
    });
  }
};

async function main() {
  production
    ? console.log(`${BRIGHT_GREEN}RUNNING AT PRODUCTION MODE${RESET}`)
    : console.log(`${BRIGHT_GREEN}RUNNING AT DEBUG MODE${RESET}`);

  const ctx = await esbuild.context({
    entryPoints: ["src/extension.ts"],
    bundle: true,
    format: "cjs",
    minify: production,
    sourcemap: !production,
    sourcesContent: !production,
    platform: "node",
    outfile: "dist/extension.js",
    external: ["vscode"],
    logLevel: "debug",
    plugins: [
      /* add to the end of plugins array */
      esbuildProblemMatcherPlugin,
      alias({
        "@": "src"
      })
    ]
  });
  if (watch) {
    await ctx.watch();
  } else {
    await ctx.rebuild();
    await ctx.dispose();
  }
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
