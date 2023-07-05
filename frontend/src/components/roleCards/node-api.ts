import path from 'path'
import fs from 'fs'
interface RoleCard {
  name:String,
  rgb:String,
  Reader:String,
  speed:Number,
  tone:Number
}
// const RoleCard={
//   name:'小帅',
//   rgb:'111',
//   Reader:'小帅',
//   speed:1,
//   tone:2
// }
function checkJSONFiles(folderPath: string, fileName: String) {
  const files = fs.readdirSync(folderPath);
   for (const file of files) {
    const filePath = path.join(folderPath, file);

    // 检查文件扩展名是否为 .json
    if (path.extname(filePath) === '.json') {
      // 获取文件名（不包括扩展名）
      const baseName = path.basename(filePath,'.json');

      // 检查文件名是否与传入的名称相同
      if (baseName == fileName) {
        return true; // 找到匹配的文件
      }else{
        console.log(baseName,fileName)
        return false
      }
    }else{
      return false
    }
    
}
}
function Addjson(dirPath: string = '../RoleCard',RoleCard:RoleCard){
  // console.log(dirPath,RoleCard)
  const jsonData = JSON.stringify(RoleCard);
  const filePath = `${dirPath}/${RoleCard.name}.json`;
  
  fs.writeFile(filePath, jsonData, (err: any) => {
    if (err) {
      console.error('发生错误：', err);
    } else {
      console.log(`JSON文件已成功创建：${filePath}`);
    }
  });
}
function readRoledir (dirPath: string = '../RoleCard',RoleCard:RoleCard){
  try{
    fs.accessSync(dirPath);
    if(checkJSONFiles(dirPath,RoleCard.name)){
      // console.log(RoleCard)
      console.log('已有json')
      return false
    }else{
      console.log('没创建json')
      Addjson(dirPath,RoleCard)
    }
  }catch(e){
    fs.mkdirSync(dirPath)
    console.log('没有创建')
  }
}

export default readRoledir