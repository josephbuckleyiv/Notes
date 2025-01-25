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
  mainWindow.loadFile(app.getAppPath() + '/dist-react/index.html');
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
