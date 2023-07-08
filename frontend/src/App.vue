
<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="80px" :class="[globalVars.isRight == 1 ? 'close' : '', globalVars.isRight == 2 ? 'open' : '']">
        <div class="iconBox">
          <div class="asideIcon" v-if="show" @click="back">
            <svg fill="currentColor" t="1688701539082" class="icon" viewBox="0 0 1025 1024" version="1.1"
              xmlns="http://www.w3.org/2000/svg" p-id="3294" width="28" height="28">
              <path
                d="M533.353412 682.285176c0-0.030118 0.060235-0.060235 0.060235-0.090353C533.172706 682.194824 533.142588 682.224941 533.353412 682.285176zM40.176941 555.008l314.398118 314.398118c27.045647 27.045647 70.927059 27.045647 97.942588 0 27.045647-26.985412 27.045647-70.866824 0-97.912471l-196.186353-196.216471 665.630118 0c38.249412 0 69.240471-31.021176 69.240471-69.240471 0-38.279529-31.021176-69.270588-69.240471-69.270588L256.361412 436.766118l196.186353-196.186353c27.045647-27.015529 27.045647-70.927059 0-97.942588C439.024941 129.114353 421.285647 122.337882 403.546353 122.337882c-17.709176 0-35.448471 6.776471-48.941176 20.299294L40.176941 457.065412C13.101176 484.111059 13.101176 527.932235 40.176941 555.008z"
                p-id="3295"></path>
            </svg>
          </div>
          <div class="asideIcon" :class="[btnmenu == 0 ? 'btn' : '']" @click="btnmenuClick(0, '/')">
            <svg fill="currentColor" t="1688650000398" viewBox="0 0 1152 1024" version="1.1"
              xmlns="http://www.w3.org/2000/svg" p-id="43658" xmlns:xlink="http://www.w3.org/1999/xlink" width="32"
              height="32">
              <path
                d="M224 128l-192 0c-17.6 0-32 14.4-32 32l0 704c0 17.6 14.4 32 32 32l192 0c17.6 0 32-14.4 32-32l0-704c0-17.6-14.4-32-32-32zM192 320l-128 0 0-64 128 0 0 64z"
                p-id="43659"></path>
              <path
                d="M544 128l-192 0c-17.6 0-32 14.4-32 32l0 704c0 17.6 14.4 32 32 32l192 0c17.6 0 32-14.4 32-32l0-704c0-17.6-14.4-32-32-32zM512 320l-128 0 0-64 128 0 0 64z"
                p-id="43660"></path>
              <path
                d="M765.088 177.472l-171.456 86.4c-15.712 7.904-22.112 27.264-14.176 42.976l287.968 571.552c7.904 15.712 27.264 22.112 42.976 14.176l171.456-86.4c15.712-7.904 22.112-27.264 14.176-42.976l-287.968-571.552c-7.904-15.712-27.264-22.112-42.976-14.176z"
                p-id="43661"></path>
              <path d="M928 864c0 17.664-14.336 32-32 32s-32-14.336-32-32c0-17.664 14.336-32 32-32s32 14.336 32 32z"
                p-id="43662"></path>
            </svg>
            <p class="asideFont">书架</p>
          </div>
          <div class="asideIcon" :class="[btnmenu == 1 ? 'btn' : '']" @click="btnmenuClick(1, '/writer')">
            <svg fill="currentColor" t="1688650152027" viewBox="0 0 1024 1024" version="1.1"
              xmlns="http://www.w3.org/2000/svg" p-id="44797" xmlns:xlink="http://www.w3.org/1999/xlink" width="32"
              height="32">
              <path
                d="M603.733333 181.333333L386.133333 401.066667c-6.4 6.4-10.666667 14.933333-12.8 25.6l-51.2 211.2c-8.533333 38.4 23.466667 74.666667 61.866667 64l200.533333-53.333334c8.533333-2.133333 17.066667-6.4 23.466667-14.933333l234.666667-236.8V853.333333c0 40.533333-32 72.533333-70.4 74.666667H170.666667c-40.533333 0-74.666667-34.133333-74.666667-74.666667V256c0-40.533333 34.133333-74.666667 74.666667-74.666667h433.066666z"
                p-id="44798"></path>
              <path
                d="M738.133333 147.2L435.2 448c-4.266667 4.266667-6.4 8.533333-8.533333 14.933333l-32 125.866667c-6.4 23.466667 14.933333 44.8 38.4 38.4l128-29.866667c6.4-2.133333 10.666667-4.266667 14.933333-8.533333l300.8-302.933333c38.4-38.4 38.4-102.4 0-140.8s-100.266667-38.4-138.666667 2.133333z"
                p-id="44799"></path>
            </svg>
            <p class="asideFont">制作</p>
          </div>
          <div class="asideIcon" :class="[btnmenu == 3 ? 'btn' : '']" @click="btnmenuClick(3, '/role')">
            <svg fill="currentColor" t="1688631768366" viewBox="0 0 1024 1024" version="1.1"
              xmlns="http://www.w3.org/2000/svg" p-id="25773" width="32" height="32">
              <path
                d="M902.4 635.1c-5.3-2.4-99.1-43.2-204.3-65.3 54.1-34.5 86.9-94.3 86.9-158.5v-112C785 195.5 700.7 111 597.1 111c-53.9 0-105.2 23.3-140.7 63.9-30.4 34.3-47.2 78.6-47.3 124.4v112c0 63.2 31.4 119.2 79.3 153.3-115.7 20-225.4 67.9-230.5 70.2-31.2 12.8-54.6 47.8-54.6 81.6v135.2h0.6c4.6 34.9 34.3 60.9 69.5 60.9h612.8c30.7 0 58.1-20 68.2-49.9l1.8-5.3V716.4c0-34-23.6-69.2-53.8-81.3"
                p-id="25774"></path>
              <path
                d="M308.6 557.9c37.4-11.7 55.1-18.2 55.1-18.2s24.2-10.7 17.8-32.1c-15-28.5-25.7-50.6-25.7-134V265.2s6.4-52.7 26.7-57.7c1.6-1.7 1.9-4.3 1.2-7.2-5.3-12.6-17.7-20.9-31.5-20.9-78.3 0.1-141.8 63.6-141.8 142v90.1c0 35 13.2 66.7 34.3 91.5-67.1 20.6-125.7 53.1-136.4 59.3-24.4 11.6-41.9 39.3-41.9 66.8v98.3c0 18.9 15.3 34.2 34.2 34.2 17.3-0.1 31.7-13.2 33.5-30.4 0.1-1.3 0.8-2.4 0.8-3.8v-15.7c2.7-32.4 14.9-91.2 65.2-118.8 8-3.5 16.1-6.9 24.3-10.1 28.2-7.9 56.3-16.2 84.2-24.9"
                p-id="25775"></path>
            </svg>
            <p class="asideFont">角色</p>
          </div>
          <div class="asideIcon" :class="[btnmenu == 2 ? 'btn' : '']" @click="btnmenuClick(2, '/search')">
            <svg fill="currentColor" t="1688650220888" class="icon" viewBox="0 0 1024 1024" version="1.1"
              xmlns="http://www.w3.org/2000/svg" p-id="47861" xmlns:xlink="http://www.w3.org/1999/xlink" width="32"
              height="32">
              <path
                d="M512 85.333333c235.648 0 426.666667 191.018667 426.666667 426.666667s-191.018667 426.666667-426.666667 426.666667S85.333333 747.648 85.333333 512 276.352 85.333333 512 85.333333z m255.317333 229.973334c9.173333-31.786667-20.778667-61.696-52.608-52.565334l-4.352 1.493334L408.746667 384.853333l-4.522667 2.133334a42.666667 42.666667 0 0 0-16.64 16.298666l-2.645333 5.333334-120.661334 301.696-1.493333 4.352c-9.130667 31.829333 20.778667 61.781333 52.608 52.608l4.352-1.493334 301.653333-120.661333 4.522667-2.133333a42.666667 42.666667 0 0 0 16.682667-16.256l2.56-5.376 120.746666-301.696 1.450667-4.352zM512 426.666667a85.333333 85.333333 0 1 1 0 170.666666 85.333333 85.333333 0 0 1 0-170.666666z"
                p-id="47862"></path>
            </svg>
            <p class="asideFont">书城</p>
          </div>
        </div>

        <div class="user">
          <div v-if="!isAvatar">
            <el-popover popper-style="display:flex; width:300px;height:130px;" placement="right" trigger="hover"
              >
              <div style="width: 64px; margin-right: 15px;
             display: flex; align-items: center; justify-content: center;">
                <el-avatar :src="circleUrl" style="width: 64px; height: 64px;" />
              </div>
              <div style=" width: 225px;">
                <h3 style="font-size: 22px; margin: 0; color:#333333">{{ userinfo.username }}</h3>
                <div style="margin: 6px 0;">
                  <p style="margin: 0; font-size: 12px; color: #7F7F7F ;">上次同步时间:</p>
                  <p style="margin: 0; font-size: 12px; color: #7F7F7F;">{{ userinfo.logintime }}</p>
                </div>
                <div style="display: flex; margin-left: 36px;">
                  <button style="margin:0 auto; background-color: #EC808D; 
                color: #F2F2F2; border-radius: 5px; border: none; height: 28px; width: 84px;"
                    @click="isSetPwd = true">更改密码</button>
                  <button style="margin:0 auto;background-color: #EC808D; 
                color: #F2F2F2; border-radius: 5px; border: none; height: 28px; width: 64px;"
                    @click="loginOut">退出</button>
                </div>
              </div>
              <template #reference>
                <el-avatar :src="circleUrl" />
              </template>
            </el-popover>
          </div>
          <div v-if="isAvatar">
            <el-popover popper-style="display:flex; width:300px;height:130px;" placement="right" trigger="hover"
              >
              <div style="width: 64px; margin-right: 15px;
             display: flex; align-items: center; justify-content: center;">
                <el-avatar :src="circleUrl" style="width: 64px; height: 64px;" />
              </div>
              <div style=" width: 225px;">
                <h3 style="font-size: 22px; margin: 0; color:#333333;">未登录</h3>
                <div style="margin: 6px 0;">
                  <p style="margin: 0; font-size: 12px; color: #7F7F7F;">上次同步时间:</p>
                  <p style="margin: 0; font-size: 12px; color: #7F7F7F;;">登录后启用云同步功能</p>
                </div>
                <div style="display: flex; margin-left: 36px;">
                  <button style="background-color: #EC808D; 
                color: #F2F2F2; border-radius: 5px; border: none; height: 28px; width: 64px; "
                    @click="isSign = true">注册</button>
                  <button style="margin-left:30px; background-color: #EC808D; 
                color: #F2F2F2; border-radius: 5px; border: none; height: 28px; width: 64px;" @click="isLogin = true">
                    登录
                  </button>
                </div>
              </div>
              <template #reference>
                <el-avatar :src="circleUrl" />
              </template>
            </el-popover>
          </div>
          <div class="login">
            <el-dialog v-model="isLogin" align-center center :show-close="false" lock-scroll style="width:448px;  box-shadow: 0 15px 20px rgba(0, 0, 0, 0.1);
              background:#F9F9F9; border-radius:15px;">
              <div style="display:flex; flex-direction:column; align-items:center
           ">
                <h2>Welcome</h2>
                <input type="text" placeholder="账号" v-model="login.username" style="padding-left: 10px;" />
                <input type="password" placeholder="密码" v-model="login.password"
                  style="padding-left: 10px; margin-bottom:8px" />

                <button @click="loginUp">登录</button>
                <a style="font-size: 12px; margin-top: 5px; color: #797979; font-family: '更纱黑体'; margin-top:10px"
                  @click="guest" href="#">游客登录</a>
              </div>
            </el-dialog>
          </div>
          <div class="setPwd">
            <el-dialog v-model="isSetPwd" align-center center :show-close="false" lock-scroll style="width:448px;  box-shadow: 0 15px 20px rgba(0, 0, 0, 0.1);
              background:#F9F9F9; border-radius:15px;">
