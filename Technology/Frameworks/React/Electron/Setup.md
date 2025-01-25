# Setup
## Model
Entire chromium browser inside of app, where the core process is a NodeJS app. First tasked with creating
windows, and has the capability of interacting with System APIs. Also furnished with IPC Event bus,
allowing main process and windoes to publish events and data.
## Vite
```js
export default defineConfig({
  plugins: [react()]
  build: {
    outDir: 'dist-react'
  },
});
```

## src.electron.main
Plate of Boiler for you.
```js
import { app, BroswerWindow } from 'electron'

app.on("ready", () => {
  const mainWindow = new BrowserWindow()
if (isDev()) {
  mainWindow.loadURL('https://localhost:5123');
} else {
  mainWindow.loadFile(app.getAppPath() + '/dist-react/index.html');
}
})
```
## package.json
Let's tell Electron where to find the script.
```js
"name": "electron-course"
"private": true
"version" : "0.0.0"
"type": "module"
"main": "src/electron/main"

```

## Then lets change TS Config
```js
compilerOptions: {
  strict: true
  target: ESNext,
  module: NodeNext,
  outDir: "../../dist-electron",
  skipLibCheck: true
}
```
## Misc.
At this point, change .gitignore, and make sure you're properly running the compiled file from the dist.
### Create Dev Check
```js
export function isDev() {
  retun process.env.NODE_ENV == 'development';
}
```
Use Vite Hot Module Reloading Server in development, but a tree-build application. Remember to set fixed port in Vite Config file.
## electron-builder.json
```js
{
  appId: "com.prepare"
  files: ["dist-electron", "dist-react"],
  "mac": {
    target: "dmg"
  }
  linux: {
    target: "AppImage",
    category: "Utility",
},
  win: { // WINDOWS REQUIRE an ICON field in this file.
    target: [ "portable", "msi" ]
  }
}
```
