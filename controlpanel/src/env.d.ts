/ <reference types="vite/client" />

declare module '*.vue' {
  import { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}
declare module 'vue-cookies'
declare module 'vue-cropper';
declare module 'vue-cropper-next';
declare module 'colorpicker-v3';
declare module '/src/http/client';