<div style="display:flex; flex-direction:column; align-items:center
           ">
                <h2>设置新密码</h2>

                <input type="text" placeholder="账号" v-model="sign.username" style="padding-left: 10px;" />
                <input type="password" placeholder="旧密码" v-model="sign.password" style="padding-left: 10px;" />
                <input type="password" placeholder="新密码" v-model="sign.password" style="padding-left: 10px;" />
                <button @click="setPwdUp">确认</button>
              </div>
            </el-dialog>
          </div>
          <div class="Sign">
            <el-dialog v-model="isSign" align-center center :show-close="false" lock-scroll style="width:448px;  box-shadow: 0 15px 20px rgba(0, 0, 0, 0.1);
              background:#F9F9F9; border-radius:15px;">
              <div style="display:flex; flex-direction:column; align-items:center
           ">
                <h2>Sign up</h2>

                <input type="text" placeholder="账号" v-model="sign.username" style="padding-left: 10px;" />
                <input type="password" placeholder="密码" v-model="sign.password" style="padding-left: 10px;" />
                <button @click="signUp">注册</button>
              </div>

            </el-dialog>
          </div>
          <div class="switch" @click="btnSwitch" :class="[isDark == true ? 'dark' : '']">
            <div class="sun" v-show="isDark == false">
              <svg fill="currentColor" t="1688632390516" viewBox="0 0 1024 1024" version="1.1"
                xmlns="http://www.w3.org/2000/svg" p-id="30668" width="26.25" height="26.25">
                <path
                  d="M512 187.733333c-177.493333 0-324.266667 146.773333-324.266667 324.266667s146.773333 324.266667 324.266667 324.266667 324.266667-146.773333 324.266667-324.266667-146.773333-324.266667-324.266667-324.266667zM512 153.6c10.24 0 17.066667-6.826667 17.066667-17.066667V34.133333c0-10.24-6.826667-17.066667-17.066667-17.066666s-17.066667 6.826667-17.066667 17.066666v102.4c0 10.24 6.826667 17.066667 17.066667 17.066667zM310.613333 194.56c3.413333 6.826667 6.826667 10.24 13.653334 10.24 3.413333 0 6.826667 0 10.24-3.413333 6.826667-3.413333 10.24-13.653333 6.826666-23.893334L286.72 88.746667c-3.413333-6.826667-13.653333-10.24-20.48-6.826667-10.24 6.826667-13.653333 17.066667-6.826667 23.893333l51.2 88.746667zM88.746667 286.72l88.746666 51.2c3.413333 3.413333 6.826667 3.413333 10.24 3.413333 6.826667 0 10.24-3.413333 13.653334-10.24 3.413333-6.826667 3.413333-17.066667-6.826667-20.48L105.813333 259.413333c-6.826667-6.826667-17.066667-3.413333-23.893333 6.826667s0 17.066667 6.826667 20.48zM153.6 512c0-10.24-6.826667-17.066667-17.066667-17.066667H34.133333c-10.24 0-17.066667 6.826667-17.066666 17.066667s6.826667 17.066667 17.066666 17.066667h102.4c10.24 0 17.066667-6.826667 17.066667-17.066667zM177.493333 686.08l-88.746666 51.2c-6.826667 3.413333-10.24 13.653333-6.826667 23.893333 3.413333 6.826667 10.24 10.24 13.653333 10.24 3.413333 0 6.826667 0 10.24-3.413333L194.56 716.8c10.24-6.826667 10.24-17.066667 6.826667-23.893333-3.413333-10.24-13.653333-13.653333-23.893334-6.826667zM334.506667 822.613333c-10.24-3.413333-20.48-3.413333-23.893334 6.826667l-51.2 88.746667c-3.413333 6.826667-3.413333 17.066667 6.826667 23.893333 3.413333 0 6.826667 3.413333 10.24 3.413333 6.826667 0 10.24-3.413333 13.653333-10.24l51.2-88.746666c3.413333-10.24 0-20.48-6.826666-23.893334zM512 870.4c-10.24 0-17.066667 6.826667-17.066667 17.066667v102.4c0 10.24 6.826667 17.066667 17.066667 17.066666s17.066667-6.826667 17.066667-17.066666v-102.4c0-10.24-6.826667-17.066667-17.066667-17.066667zM713.386667 829.44c-3.413333-6.826667-13.653333-10.24-23.893334-6.826667-6.826667 3.413333-10.24 13.653333-6.826666 23.893334l51.2 88.746666c3.413333 6.826667 10.24 10.24 13.653333 10.24 3.413333 0 6.826667 0 10.24-3.413333 6.826667-3.413333 10.24-13.653333 6.826667-23.893333l-51.2-88.746667zM935.253333 737.28l-88.746666-51.2c-6.826667-3.413333-17.066667-3.413333-23.893334 6.826667s-3.413333 17.066667 6.826667 23.893333l88.746667 51.2c3.413333 0 6.826667 3.413333 10.24 3.413333 6.826667 0 10.24-3.413333 13.653333-10.24 3.413333-10.24 0-20.48-6.826667-23.893333zM989.866667 494.933333h-102.4c-10.24 0-17.066667 6.826667-17.066667 17.066667s6.826667 17.066667 17.066667 17.066667h102.4c10.24 0 17.066667-6.826667 17.066666-17.066667s-6.826667-17.066667-17.066666-17.066667zM836.266667 341.333333c3.413333 0 6.826667 0 10.24-3.413333l88.746666-51.2c6.826667-3.413333 10.24-13.653333 6.826667-23.893333-3.413333-6.826667-13.653333-10.24-23.893333-6.826667L829.44 307.2c-10.24 6.826667-10.24 17.066667-6.826667 27.306667 3.413333 3.413333 10.24 6.826667 13.653334 6.826666zM692.906667 201.386667c0 3.413333 3.413333 3.413333 6.826666 3.413333 6.826667 0 10.24-3.413333 13.653334-10.24l51.2-88.746667c3.413333-6.826667 3.413333-17.066667-6.826667-23.893333-6.826667-3.413333-17.066667-3.413333-23.893333 6.826667L682.666667 177.493333c-3.413333 10.24 0 20.48 10.24 23.893334z"
                  p-id="30669"></path>
              </svg>
            </div>
            <div class="moon" v-show="isDark == true">
              <svg fill="currentColor" t="1688632491243" viewBox="0 0 1024 1024" version="1.1"
                xmlns="http://www.w3.org/2000/svg" p-id="31690" width="30" height="30">
                <path
                  d="M524.8 938.666667h-4.266667a439.893333 439.893333 0 0 1-313.173333-134.4 446.293333 446.293333 0 0 1-11.093333-597.333334 432.213333 432.213333 0 0 1 170.666666-116.906666 42.666667 42.666667 0 0 1 45.226667 9.386666 42.666667 42.666667 0 0 1 10.24 42.666667 358.4 358.4 0 0 0 82.773333 375.893333 361.386667 361.386667 0 0 0 376.746667 82.773334 42.666667 42.666667 0 0 1 54.186667 55.04A433.493333 433.493333 0 0 1 836.266667 810.666667a438.613333 438.613333 0 0 1-311.466667 128z"
                  p-id="31691"></path>
              </svg>
            </div>
          </div>
          <div class="asideIcon" :class="[btnmenu == 4 ? 'btn' : '']" @click="btnmenuClick(4, '/setting')">
            <svg fill="currentColor" t="1688633582143" viewBox="0 0 1024 1024" version="1.1"
              xmlns="http://www.w3.org/2000/svg" p-id="35296" width="30" height="30">
              <path
                d="M885.196958 871.528145c-14.753141-5.464126-25.134981-8.742602-34.970408-12.56749-30.052695-11.474665-59.012564-24.588568-89.611671-34.970408-8.742602-2.732063-21.310093-1.639238-29.506282 2.732063-21.310093 10.38184-42.073773 22.402918-61.744627 36.063234-6.556952 4.371301-12.021078 14.206728-13.660316 22.402918-6.556952 36.063234-11.474665 72.67288-15.845966 109.282526-2.732063 20.76368-12.56749 30.052695-33.877583 29.506282-62.837452-0.546413-125.674905-0.546413-188.512357 0-20.76368 0-31.691933-8.196189-34.423996-28.959869-4.917714-36.609646-10.928253-72.67288-15.299554-109.282526-1.639238-12.56749-6.010539-20.217267-18.031617-26.227806-20.217267-9.835427-38.248884-23.495743-58.466151-33.33117-7.103364-3.824888-18.031617-5.464126-25.681394-2.732063-35.516821 12.56749-69.940817 26.774219-104.911225 40.434535-25.134981 9.835427-32.238345 7.649777-46.445074-16.392379-31.14552-53.002025-61.744627-106.00405-92.343735-159.006075-14.206728-24.588568-13.660316-30.599107 8.742602-48.084311 29.506282-23.495743 59.558977-45.898661 88.518846-69.940817 5.464126-4.371301 9.289015-14.206728 9.289015-21.310093 1.092825-20.217267-3.278476-41.52736 0-61.198215 3.824888-21.310093-4.917714-32.238345-20.217267-43.71301-27.320632-19.670855-52.455613-41.52736-79.229831-61.198215-15.845966-12.021078-19.670855-24.588568-9.289015-42.073773 33.33117-56.280501 66.115928-113.107414 98.900686-169.934328 9.289015-15.845966 21.310093-19.124442 38.248884-12.56749 36.063234 14.206728 72.126467 28.959869 108.736113 42.073773 7.649777 2.732063 19.670855 1.092825 27.320632-2.732063 20.76368-10.38184 39.341709-24.042156 60.105389-34.423996 10.928253-5.464126 14.753141-12.021078 15.845966-22.94933 4.917714-37.156059 10.928253-74.312118 15.845966-111.468177C386.322226 7.649777 397.796892 0 418.014159 0c62.837452 0.546413 125.674905 0.546413 188.512357 0 19.670855 0 30.599107 8.196189 32.784758 27.867044 4.917714 36.609646 10.928253 72.67288 15.299554 109.282526 1.639238 13.660316 6.556952 21.856505 19.670855 27.867044 19.124442 8.742602 37.156059 19.670855 54.641263 31.691933 11.474665 8.196189 21.310093 8.196189 33.877583 2.732063 33.877583-14.206728 68.301579-26.774219 102.725575-40.980947 19.124442-7.649777 32.238345-3.824888 42.620185 14.753141 31.14552 55.187676 63.930278 109.828939 95.62221 164.470202 12.56749 21.856505 10.928253 28.959869-9.289015 44.805836-28.959869 22.402918-57.919739 45.352248-87.426021 67.755166-9.835427 7.649777-12.56749 14.753141-11.474665 27.320632 2.185651 22.94933-0.546413 46.445074 0.546413 69.394404 0.546413 8.196189 4.371301 18.031617 10.38184 22.94933 28.959869 24.042156 59.012564 46.445074 88.518846 69.394404 20.217267 15.845966 21.310093 22.94933 8.742602 44.805836-32.238345 55.734088-64.47669 110.921764-97.261448 166.10944C901.042924 860.05348 891.753909 865.517606 885.196958 871.528145zM510.357893 692.304803c98.900686 0.546413 180.862581-79.776244 181.408993-177.584105 0.546413-100.539924-79.229831-181.955406-178.67693-182.501819-99.447099-0.546413-181.408993 80.322657-181.408993 179.769755C331.134551 610.342908 412.00362 691.75839 510.357893 692.304803z"
                p-id="35297"></path>
            </svg>
          </div>
        </div>
      </el-aside>
      <!-- 添加控制箭头是否显示 -->
      <div class="left" @click="checkRight" v-if="globalVars.isHidden">
        <svg t="1688729279262" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
          p-id="2586" width="20" height="20" v-if="globalVars.isRight == 2">
          <path d="M610.88 512L192 93.12 285.12 0l512 512-512 512L192 930.88z" fill="#262626" p-id="2587"></path>
        </svg>
        <svg t="1688729212223" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
          p-id="1610" width="20" height="20" v-if="globalVars.isRight == 1">
          <path d="M378.24 512l418.88 418.88L704 1024 192 512l512-512 93.12 93.12z" fill="#262626" p-id="1611"></path>
        </svg>
      </div>

      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>
