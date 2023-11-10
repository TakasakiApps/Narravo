import { app, BrowserWindow,ipcMain,nativeTheme,ipcRenderer,Menu } from 'electron'
import path from 'node:path'
import cosi from './conslo'
// const childProcess = require('child_process');
// import { nodeinit} from '../src/node-api/nodeinit'
// import cutHile from '../src/node-api/cutFile'
const fileNames = require('../src/node-api/nodeinit')
const {cutFile}  = require('../src/node-api/cutFile')
const delFile = require('../src/node-api/delflie')
// 分割文章
const cuttxt = require('../src/node-api/txtliemarr');
const readAndSortTxtFiles = require('../src/node-api/readAndSort');
// 配置网络小说相关
const {addObject,deleteObjectById,readIndex,findObjectByName,readBookChapter} = require('../src/node-api/webbook');
// 判断路径模块
const checkFolderExistence = require('../src/node-api/textfiles');
//更新阅读记录
const updateReadIndex = require('../src/node-api/updataIndex');

// The built directory structure
//
// ├─┬─┬ dist
// │ │ └── index.html
// │ │
// │ ├─┬ dist-electron
// │ │ ├── main.jss
// │ │ └── preload.js
// │

process.env.DIST = path.join(__dirname, '../dist')
process.env.PUBLIC = app.isPackaged ? process.env.DIST : path.join(process.env.DIST, '../public')

app.commandLine.appendSwitch('disable-features', 'OutOfBlinkCors');

