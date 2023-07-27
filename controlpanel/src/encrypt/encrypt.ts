import {AesCipher, MD5} from "./crypto_impl";

class EncryptResult {
  private data: string;
  private integrityKey: string;
  private timestamp: string;

  constructor(data: string, integrityKey: string, timestamp: string) {
    this.data = data;
    this.integrityKey = integrityKey;
    this.timestamp = timestamp;
  }
}

const generateKey = (): string => {
  const characters = 'abcdefghijklmnopqrstuvwxyz0123456789';
  let result = '';

  for (let i = 0; i < 32; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length);
    result += characters[randomIndex];
  }

  return result;
};

export const Encrypt = (data: Object): EncryptResult => {
  const originalData = JSON.stringify(data);
  const encryptKey = generateKey();

  const milliTimestamp = Date.now().toString();
  const milliTimestampMD5 = MD5.cal(milliTimestamp);

  return new EncryptResult(
    new AesCipher(encryptKey).encrypt(originalData),
    new AesCipher(milliTimestampMD5).encrypt(encryptKey),
    milliTimestamp
  );
};