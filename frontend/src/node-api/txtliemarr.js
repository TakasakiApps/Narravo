const fs = require('fs');
const path = require('path');
// 读取文本文件内容
function cuttxt(name){
const filePath = `./book/`+name+`.txt`;
const novelText = fs.readFileSync(filePath, 'utf8');
// 正则表达式匹配章节标题和内容
const chapterRegex = /(第[\d零一二三四五六七八九十百千万亿]+章)\s+([\s\S]*?)(?=\s*第[\d零一二三四五六七八九十百千万亿]+章|$)/g;

// 提取章节标题和内容并存入数组
const chapters = [];
let match;
while ((match = chapterRegex.exec(novelText)) !== null) {
  const title = match[1];
  const content = match[2];
  chapters.push({ title, content });
}


const dirPath = `./book/${name}`;
// console.log(fs.existsSync(dirPath),dirPath)
if (!fs.existsSync(dirPath)) {
// console.log('true')
fs.mkdirSync(dirPath);

}

// 循环遍历每个章节，将内容写入文件
chapters.forEach((chapter) => {
  const { title, content } = chapter;
  const fileName = `${title}.txt`;
  const filePath = path.join(dirPath, fileName);
  
  fs.writeFileSync(filePath, content, 'utf8');
});

console.log('文件输出完成！');
 return chapters
}

module.exports = cuttxt;