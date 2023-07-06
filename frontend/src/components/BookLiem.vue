<template>
  <div class="book-liem-main "  v-contextmenu:bookshelf>
    <div class="book-img box-margin">
      <img src="../assets/猫猫.png" alt="" class="book-img-main 
      shadow-effect 
      no-select 
      no-drag
      cursor-pointer">
    </div>
    <div class="book-main box-margin">
      <h1 class="book-name">{{ props.name }}</h1>
      <div class="book-liem-act">
        <div class="book-liem-inln">
          <h5>{{ props.act }}</h5>
        </div>
        <div class="book-liem-inln">

          <el-tag v-for="(ilem, index ) in props.attr" :key="index" class="tag-box" effect="dark">
            {{ ilem }}
          </el-tag>
        </div>
      </div>
    </div>
    <!-- <el-tooltip content="加入书架" placement="bottom">
      <el-button class="book-liem-Brief">
        <el-icon size="25"><Plus /></el-icon>
      </el-button>
    </el-tooltip> -->
  </div>
  <v-contextmenu ref="bookshelf" class="bookshelf">
    <v-contextmenu-item><span>加入书架</span></v-contextmenu-item>
  </v-contextmenu>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

import { ref } from 'vue'
// import {Plus} from '@element-plus/icons-vue'
const router: any = useRouter()
const props = defineProps({
  id: Number,
  name: {
    type: String,
    default: '佚名'
  },
  act: {
    type: String,
    default: '作者',
  },
  Brief: String,
  exist: {
    type: Boolean,
    default: 'ture'
  },
  attr: {
    type: Object,
    default: {
      0: '玄幻',
      1: '有声'
    }
  },

})

</script>

<style scoped>


.box-margin {
  margin-top: 1.25rem;
  margin-left: 3rem;
}

.book-liem-main {
  display: block;
  height: 200px;
  /* border-bottom: 1px solid black; */
  max-width: 900px;
  margin: 10px auto;
  transition: box-shadow 0.5s ease-in-out;
  border-radius: 25px;
}

.book-liem-main:hover {
  box-shadow: 1px 1px 6px 8px #ccc;
  /* border: 1px solid black; */

}

.tag-box {
  margin-right: 1rem;


}

.book-img {
  display: inline-flex;
  float: left;
  max-height: 200px;
  /* box-shadow: 1px 1px 5px 5px #ccc; */
  
  
}

.book-img-main {
  width: 150px;border-radius: 25px;
}

.book-main {
  display: inline;
  float: left;
  text-align: left;
}

.book-liem-act {
  margin-top: 1rem;
  height: 82px;
}

.book-liem-inln {
  display: inline;
}

.book-liem-Brief {
  background-color: rgba(255, 255, 255, 0);
  border: none;
  float: right;
  display: block;

}

.el-icon:hover {
  animation: up 1s;
}

.book-button {
  margin-top: 2rem;
}

.book-stare {
  height: 45px;
  width: 200px;
  margin-right: 1rem;
}

h5 {
  opacity: 0.5;
}

@keyframes up {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}
</style>