<script setup lang='ts'>

import { ref, inject, watch, reactive, onBeforeMount } from 'vue'
import { client } from './http/client'
import { useColorMode } from '@vueuse/core'
import router from './router'
import { ipcRenderer } from 'electron'


//用户头像
const circleUrl = ref('https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png')
//图片点击高亮显示
const btnmenu = ref(10)
const btnmenuClick = (id: any, path: any) => {
  btnmenu.value = id
  router.push({ path })
}
//切换暗——亮模式
const isDark = ref(false)
useColorMode({
  // 如果模式为auto也需要回显回auto
  emitAuto: true,
  // 默认模式先默认auto,后续通过Electorn拿到当前App主题
  initialValue: "auto",
});
const btnSwitch = () => {
  isDark.value = !isDark.value
  if (isDark.value == true) {
    ipcRenderer.invoke("dark")
  } else {
    ipcRenderer.invoke('light')
  }
}
//用户信息
const userinfo = reactive({
  username: localStorage.username,
  logintime: '',

})
//切换头像弹出窗是否显示 true:未登录,false:已登录
const isAvatar = ref(true)
//显示登录,注册按钮
const isLogin = ref(false)
const isSign = ref(false)

//获取时间的函数
const addZero = (v) => {
  if (v < 10)
    return '0' + v;
  return v.toString();
}
const getDate = () => {
  var d = new Date();
  var str = d.getFullYear().toString() + '-' + addZero(d.getMonth() + 1) + '-' + addZero(d.getDate()) + ' ' + addZero(d.getHours()) + ':' + addZero(d.getMinutes()) + ':' + addZero(d.getSeconds());
  return str
}
//检查本地是否存的有登录信息，有就显示（已登录），没有则显示（未登录）
onBeforeMount(() => {
  if(localStorage.token){
    isAvatar.value = false
    userinfo.logintime = localStorage.logintime
    userinfo.username = localStorage.username
  }
})
//登录，注册功能实现
//1.注册
const sign = reactive({
  username: '',
  password: ''
})
const signUp = () => {
  //需要转成字符串发送数据
  client.post('/api/v1/auth/register', JSON.stringify(sign)).then(res => {
    if (res.status == 200) {
      ElNotification.success({ title: '注册成功!', message: '恭喜你注册成功' })
      //拿到token
      client.post('/api/v1/auth/login', JSON.stringify(sign)).then(res => {
        //注册成功后去获取token
        localStorage.token = res.data.token,
          //保存用户的账号
          localStorage.username = sign.username
        userinfo.username = localStorage.username
        //获取登录时间
        userinfo.logintime = getDate()
        localStorage.logintime = getDate()
        //关闭注册框
        isSign.value = false
      }).catch(err => {
        console.log(err);
      })
      console.log(res);
      //显示登录头像
      isAvatar.value = false
    }
    else if (res.status == 401) {
      ElNotification.error({ title: '错误', message: '该用户已被注册' })
    }
    else
      ElNotification.error({ title: '注册失败' })

  }).catch(err => {
    console.log(err);
  })

}
//2.登录
const login = reactive({
  username: '',
  password: ''
})
const loginUp = () => {
  if (login.username != '' && login.password != '') {
    client.post('/api/v1/auth/login', JSON.stringify(login)).then(res => {
      ElNotification({ title: '登录成功！', message: '欢迎您', type: 'success' })
      localStorage.token = res.data.token
      localStorage.username = sign.username
        userinfo.username = localStorage.username
        //获取登录时间
        userinfo.logintime = getDate()
        localStorage.logintime = getDate()
        //关闭登录框
        isLogin.value = false
    })
  } else {
    ElNotification({ message: '账号或密码不能为空！', type: 'warning' })
  }
}
//3.游客登录
const guest = () => {
  const guests = {
    username: 'guest',
    password: 'guest'
  }
  client.post('/api/v1/auth/login', JSON.stringify(guests)).then(res => {
    ElNotification({ title: '游客登录成功！', message: '你当前使用的是游客账号', type: 'success' })
    localStorage.token = res.data.token
    localStorage.username = '游客'
    userinfo.username = localStorage.username
    localStorage.logintime = getDate()
    userinfo.logintime = localStorage.logintime
    isLogin.value = false
  }).catch(err => {
    console.log(err);
  })
}
//4.更改密码功能
const isSetPwd = ref(false) //控制更改密码页面是否显示

  const setPwd = reactive({
    username: '',
    password: '',
    newPassword: ''
  })
