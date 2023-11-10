const fs = require('fs');
const path = require('path');

function readAndSortTxtFiles(Path) {
  let folderPath = './book/'+Path
  try {
    // 读取文件夹下的所有文件
    const files = fs.readdirSync(folderPath);

    // 过滤出以 '.txt' 结尾的文件
    const txtFiles = files.filter(file => path.extname(file) === '.txt');

    // 对文件进行排序
    const sortedTxtFiles = txtFiles.sort((fileA, fileB) => {
      // 提取文件名中的数字部分进行比较
      const numA = parseInt(path.basename(fileA, '.txt').replace('第', '').replace('章', ''));
      const numB = parseInt(path.basename(fileB, '.txt').replace('第', '').replace('章', ''));
      return numA - numB;
    });

    return sortedTxtFiles;
  } catch (error) {
    // 处理错误
    console.error('Error:', error);
    return [];
  }
}

module.exports = readAndSortTxtFiles;