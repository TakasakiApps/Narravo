<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="80px">
        <el-avatar :size="68" :src="circleUrl" />
        <div class="iconBox">
          <button class="icon" :class="[btnmenu == 0 ? 'btn' : '']" @click="btnmenuClick(0, '/')">
            <el-icon size="36">
              <Reading />
            </el-icon>
          </button>
        </div>

        <div class="iconBox">
          <button class="icon" :class="[btnmenu == 1 ? 'btn' : '']" @click="btnmenuClick(1, '/writer')">
            <el-icon size="36">
              <EditPen />
            </el-icon>
          </button>
        </div>
        <div class="iconBox">
          <button class="icon" :class="[btnmenu == 2 ? 'btn' : '']" @click="btnmenuClick(2, '/search')">
            <el-icon size="36">
              <svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" width="30" height="30">
                <path
                  d="M967.9234 816.206954l-147.416093-192.48379c52.29229-49.884092 85.146985-119.721821 85.146985-197.472199 0-150.684361-122.990089-273.330422-274.018478-273.330422-151.200403 0-274.018478 122.646061-274.018478 273.330422s122.990089 273.330422 274.018478 273.330422c48.679993 0 94.263733-12.901058 133.998992-35.090879l147.932135 193.171846c6.70855 8.77272 16.857383 13.417101 27.350244 13.417101 7.224593 0 14.621199-2.236183 20.641693-7.052579C976.524105 852.501932 979.448345 831.172182 967.9234 816.206954L967.9234 816.206954zM421.090543 426.250966c0-115.765496 94.435747-209.857215 210.545271-209.857215s210.545271 94.091718 210.545271 209.857215-94.435747 209.857215-210.545271 209.857215S421.090543 542.016462 421.090543 426.250966L421.090543 426.250966zM288.467663 288.811692 83.254829 288.811692c-18.921552 0-34.230808-15.309256-34.230808-34.058794s15.309256-34.058794 34.230808-34.058794l205.212834 0c18.921552 0 34.230808 15.309256 34.230808 34.058794C322.698471 273.502436 307.389216 288.811692 288.467663 288.811692L288.467663 288.811692zM280.210986 626.131362 83.254829 626.131362c-18.921552 0-34.230808-15.309256-34.230808-34.058794s15.309256-34.058794 34.230808-34.058794l196.956157 0c18.921552 0 34.230808 15.309256 34.230808 34.058794C314.441794 610.822107 299.132538 626.131362 280.210986 626.131362L280.210986 626.131362zM239.78767 457.041492l-156.532841 0c-18.921552 0-34.230808-15.309256-34.230808-34.230808 0-18.921552 15.309256-34.230808 34.230808-34.230808l156.532841 0c18.921552 0 34.230808 15.309256 34.230808 34.230808C273.846464 441.732236 258.537208 457.041492 239.78767 457.041492L239.78767 457.041492zM460.137746 801.241727 83.254829 801.241727c-18.921552 0-34.230808-15.309256-34.230808-34.230808s15.309256-34.230808 34.230808-34.230808l376.882916 0c18.921552 0 34.230808 15.309256 34.230808 34.230808C494.368554 785.932471 479.059298 801.241727 460.137746 801.241727L460.137746 801.241727z"
                  fill="#000"></path>
              </svg>
            </el-icon></button>
        </div>
        <div class="setBox" >
          <button class="set" :class="[btnmenu == 3 ? 'btn' : '']" @click="btnmenuClick(3, '/setting')">
          <el-icon size="36">
            <Setting />
          </el-icon>
        </button>

        </div>
        

      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>
  
<script setup lang='ts'>
import { ref } from 'vue'
import { Reading, EditPen, Setting } from '@element-plus/icons-vue'

import router from './router'

const circleUrl = ref('https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png')

const btnmenu = ref(0)
const btnmenuClick = (id, path) => {
  btnmenu.value = id
  router.push({ path })
  console.log(path,id);
  
}

</script>
  
<style scoped lang="scss">
$background-color: rgb(239, 122, 122);
$icon-color: #ccc;

.btn {
  
  background-color: $background-color !important;
  animation: big 0.2s;
}

.common-layout,
.el-container {


  height: 100%;
  min-height: auto;

  .el-aside {

    display: flex;
    flex-direction: column;
    background-color: #eeecec;
    border-radius: 0 8px 8px 0;
    width: 90px;
    align-items: center;
   
    overflow: auto;

    .el-avatar {
      width: 68px;
      height: 68px;
      margin: auto;
      
      border: 1px solid #929290;
    }

    .iconBox {
      display: flex;
      width: 68px;
      height: 68px;
      margin-bottom: 20%;
      align-items: center;
      justify-content: center;

      .icon {
        width: 60px;
        height: 60px;
        border-radius: 25%;
        background-color: $icon-color;
        align-items: center;
        justify-content: center;
        display: flex;
        border: none;
      }
    }


    .icon:hover,
    .set:hover {
      width: 68px;
      height: 68px;
      background-color: $background-color;
    }

    .setBox{
      display: flex;
      width: 68px;
      height: 68px;
      margin: auto;
      align-items: center;
      justify-content: center;

    }
    .set {
      width: 60px;
      height: 60px;
      border-radius: 50%;
      
      background-color: #ccc;
      border: none;
      
    }
  }
}
@keyframes big{
  from{width: 60px; height: 60px;}
  to{width: 68px; height: 68px;}
}
</style>

