const fs = require('fs');
const chokidar = require('chokidar');
const nodepath = require('path');
const folderPath = './book'
let fileNames = [];
const pattern = /\\(.*)/;
// 判断后缀名
function checkFileExtension(filename, extension) {
  const fileExtension = filename.slice(filename.lastIndexOf('.') + 1);

  return fileExtension === extension;
}
// 判断是否为文件夹下
function hasExcessiveBackslashes(str) {
  const backslashCount = (str.match(/\\/g) || []).length;
  return backslashCount >= 2 ? false : true;
}
function createDirectory(path) {

  fs.access(path, (error) => {
    console.log('1111')
    if (error) {
      // 路径不存在，创建目录
      fs.mkdir(path, (error) => {
        if (error) {
          console.error('创建目录失败:', error);
          return;
        }
        console.log('目录已成功创建！');
      });
    } else {
      // console.log('路径已存在！');
    }
  });



// 获取当前文件夹下的全部文件名
// const getCurrentFileNames = () => {
//   fs.readdir(folderPath, (err, files) => {
//     if (err) {
//       console.error('无法读取文件夹内容:', err);
//       return;
//     }
//     fileNames =files;
//     // 更新 fileNames 数组
//     console.log('12',files);
//   });
// };

// 初始化，获取当前文件夹下的全部文件名
// getCurrentFileNames();

// 监听文件夹变化
//监听器
 const watcher = chokidar.watch(folderPath, {
  persistent: true // 保持监听状态，即使没有文件改动
});
function readbookdata (path){
  console.log(path)
  const stats = fs.statSync(path);
  console.log(stats)
  console.log(hasExcessiveBackslashes(path))
  // if(hasExcessiveBackslashes(path)){
  // }
  return{
    author : stats.uid,
    modifiedTime : stats.mtime,
    local:1,
    size:stats.size,
    name: 'null',
    path:path
  }
}
  watcher
  .on('add',(filePath)=>{

    let bookdata =  readbookdata(filePath)
    const fileName = nodepath.basename(filePath);
    // console.log(bookdata,'fileName')
    // console.log(bookdata.path,"path")
    bookdata.name = fileName
    if(checkFileExtension(bookdata.name,'txt')){
    // console.log('aaaaaaaaaa',path)
    if(hasExcessiveBackslashes(bookdata.path)){
      delete bookdata.path
      fileNames.push(bookdata)
    }
    // console.log(bookdata);
    // console.log(bookdata,'bookdata')
    
    // console.log(`文件添加: ${fileName}`);
    return fileNames
    }
    
  })
  .on('unlink', (filePath) => {
    const fileName = nodepath.basename(filePath);
    // fileNames = fileNames.filter(fileNames => fileNames !== filePath);

    const unfilename =  pattern.exec(filePath)[1]

    for(let i =0;i< fileNames.length  ;i++){
      if(fileNames[i].name == unfilename){
        fileNames.splice(i,1)
        // console.log(fileNames)
        return fileNames
      }
    }
    console.log(`文件删除: ${fileName},`+fileNames);
    // 在这里可以执行进一步的操作
  });
}
createDirectory(folderPath)
function deleteFile(fileName) {
  const filePath = path.join(__dirname, 'book', fileName);
  fs.unlink(filePath, (err) => {
    if (err) {
      console.error(err);
      return;
    }

    console.log(`成功删除文件: ${fileName}`);
  });
}
// fileWatcher.js
// 存储当前文件夹下的全部文件名的数组

module.exports={
  fileNames
  // createDirectory
};