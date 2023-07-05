<template>
    <div class="wrap">
        <div class="search">
            <input type="text" class="searchTerm" placeholder="What are you looking for?">
            <button type="submit" class="searchButton">
                <svg t="1688294823243" class="icon" viewBox="0 0 1036 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
                    p-id="2665" width="20" height="20">
                    <path
                        d="M1036.303023 935.288732l-88.711269 88.711268-217.439737-214.266852C653.420008 869.499937 558.103959 906.53851 453.269255 906.53851a453.269255 453.269255 0 1 1 453.269255-453.269255 449.772607 449.772607 0 0 1-88.51701 266.910839l218.281523 215.108638zM453.269255 129.505501a323.763754 323.763754 0 1 0 0.064753 647.59226A323.763754 323.763754 0 0 0 453.269255 129.505501z"
                        fill="#e6e6e6" p-id="2666"></path>
                </svg>
            </button>
        </div>

    </div>
    <ul class="box" v-infinite-scroll="load" infinite-scroll-immediate="false" style="overflow: auto; ">
        <el-scrollbar height="490px">
            <div v-for="i in 10" class="box-item">
                <BookLiem />
            </div>
        </el-scrollbar>
    </ul>
</template>
    
<script setup lang='ts'>
import BookLiem from '../components/BookLiem.vue'
import { ref } from 'vue'
const count = ref(2)
const load = () => {
    console.log(111);
}
</script>
    
<style scoped lang="scss">
@import url(https://fonts.googleapis.com/css?family=Open+Sans);

.wrap {
    width: 30%;
    position: absolute;
    top: 10%;
    left: 50%;
    transform: translate(-50%, -50%);
    display: block;

    .search {
        width: 100%;
        position: relative;
        display: flex;

        .searchTerm {
            width: 100%;
            border: 3px solid #00B4CC;
            border-right: none;
            padding: 5px;
            height: 25px;
            border-radius: 5px 0 0 5px;
            outline: none;
            color: #9DBFAF;
        }

        .searchTerm:focus {
            color: #00B4CC;
        }

        .searchButton {
            width: 40px;
            height: 40px;
            border: 1px solid #00B4CC;
            background: #00B4CC;
            text-align: center;
            color: #fff;
            border-radius: 0 5px 5px 0;
            cursor: pointer;
            font-size: 20px;
            display: flex;
            align-items: center;
        }
    }
}
.box {
    margin-top: 10%;
}
.box-item{
    transition: box-shadow 0.3s ease-in-out;
}
</style>