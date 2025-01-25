# IPC
App and UI need to communicate with one another. Do this via the IPC-Bus. Don't want to reveal 
too much to the UI, otherwise calls can be made to file system from the front-end. Use
preload.js script -- object that wraps functionality we require in our Electron app.