let win: BrowserWindow | null
// 🚧 Use ['ENV_NAME'] avoid vite:define plugin - Vite@2.x
const VITE_DEV_SERVER_URL = process.env['VITE_DEV_SERVER_URL']
// const workerProcess = childProcess.fork('./src/node-api/index.js');
// workerProcess.on('message', (message) => {
//   console.log('收到子进程的消息:', message);
// });
//处理后缀名
function removeFileExtension(filename) {
  const lastDotIndex = filename.lastIndexOf('.');
  if (lastDotIndex === -1) {
    return filename; // 如果没有找到后缀名分隔符，则直接返回原始文件名
  }

  return filename.slice(0, lastDotIndex); // 返回删除了后缀名的文件名部分
}
// 汉字排序
function compareChineseNumbers(a, b) {
  const chineseNumberRegex = /第(\d+)章/;
  const numberA = parseInt(a.match(chineseNumberRegex)[1]);
  const numberB = parseInt(b.match(chineseNumberRegex)[1]);

  return numberA - numberB;
}
function createWindow() {
  win = new BrowserWindow({
    width:1130,
    height:700,
    minHeight:600,
    minWidth:660,
    icon: path.join(process.env.PUBLIC, 'electron-vite.svg'),
    webPreferences: {
      // 预处理脚本
      preload: path.join(__dirname, 'preload.js'),
      nodeIntegration: true,
      contextIsolation: false,
      webSecurity: false
    },
  })
   nativeTheme.themeSource = 'light';
  // 黑夜模式
  ipcMain.handle("dark", () => {
      //cutHile('11')
      // console.log(cutFile)
      // console.log('1')
      // createDirectory('../book')
      nativeTheme.themeSource = 'dark'
    

    return nativeTheme.themeSource;
  });
  //首页点击
  ipcMain.on('testliem',(e,name)=>{
    // checkFolderExistence(name)
    let newname = removeFileExtension(name)
    if(checkFolderExistence(newname)){
      // 存在同名文件夹
      const sortedFiles = readAndSortTxtFiles(newname)
      // const bleArr = sortedFiles.map(item=> removeFileExtension(item))
      // let newFiles = bleArr.sort(compareChineseNumbers)
      win?.webContents.send('sortedFiles',sortedFiles)
    }else{
      let arr =  cuttxt(newname)
      const newObject = { readin: arr.map(obj => ({name:obj.title})) };
      newObject.name = removeFileExtension(name)
      newObject.local = '1'
      newObject.id = Math.floor(Math.random() * 1000000);
      // console.log(newObject)
      addObject(newObject)

      const sortedFiles = readAndSortTxtFiles(newname)
      // const bleArr = sortedFiles.map(item=> removeFileExtension(item))
      // let newFiles = sortedFiles.sort(compareChineseNumbers)
      win?.webContents.send('sortedFiles',sortedFiles)
    }

  })
  ipcMain.handle('light',()=>{
    nativeTheme.themeSource = 'light'
    return nativeTheme.themeSource
  })
  // 获取书籍信息
  ipcMain.on('Upbookdata', (event, data) => {
    // console.log('接收到来自渲染进程的消息：', data);
    // console.log(fileNames)
    event.reply('GetBookliem', fileNames);
  });
  // 删除书
  ipcMain.on('delbkliem', (event, data) => {
    // console.log('接收到来自渲染进程的消息：', data);
    // console.log(fileNames)
    // event.reply('GetBookliem', fileNames);
  });
  //本地书内容
  ipcMain.on('getchapter',(e,data)=>{
    removeFileExtension(data.name)
    // console.log(data)
    readBookChapter(removeFileExtension(data.name),data.chaptername).then(
      res=>{
        win?.webContents.send('chaptersdata',res)
      }
    ).catch(error => {
    console.error(error);
  });
  })
  //本地书列表
  ipcMain.on('Getpassliem',(e,bookname)=>{
    // console.log('get：',bookname)
    // event.reply('GetBookliem', '列表');
    let newname = removeFileExtension(bookname)

    let filteredArr = findObjectByName(newname)
    //发送给前端
    win?.webContents.send('Getpassliem', filteredArr)
  });
  //读取网络书架
  ipcMain.on('GetwebBook',(e) => {
    const indexData = readIndex();

    const filteredArr = indexData.filter(item => item.local == '0');
    // console.log(filteredArr);
    // event.reply('GetwebBook', indexData);
    win?.webContents.send('GetwebBook', filteredArr)
  });
  ipcMain.on('addbookliem',(e, bookdata)=>{
    try {
      addObject(bookdata)
      win?.webContents.send('ok','ok')
    } catch (error) {
      win?.webContents.send('noaddliem','no')
    }
    
    
    // console.log('addObkjct')
  })
  ipcMain.on('addfrbookliem',(e, bookdata)=>{
    try {
      addObject(bookdata)
      win?.webContents.send('storyOkadd','ok')
    } catch (error) {
      win?.webContents.send('noaddliem','no')
    }
    
    
    // console.log('addObkjct')
  })
  ipcMain.on('Recordindex',(e,id,index,a)=>{
    updateReadIndex(id,index,a)
  })
  //删除网络书本
  ipcMain.on('removewebBoke',(e,id)=>{
    try {
      // console.log(id)
    deleteObjectById(id)
    win?.webContents.send('okremove','ok')
    } catch (error) {
      win?.webContents.send('noremove','no')
    }
  })
  // 删除本地书本
  ipcMain.on('removelocalBook',(e,name)=>{
    try {
      // console.log(name)
    // console.log(delFile)
    delFile(name)
     win?.webContents.send('okdel',name)
    } catch (error) {
      win?.webContents.send('no','no')
    }
  })
  
  app.on('browser-window-blur', () => {
    // const indexData = bookModule.getIndexData();
  // 当窗口失去焦点时调用重新获取焦点的函数
   win?.webContents.send('GetBookliem', fileNames);
  //  win?.webContents.send('GetwebBook', indexData )
  });
  // 设置APP主题模式
  // ipcMain.handle("dark-mode:change", (_, type: "system" | "light" | "dark") => {
  //   nativeTheme.themeSource = type;
  //   return nativeTheme.themeSource;
    
  // });
  
//  win.show =()=>{
  
//  }
  // Test active push message to Renderer-process.
  win.webContents.on('did-finish-load', () => {
    win?.webContents.send('main-process-message', (new Date).toLocaleString());
  })

  if (VITE_DEV_SERVER_URL) {
    win.loadURL(VITE_DEV_SERVER_URL)
  } else {
    // win.loadFile('dist/index.html')
    win.loadFile(path.join(process.env.DIST, 'index.html'))
  }
}
// 隐藏菜单栏
app.on('ready', () => {
  // 创建一个空的菜单栏
  const menu = Menu.buildFromTemplate([]);
  
  // 隐藏菜单栏
  Menu.setApplicationMenu(menu);
});
app.on('window-all-closed', () => {
  win = null
})

// Menu.setApplicationMenu(null) 
app.whenReady().then(createWindow)
