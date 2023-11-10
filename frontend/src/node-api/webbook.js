const fs = require('fs');
const path = require('path');

const indexPath = './book/index.json';

function readIndex() {
  try {
    // 读取index.json文件
    const indexData = fs.readFileSync(indexPath, 'utf8');
    return JSON.parse(indexData);
  } catch (err) {
    // 文件不存在或解析失败，返回空数组
    return [];
  }
}

function writeIndex(indexData) {
  fs.writeFileSync(indexPath, JSON.stringify(indexData, null, 2));
}

function deleteObjectById(id) {
  let indexData = readIndex();

  // 查找并删除指定id的对象
  indexData = indexData.filter(obj => obj.id !== id);

  writeIndex(indexData);
  return indexData;
}
// 查找name
function findObjectByName(name) {
 let indexData = readIndex();

 // 查找指定name的对象
 const foundObj = indexData.find(obj => obj.name === name);

 return foundObj;
}
function addObject(obj) {
  let indexData = readIndex();

  // 检查id是否重复
  const isIdExists = indexData.some(item => item.id === obj.id);
  if (isIdExists) {
    throw new Error('ID already exists.');
  }

  // 添加新对象
  indexData.push(obj);

  writeIndex(indexData);
  return indexData;
}
function readBookChapter(name, chaptername) {
    const filePath = `./book/${name}/${chaptername}.txt`;
    return new Promise((resolve, reject) => {
        fs.readFile(filePath, 'utf8', (err, data) => {
            if (err) {
                reject(err);
            } else {
                resolve(data);
            }
        });
    });
}
module.exports = {
  deleteObjectById,
  addObject,
  readIndex, // 导出读取index数据的函数
  findObjectByName,
  readBookChapter
};