const setPwdUp = () => {
  client.post('/api/v1/auth/reset/password',JSON.stringify(setPwd)).then(res => {
    if(res.status == 200){
      ElNotification.success({title:'密码修改成功'})
    }else{
      ElNotification.success({title:'密码修改失败'})
    }
  }).catch(err => {
    console.log(err);
  })
}
//5.退出账号
const loginOut = () => {
  ElNotification.success({ title: '退出账号成功' })
  localStorage.clear()
  isAvatar.value = true
}
//监听全局变量
const globalVars = inject('globalVars')
//控制返回图标是否出现
const show = ref(false)
//点击时返回上级目录
const back = () => {
  router.back()
}
watch(() => globalVars.isInfo, (newValue) => {

  if (newValue) {
    show.value = true
  } else {
    show.value = false
  }
})
//点击箭头改变箭头样式
const checkRight = () => {
  if (globalVars.isRight == 1) {
    globalVars.isRight = 2
  } else if (globalVars.isRight == 2) {
    globalVars.isRight = 1
  } else
    globalVars.isRight = !globalVars.isRight
}
</script>
  
<style scoped lang="scss">
$background-color: rgb(239, 122, 122);
$icon-color: #ccc;

.btn {

  svg,
  .asideFont {
    color: rgba(191, 70, 113) !important;
  }
}

