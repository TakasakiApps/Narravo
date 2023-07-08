<template>
    <div class="box">
        <img :src="img1" alt="">
        <div class="info">
            <h1>诡秘之主</h1>
            <h2>爱潜水的乌贼</h2>
            <h2 style="color: #7F7F7F;">玄幻 446.52万字</h2>
            <div style="display: flex;">
                <h3 style="margin-right: 32px;">最新章节：第八卷 愚者</h3>
                <h3>更新时间：2020-05-01</h3>
            </div>
            <h3>书籍简介</h3>
            <p style="width: 450px; " >蒸汽与机械的浪潮中，谁能触及非凡。历史和黑暗的迷雾里，又是谁在耳语。我从诡秘中醒来，睁眼看见这个世界：
                枪械，大炮，巨舰，飞空艇，差分机；魔药，占卜，诅咒，倒吊人，封印物……光明依旧照耀，神秘从未远离，这是一段“愚者”的传说</p>
            <div class="btn">
                <button style="width: 100px; height: 48px; margin-right: 20px;" @click="goRead('/read')">阅读</button>
                <button style="width: 142px; height: 48px; margin-right: 168px;">加入书架</button>
                <button style="width: 120px; height: 48px;">去制作</button>
            </div>
        </div>
    </div>
    <router-view></router-view>
</template>
    
<script setup lang='ts'>
import img1 from '../assets/诡秘.png'
import router from '../router'


const goRead = (path:any) => {
    router.push(path)
    console.log(path);
}
</script>
    
<style scoped lang="scss">
.box {
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 170px auto;

    img {
        width: 241px;
        height: 376px;
        margin-right: 105px;
    }

    .info {
        h1{
            font-family: 'Arial';
            font-size: 54px;
            color: #333333;
            margin-top: 0;
            margin-bottom: 24px;
        }
        h2{
            font-family: 'Roboto';
            color: #333333;
            font-size: 20px;
            margin-top: 0;
            margin-bottom: 24px;
        }
        h3{
            font-family: 'Arial';
            font-size: 20px;
            color:#7F7F7F ;
            margin-top: 0;
            margin-bottom: 24px;
        }
        p{
            font-family: 'Roboto';
            font-size: 18px;
            color: #7F7F7F;
        }
        
    }
    .btn{
        button{
            background-color: #EC808D;
            border: none;
            border-radius: 5px;
            color: #F2F2F2;
            font-family: '更纱黑体 SC';
            font-size: 20px;
        }
    }
}

button {}
</style>