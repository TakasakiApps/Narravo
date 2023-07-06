import { app, BrowserWindow,ipcMain,nativeTheme } from 'electron'
import path from 'node:path'

// The built directory structure
//
// ├─┬─┬ dist
// │ │ └── index.html
// │ │
// │ ├─┬ dist-electron
// │ │ ├── main.js
// │ │ └── preload.js
// │
process.env.DIST = path.join(__dirname, '../dist')
process.env.PUBLIC = app.isPackaged ? process.env.DIST : path.join(process.env.DIST, '../public')


let win: BrowserWindow | null
// 🚧 Use ['ENV_NAME'] avoid vite:define plugin - Vite@2.x
const VITE_DEV_SERVER_URL = process.env['VITE_DEV_SERVER_URL']

function createWindow() {
  win = new BrowserWindow({
    width:800,
    height:700,
    minHeight:600,
    icon: path.join(process.env.PUBLIC, 'electron-vite.svg'),
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
      nodeIntegration: true,
        contextIsolation:false
    },
  })
  ipcMain.handle("dark", () => {
    
      nativeTheme.themeSource = 'dark'
    

    return nativeTheme.themeSource;
  });
  ipcMain.handle('light',()=>{
    nativeTheme.themeSource = 'light'
    return nativeTheme.themeSource
  })
  // 设置APP主题模式
  // ipcMain.handle("dark-mode:change", (_, type: "system" | "light" | "dark") => {
  //   nativeTheme.themeSource = type;
  //   return nativeTheme.themeSource;
    
  // });
  
 
  // Test active push message to Renderer-process.
  win.webContents.on('did-finish-load', () => {
    win?.webContents.send('main-process-message', (new Date).toLocaleString())
  })

  if (VITE_DEV_SERVER_URL) {
    win.loadURL(VITE_DEV_SERVER_URL)
  } else {
    // win.loadFile('dist/index.html')
    win.loadFile(path.join(process.env.DIST, 'index.html'))
  }
}

app.on('window-all-closed', () => {
  win = null
})

app.whenReady().then(createWindow)
