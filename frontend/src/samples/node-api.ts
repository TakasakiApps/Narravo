var CryptoJs = require(`crypto-js`)

function NewMathDom(){
      let NewMathdom =  Math.floor(Math.random()*10000000000000000)
      return NewMathdom
}


function NewToken(){
  let Newtoken = String(Date.now()*1000)
  let Newkey =  String(NewMathDom())
  // console.log(Newkey,1)
  var encrypted = CryptoJs.AES.encrypt(Newtoken,Newkey);
  // var decrypted = CryptoJs.AES.decrypt(encrypted, Newkey);
  // console.log(decrypted.toString(CryptoJs.enc.Utf8))
}
export default NewToken











// import { lstat } from 'node:fs/promises'
// import { cwd } from 'node:process'
// import { ipcRenderer } from 'electron'

// ipcRenderer.on('main-process-message', (_event, ...args) => {
//   console.log('[Receive Main-process message]:', ...args)
// })

// lstat(cwd()).then(stats => {
//   console.log('[fs.lstat]', stats)
// }).catch(err => {
//   console.error(err)
// })