.common-layout,
.el-container {
  height: 100%;
  min-height: auto;

  .el-aside {

    display: flex;
    flex-direction: column;
    background: #F2F2F2;
    transition: 0.5s;
    border-radius: 0 8px 8px 0;
    width: 77px;
    align-items: center;
    justify-content: space-between;




    .iconBox {
      margin-top: 49px;
      display: flex;
      width: 70px;
      height: 326px;
      // margin-bottom: 20%;
      align-items: center;
      justify-content: center;
      flex-direction: column;

      .asideIcon {
        display: flex;
        flex-direction: column;
        width: 68px;
        height: 68px;
        align-items: center;
        justify-content: center;
        border: none;
        margin-bottom: 18px;

      }

      .asideFont {
        width: 24px;
        height: 16px;
        font-size: 12px;
        color: #797979;
        font-family: '更纱黑体';
        margin: 0 auto;
      }
    }

    .asideIcon:hover {

      svg,
      .asideFont {
        color: rgba(191, 70, 113) !important;

      }
    }


    .user {
      width: 70px;
      margin-bottom: 44px;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;



      .el-avatar {
        width: 48px;
        height: 48px;
        margin: auto;
        border: 1px solid #929290;
      }



      .switch {
        display: flex;
        flex-direction: column;
        width: 68px;
        height: 68px;
        align-items: center;
        justify-content: center;
        border: none;
        // margin-bottom: 18px;
      }
    }

    svg {
      color: rgba(97, 102, 109);
    }
  }

  .left {
    margin: 340px 0 0 0;
    background-color: #F2F2F2;
    width: 20px;
    height: 50px;
    display: flex;
    align-items: center;
    border-radius: 0 15px 15px 0;
  }

  .left:hover {
    background-color: #b7b7b7;
  }


}


