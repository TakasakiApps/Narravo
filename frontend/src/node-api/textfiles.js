const fs = require('fs');

function checkFolderExistence(folderName) {
  const folderPath = `./book/${folderName}`;
  // console.log(fs.statSync(folderPath))
  try {
    
    const stats = fs.statSync(folderPath);
    
    if (stats.isDirectory()) {
      return true;
    }
  } catch (error) {
    return false;
  }
}

module.exports = checkFolderExistence;