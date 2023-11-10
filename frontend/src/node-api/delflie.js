const fs = require('fs');
// const path = require('path');

function delFile(fileName) {
  console.log('1111')
  const filePath = `./book/${fileName}`;

  fs.unlink(filePath, (err) => {
    if (err) {
      console.error(err);
      return;
    }

    console.log(`成功删除文件: ${fileName}`);
  });
}

module.exports = delFile;