.open {
  width: 0 !important;
  // animation: close 1.5s !important;
}

//动画无效
// @keyframes close{
//   from{width:77px;}
//   to{width: 1px;}
// }
.close {
  width: 77px !important;
  // animation: open 1.5s !important; 

}

//动画无效
// @keyframes open{
//   from{width:1px;}
//   to{width:77px;}
// }
.login {
  h2 {
    font-size: 24px;
    color: #333333;
    margin-bottom: 37px;
    font-family: 'Arial';
  }

  input {
    margin-bottom: 18px;
    border: 2px solid #EC808D;
    box-shadow: none;
    width: 274px;
    height: 34px;
    border-radius: 10px;
  }

  button {
    width: 80px;
    height: 30px;
    color: #F2F2F2;
    border: none;
    background-color: #EC808D;
    border-radius: 5px;
    font-family: '更纱黑体';
  }
}

.Sign {
  h2 {
    font-size: 24px;
    color: #333333;
    margin-bottom: 37px;
    font-family: 'Arial';
  }

  input {
    margin-bottom: 18px;
    border: 2px solid #EC808D;
    box-shadow: none;
    width: 274px;
    height: 34px;
    border-radius: 10px;
  }

  button {
    width: 80px;
    height: 30px;
    color: #F2F2F2;
    border: none;
    background-color: #EC808D;
    border-radius: 5px;
    font-family: '更纱黑体';
  }
}
.setPwd{
  h2 {
    font-size: 24px;
    color: #333333;
    margin-bottom: 37px;
    font-family: 'Arial';
  }

  input {
    margin-bottom: 18px;
    border: 2px solid #EC808D;
    box-shadow: none;
    width: 274px;
    height: 34px;
    border-radius: 10px;
  }

  button {
    width: 80px;
    height: 30px;
    color: #F2F2F2;
    border: none;
    background-color: #EC808D;
    border-radius: 5px;
    font-family: '更纱黑体';
  }
}
.switch {
  width: 68px;
  height: 68px;
  border: none;
}

html.dark .el-aside {
  animation: asideBackground 1.5s;
  background-color: rgb(67, 63, 63);

}

html.dark svg {
  color: rgba(146, 151, 157) !important;
}
</style>

