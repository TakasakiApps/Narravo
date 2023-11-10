const fs = require('fs');

function updateReadIndex(id, readIndex,a) {
  // 读取'./book/index.json'文件
  
  fs.readFile('./book/index.json', 'utf8', (err, data) => {
    if (err) {
      console.error('读取文件时出错:', err);
      return;
    }

    try {
      // 解析JSON数据
      const index = JSON.parse(data);
      // let id = toString(nid)
      // 根据id查找对应的子元素
      console.log(a)
      let child = null
      if(a){
        child = index.find(item => item.name == id);
      }else{
        child = index.find(item => item.id == id);
      }
      
      if (child) {
        // 检查是否存在readindex属性，如果不存在则添加
        if (!child.hasOwnProperty('readindex')) {
          child.readindex = readIndex;
        } else {
          // 存在readindex属性，直接更新其值
          child.readindex = readIndex;
        }

        // 将更新后的数据写入文件
        fs.writeFile('./book/index.json', JSON.stringify(index, null, 2), 'utf8', (err) => {
          if (err) {
            console.error('写入文件时出错:', err);
          } else {
            console.log('文件更新成功');
          }
        });
      } else {
        console.log('未找到匹配的子元素');
      }
    } catch (error) {
      console.error('解析JSON时出错:', error);
    }
  });
}

module.exports = updateReadIndex;