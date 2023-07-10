import {randomBytes, createCipheriv, createDecipheriv, createHash} from 'crypto';

class AesCipher {
  key: Buffer;

  constructor(key: string) {
    if (![16, 24, 32].includes(key.length)) {
      throw new Error('invalid key size');
    }
    this.key = Buffer.from(key);
  }

  encrypt(unencrypted: string): string {
    const plainText = Buffer.from(unencrypted);
    const iv = randomBytes(16);
    const cipher = createCipheriv('aes-256-cbc', this.key, iv);
    let cipherText = cipher.update(plainText);
    cipherText = Buffer.concat([cipherText, cipher.final()]);
    return iv.toString('hex') + cipherText.toString('hex');
  }

  decrypt(encrypted: string): string {
    const iv = Buffer.from(encrypted.slice(0, 32), 'hex');
    const cipherText = Buffer.from(encrypted.slice(32), 'hex');
    const decipher = createDecipheriv('aes-256-cbc', this.key, iv);
    let plainText = decipher.update(cipherText);
    plainText = Buffer.concat([plainText, decipher.final()]);
    return plainText.toString();
  }
}

class MD5 {
  static cal(data: string) {
    const md5Hash = createHash('md5');

    md5Hash.update(data, 'utf-8');

    return md5Hash.digest('hex');
  }
}

export {
  AesCipher,
  MD5
};