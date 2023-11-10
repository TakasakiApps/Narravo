const fs = require('fs');
const iconv = require('iconv-lite');
const chardet = require('chardet');
const createDirectory  =require('./nodeinit.js')

// console.log(nodeinit)
// const findTextBetweenStrings = require('./1.js')

// let texty =  findTextBetweenStrings("这是一段示例文本，我要在这个文本中查找两个字符串之间的内容。", "示例", "查找")
// console.log(texty)
// nodeinit()
// console.log(nodeinit)

// let res = '陛下难当'
function cutFile(res) {
// 定义一个字符串来存储读取的小说内容
let src = "";
// 文件读取流
let filesP = `./book/${res}.txt`;
// console.log(filesP,'32');


try {
// 从指定路径读取小说
// const filesPath = path.resolve(__dirname, filesP)

const buffer = fs.readFileSync(filesP);
// 解决编码格式
const detectedEncoding = chardet.detect(buffer);
let sce = iconv.decode(buffer, detectedEncoding);

cutCatlog(sce,res);
  
// src = fs.readFileSync(filesP, 'utf-8');

} catch (error) {
  console.log('读取失败',error)
 return{
  type:'404',
  mage:'读取失败'
 }
}


}

function cutCatlog(src,res) {

// 匹配规则
const pest = "(正文){0,1}(\s|\n)(第)([\u4e00-\u9fa5a-zA-Z0-9]{1,7})[章][^\n]{1,35}(|\n)";
// 替换规则
const washpest = "(PS|ps)(.)*(|\n)";
// 将小说内容中的 PS 全部替换为 ""
src = src.replace(new RegExp(washpest, "g"), "");

// 用于存储章节内容
const list = [];
const namelist = [];
// 根据匹配规则将小说分为一章一章的，并存入 list



for (const s of src.split(new RegExp(pest, "g"))) {
  // console.log(s)
list.push(s);
// 
}

// console.log("size" + src.split(new RegExp(pest, "g")).length);
const p = new RegExp(pest, "g");
let m;
let i = 1;
let j = 1;
const newlist = [];
let newstr = null;

// 

// 循环匹配
while ((m = p.exec(src)) !== null) {
newstr = "";

// 替换退格符
let temp = m[0].replace(/ /g, "").replace(/\r/g, "");

if (i === list.length) break;
newstr = temp + list[i];
i++;

newlist.push(newstr);


temp = temp.replace(/[(（].*?[）)]/g, "").replace(/：/g, "");
temp = temp.replace(/[\\"]/g, "").replace(/[\\/]/g, "").replace(/[|]/g, "");

temp = temp.replace(/[?]/g, "").replace(/[*]/g, "").replace(/[(](.)*[)]/g, "");

temp = temp.replace(/\n/g)

temp = temp.replace(/^undefined/, "");
// console.log("j=" + j + " temp=" + temp + ".txt");
j++;
namelist.push(temp);
temp = "";
}

// 创建目录
const bookname = res;
const dirPath = `./book/${bookname}`;
// console.log(fs.existsSync(dirPath),dirPath)
if (!fs.existsSync(dirPath)) {
// console.log('true')
fs.mkdirSync(dirPath);

}
// console.log(fs.existsSync(dirPath))
// 生成章节 TXT 文件
for (i = 0; i < newlist.length; i++) {

const ctl = namelist[i];
// console.log(newlist[i])//
const bloodbath = `${dirPath}/${ctl}.txt`; 
const bookContent = newlist[i];


try {
  fs.writeFileSync(bloodbath, bookContent);
} catch (error) {
  console.error(error);
}
}
}
// 调用或者导出 cutFile() 函数
module.exports ={
  createDirectory,
  cutFile
}
