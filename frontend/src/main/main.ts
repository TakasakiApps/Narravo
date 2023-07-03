import { app, BrowserWindow, ipcMain, session, globalShortcut, Menu, nativeTheme } from 'electron';
import { join } from 'path';





function createWindow() {

  const mainWindow = new BrowserWindow({
    width: 800,
    height: 700,
    minHeight:600,
    webPreferences: {
      preload: join(__dirname, 'preload.js'),
      nodeIntegration: true,
      contextIsolation: true,
    }
  });

  if (process.env.NODE_ENV === 'development') {
    const rendererPort = process.argv[2];
    mainWindow.loadURL(`http://localhost:${rendererPort}`);
  }
  else {
    mainWindow.loadFile(join(app.getAppPath(), 'renderer', 'index.html'));
  }

  

  app.on('ready', function ()  {
    globalShortcut.register('CommandOrControl+Shift+i', function () {
      mainWindow.webContents.openDevTools()
    })
    

    
createWindow();

  })
// Menu.setApplicationMenu(null); 
//调试完成删掉上面的注释标识符
}

app.whenReady().then(() => {
  createWindow();
  session.defaultSession.webRequest.onHeadersReceived((details, callback) => {
    callback({
      responseHeaders: {
        ...details.responseHeaders,
        'Content-Security-Policy': ['script-src \'self\'']
      }
    })
  })

  app.on('activate', function () {
    // On macOS it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow();
    }
  });
});

app.on('window-all-closed', function () {
  if (process.platform !== 'darwin') app.quit()
});

ipcMain.on('message', (event, message) => {
  console.log(message);
})





